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

// options defines what elements can be overriden in a config file
type options struct {
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
	FreeSystemPageTableEnteries     uint64
	ModifiedPageListBytes           uint64
	PageFaultsPersec                uint64
	PageReadsPersec                 uint64
	PagesInputPersec                uint64
	PagesOutputPersec               uint64
	PagesPersec                     uint64
	PageWritesPersec                uint64
	PercentCommittedBytesInUse      uint64
	PoolNonpagedAllocs              uint64
	PoolPagedAllocs                 uint64
	PoolPagedBytes                  uint64
	PoolPagedResidentBytes          uint64
	StandbyCacheCoreBytes           uint64
	StandbyCacheNormalPriorityBytes uint64
	StandbyCacheResidentBytes       uint64
	SystemCacheResidentBytes        uint64
	SystemCodeResidentBytes         uint64
	SystemCodeTotalBytes            uint64
	SystemDriverTotalBytes          uint64
	TransitionFaultPersec           uint64
	TransitionPagesRePurposePersec  uint64
	WriteCopiesPersec               uint64
}

// NewMemoryCollector creates new wmi collector
func NewMemoryCollector(cfgBaseName string) (collector.Collector, error) {
	id := "memory"
	c := Memory{}

	var dst Win32_PerfFormattedData_PerfOS_Memory
	c.query = wmi.CreateQuery(&dst, "")

	c.id = id
	c.logger = log.With().Str("pkg", "builtins.wmi.memory").Logger()
	c.metricStatus = map[string]bool{}
	c.metricDefaultActive = true
	c.lastMetrics = cgm.Metrics{}

	if cfgBaseName == "" {
		return &c, nil
	}

	var cfg options
	err := config.LoadConfigFile(cfgBaseName, &cfg)
	if err != nil {
		p.logger.Debug().Err(err).Str("file", cfgBaseName).Msg("loading config file")
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

	var dst Win32_PerfFormattedData_PerfOS_Processor
	if err := wmi.Query(c.query, &dst); err != nil {
		c.setStatus(metrics, err)
		return errors.Wrap(err, "wmi.memory")
	}

	pfx := c.id + "`"
	d := structs.Map(dst)
	for name, val := range d {
		addMetric(pfx+name, "L", val)
	}

	// addMetric(pfx+"AvailableBytes", "L", dst.AvailableBytes)
	// addMetric(pfx+"CacheBytes", "L", dst.CacheBytes)
	// addMetric(pfx+"CacheFaultsPersec", "L", dst.CacheFaultsPersec)
	// addMetric(pfx+"CommittedBytes", "L", dst.CommittedBytes)
	// addMetric(pfx+"DemandZeroFaultsPersec", "L", dst.DemandZeroFaultsPersec)
	// addMetric(pfx+"FreeAndZeroPageListBytes", "L", dst.FreeAndZeroPageListBytes)
	// addMetric(pfx+"FreeSystemPageTableEnteries", "L", dst.FreeSystemPageTableEnteries)
	// addMetric(pfx+"ModifiedPageListBytes", "L", dst.ModifiedPageListBytes)
	// addMetric(pfx+"PageFaultsPersec", "L", dst.PageFaultsPersec)
	// addMetric(pfx+"PageReadsPersec", "L", dst.PageReadsPersec)
	// addMetric(pfx+"PagesInputPersec", "L", dst.PagesInputPersec)
	// addMetric(pfx+"PagesOutputPersec", "L", dst.PagesOutputPersec)
	// addMetric(pfx+"PagesPersec", "L", dst.PagesPersec)
	// addMetric(pfx+"PageWritesPersec", "L", dst.PageWritesPersec)
	// addMetric(pfx+"PercentCommittedBytesInUse", "L", dst.PercentCommittedBytesInUse)
	// addMetric(pfx+"PoolNonpagedAllocs", "L", dst.PoolNonpagedAllocs)
	// addMetric(pfx+"PoolPagedAllocs", "L", dst.PoolPagedAllocs)
	// addMetric(pfx+"PoolPagedBytes", "L", dst.PoolPagedBytes)
	// addMetric(pfx+"PoolPagedResidentBytes", "L", dst.PoolPagedResidentBytes)
	// addMetric(pfx+"StandbyCacheCoreBytes", "L", dst.StandbyCacheCoreBytes)
	// addMetric(pfx+"StandbyCacheNormalPriorityBytes", "L", dst.StandbyCacheNormalPriorityBytes)
	// addMetric(pfx+"StandbyCacheResidentBytes", "L", dst.StandbyCacheResidentBytes)
	// addMetric(pfx+"SystemCacheResidentBytes", "L", dst.SystemCacheResidentBytes)
	// addMetric(pfx+"SystemCodeResidentBytes", "L", dst.SystemCodeResidentBytes)
	// addMetric(pfx+"SystemCodeTotalBytes", "L", dst.SystemCodeTotalBytes)
	// addMetric(pfx+"SystemDriverTotalBytes", "L", dst.SystemDriverTotalBytes)
	// addMetric(pfx+"TransitionFaultPersec", "L", dst.TransitionFaultPersec)
	// addMetric(pfx+"TransitionPagesRePurposePersec", "L", dst.TransitionPagesRePurposePersec)
	// addMetric(pfx+"WriteCopiesPersec", "L", dst.WriteCopiesPersec)

	c.setStatus(metrics, nil)
	return nil
}
