// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

import (
	"regexp"
	"strings"
	"time"

	"github.com/StackExchange/wmi"
	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
	"github.com/circonus-labs/circonus-agent/internal/config"
	cgm "github.com/circonus-labs/circonus-gometrics"
	"github.com/fatih/structs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Memory metrics from the Windows Management Interface (wmi)
type Memory struct {
	wmicommon
}

// memoryOptions defines what elements can be overriden in a config file
type memoryOptions struct {
	ID                   string   `json:"id" toml:"id" yaml:"id"`
	MetricsEnabled       []string `json:"metrics_enabled" toml:"metrics_enabled" yaml:"metrics_enabled"`
	MetricsDisabled      []string `json:"metrics_disabled" toml:"metrics_disabled" yaml:"metrics_disabled"`
	MetricsDefaultStatus string   `json:"metrics_default_status" toml:"metrics_default_status" toml:"metrics_default_status"`
	RunTTL               string   `json:"run_ttl" toml:"run_ttl" yaml:"run_ttl"`
}

// Win32_PerfFormattedData_PerfOS_Memory defines the metrics to collect
type Win32_PerfFormattedData_PerfOS_Memory struct {
	AvailableBytes                  uint64
	CacheBytes                      uint64
	CacheFaultsPersec               uint64
	CommittedBytes                  uint64
	DemandZeroFaultsPersec          uint64
	FreeAndZeroPageListBytes        uint64
	FreeSystemPageTableEntries      uint64
	ModifiedPageListBytes           uint64
	PageFaultsPersec                uint64
	PageReadsPersec                 uint64
	PagesInputPersec                uint64
	PagesOutputPersec               uint64
	PagesPersec                     uint64
	PageWritesPersec                uint64
	PercentCommittedBytesInUse      uint64
	PoolNonpagedAllocs              uint64
	PoolNonpagedBytes               uint64
	PoolPagedAllocs                 uint64
	PoolPagedBytes                  uint64
	PoolPagedResidentBytes          uint64
	StandbyCacheCoreBytes           uint64
	StandbyCacheNormalPriorityBytes uint64
	StandbyCacheReserveBytes        uint64
	SystemCacheResidentBytes        uint64
	SystemCodeResidentBytes         uint64
	SystemCodeTotalBytes            uint64
	SystemDriverTotalBytes          uint64
	TransitionFaultsPersec          uint64
	TransitionPagesRePurposedPersec uint64
	WriteCopiesPersec               uint64
}

// NewMemoryCollector creates new wmi collector
func NewMemoryCollector(cfgBaseName string) (collector.Collector, error) {
	id := "memory"
	c := Memory{}

	var dst []Win32_PerfFormattedData_PerfOS_Memory
	c.query = wmi.CreateQuery(&dst, "")

	c.id = id
	c.logger = log.With().Str("pkg", "builtins.wmi.memory").Logger()
	c.metricStatus = map[string]bool{}
	c.metricDefaultActive = true
	c.lastMetrics = cgm.Metrics{}

	if cfgBaseName == "" {
		return &c, nil
	}

	var cfg memoryOptions
	err := config.LoadConfigFile(cfgBaseName, &cfg)
	if err != nil {
		c.logger.Debug().Err(err).Str("file", cfgBaseName).Msg("loading config file")
		if strings.Contains(err.Error(), "no config found matching") {
			return &c, nil
		}
		return nil, errors.Wrap(err, "wmi.memory config")
	}

	c.logger.Debug().Interface("config", cfg).Msg("loaded config")

	if cfg.ID != "" {
		c.id = cfg.ID
	}

	if len(cfg.MetricsEnabled) > 0 {
		for _, name := range cfg.MetricsEnabled {
			c.metricStatus[name] = true
		}
	}
	if len(cfg.MetricsDisabled) > 0 {
		for _, name := range cfg.MetricsDisabled {
			c.metricStatus[name] = false
		}
	}

	if cfg.MetricsDefaultStatus != "" {
		if ok, _ := regexp.MatchString(`^(enabled|disabled)$`, strings.ToLower(cfg.MetricsDefaultStatus)); ok {
			c.metricDefaultActive = strings.ToLower(cfg.MetricsDefaultStatus) == "enabled"
		} else {
			return nil, errors.Errorf("wmi.memory invalid metric default status (%s)", cfg.MetricsDefaultStatus)
		}
	}

	if cfg.RunTTL != "" {
		dur, err := time.ParseDuration(cfg.RunTTL)
		if err != nil {
			return nil, errors.Wrap(err, "wmi.memory parsing run_ttl")
		}
		c.runTTL = dur
	}

	return &c, nil
}

// Flush returns the last metrics
func (c *Memory) Flush() cgm.Metrics {
	c.Lock()
	defer c.Unlock()
	return c.lastMetrics
}

// Collect metrics from the wmi resource
func (c *Memory) Collect() error {
	metrics := cgm.Metrics{}

	c.Lock()

	if c.runTTL > time.Duration(0) {
		if time.Since(c.lastEnd) < c.runTTL {
			c.logger.Warn().Msg(collector.ErrTTLNotExpired.Error())
			c.Unlock()
			return collector.ErrTTLNotExpired
		}
	}
	if c.running {
		c.logger.Warn().Msg(collector.ErrAlreadyRunning.Error())
		c.Unlock()
		return collector.ErrAlreadyRunning
	}

	c.running = true
	c.lastStart = time.Now()
	c.Unlock()

	var dst []Win32_PerfFormattedData_PerfOS_Memory
	if err := wmi.Query(c.query, &dst); err != nil {
		c.logger.Error().Err(err).Str("query", c.query).Msg("wmi error")
		c.setStatus(metrics, err)
		return errors.Wrap(err, "wmi.memory")
	}

	pfx := c.id + "`"
	d := structs.Map(dst[0])
	for name, val := range d {
		c.addMetric(&metrics, pfx+name, "L", val)
	}

	c.setStatus(metrics, nil)
	return nil
}
