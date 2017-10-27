// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

import (
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/StackExchange/wmi"
	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
	"github.com/circonus-labs/circonus-agent/internal/config"
	cgm "github.com/circonus-labs/circonus-gometrics"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Processor metrics from the Windows Management Interface (wmi)
type Processor struct {
	wmicommon
	numCPU        float64
	reportAllCPUs bool // may be overriden in config file
}

// processorOptions defines what elements can be overriden in a config file
type processorOptions struct {
	ID                   string   `json:"id" toml:"id" yaml:"id"`
	AllCPU               string   `json:"report_all_cpus" toml:"report_all_cpus" yaml:"report_all_cpus"`
	MetricsEnabled       []string `json:"metrics_enabled" toml:"metrics_enabled" yaml:"metrics_enabled"`
	MetricsDisabled      []string `json:"metrics_disabled" toml:"metrics_disabled" yaml:"metrics_disabled"`
	MetricsDefaultStatus string   `json:"metrics_default_status" toml:"metrics_default_status" toml:"metrics_default_status"`
	RunTTL               string   `json:"run_ttl" toml:"run_ttl" yaml:"run_ttl"`
}

// Win32_PerfFormattedData_PerfOS_Processor defines the metrics to collect
type Win32_PerfFormattedData_PerfOS_Processor struct {
	Name                  string
	C1TransitionsPersec   uint64
	C2TransitionsPersec   uint64
	C3TransitionsPersec   uint64
	DPCsQueuedPersec      uint32
	InterruptsPersec      uint32
	PercentC1Time         uint64
	PercentC2Time         uint64
	PercentC3Time         uint64
	PercentDPCTime        uint64
	PercentIdleTime       uint64
	PercentInterruptTime  uint64
	PercentPrivilegedTime uint64
	PercentProcessorTime  uint64
	PercentUserTime       uint64
}

// NewProcessorCollector creates new wmi collector
func NewProcessorCollector(cfgBaseName string) (collector.Collector, error) {
	id := "processor"
	c := Processor{}

	var dst []Win32_PerfFormattedData_PerfOS_Processor
	c.query = wmi.CreateQuery(&dst, "")

	c.id = id
	c.logger = log.With().Str("pkg", "builtins.wmi.processor").Logger()
	c.numCPU = float64(runtime.NumCPU())
	c.metricStatus = map[string]bool{}
	c.metricDefaultActive = true
	c.reportAllCPUs = true
	c.lastMetrics = cgm.Metrics{}

	if cfgBaseName == "" {
		return &c, nil
	}

	var cfg processorOptions
	err := config.LoadConfigFile(cfgBaseName, &cfg)
	if err != nil {
		c.logger.Debug().Err(err).Str("file", cfgBaseName).Msg("loading config file")
		if strings.Contains(err.Error(), "no config found matching") {
			return &c, nil
		}
		return nil, errors.Wrap(err, "wmi.processor config")
	}

	c.logger.Debug().Interface("config", cfg).Msg("loaded config")

	if cfg.ID != "" {
		c.id = cfg.ID
	}

	if cfg.AllCPU != "" {
		rpt, err := strconv.ParseBool(cfg.AllCPU)
		if err != nil {
			return nil, errors.Wrap(err, "wmi.processor parsing report_all_cpus")
		}
		c.reportAllCPUs = rpt
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
			return nil, errors.Errorf("wmi.processor invalid metric default status (%s)", cfg.MetricsDefaultStatus)
		}
	}

	if cfg.RunTTL != "" {
		dur, err := time.ParseDuration(cfg.RunTTL)
		if err != nil {
			return nil, errors.Wrap(err, "wmi.processor parsing run_ttl")
		}
		c.runTTL = dur
	}

	return &c, nil
}

// Flush returns the last metrics
func (c *Processor) Flush() cgm.Metrics {
	c.Lock()
	defer c.Unlock()
	return c.lastMetrics
}

// Collect metrics from the wmi resource
func (c *Processor) Collect() error {
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

	var dst []Win32_PerfFormattedData_PerfOS_Processor
	if err := wmi.Query(c.query, &dst); err != nil {
		c.setStatus(metrics, err)
		return errors.Wrap(err, "wmi.processor")
	}

	for _, group := range dst {
		pfx := c.id + "`"
		if strings.Contains(group.Name, "_Total") {
			pfx += "all"
		} else {
			if !c.reportAllCPUs {
				continue
			}
			pfx += group.Name
		}

		c.addMetric(&metrics, pfx+"PercentC1Time", "L", group.PercentC1Time)
		c.addMetric(&metrics, pfx+"PercentC2Time", "L", group.PercentC2Time)
		c.addMetric(&metrics, pfx+"PercentC3Time", "L", group.PercentC3Time)
		c.addMetric(&metrics, pfx+"PercentIdleTime", "L", group.PercentIdleTime)
		c.addMetric(&metrics, pfx+"PercentInterruptTime", "L", group.PercentInterruptTime)
		c.addMetric(&metrics, pfx+"PercentDPCTime", "L", group.PercentDPCTime)
		c.addMetric(&metrics, pfx+"PercentPrivilegedTime", "L", group.PercentPrivilegedTime)
		c.addMetric(&metrics, pfx+"PercentUserTime", "L", group.PercentUserTime)
		c.addMetric(&metrics, pfx+"PercentProcessorTime", "L", group.PercentProcessorTime)
		c.addMetric(&metrics, pfx+"C1TransitionsPersec", "L", group.C1TransitionsPersec)
		c.addMetric(&metrics, pfx+"C2TransitionsPersec", "L", group.C2TransitionsPersec)
		c.addMetric(&metrics, pfx+"C3TransitionsPersec", "L", group.C3TransitionsPersec)
		c.addMetric(&metrics, pfx+"InterruptsPersec", "L", group.InterruptsPersec)
		c.addMetric(&metrics, pfx+"DPCsQueuedPersec", "L", group.DPCsQueuedPersec)
	}

	c.setStatus(metrics, nil)
	return nil
}
