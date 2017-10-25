// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

import (
	"encoding/json"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/StackExchange/wmi"
	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
	"github.com/circonus-labs/circonus-agent/internal/config/defaults"
	cgm "github.com/circonus-labs/circonus-gometrics"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// CPU metrics from the Windows Management Interface (wmi)
type CPU struct {
	wmicommon
	numCPU        float64
	reportAllCPUs bool // may be overriden in config file
}

// config defines what elements can be overriden in a config file
type config struct {
	ID                  string          `json:"id"`
	AllCPU              bool            `json:"report_all_cpus"`
	Metrics             map[string]bool `json:"metrics"`
	DefaultMetricStatus string          `json:"metric_default_status"`
	RunTTL              string          `json:"run_ttl"`
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

// NewCPUCollector creates new wmi cpu collector
func NewCPUCollector() (collector.Collector, error) {
	id := "cpu"
	cpu := CPU{}

	var dst []Win32_PerfFormattedData_PerfOS_Processor
	cpu.query = wmi.CreateQuery(&dst, "")

	cpu.id = id
	cpu.logger = log.With().Str("pkg", "builtins.wmi.cpu").Logger()
	cpu.numCPU = float64(runtime.NumCPU())
	cpu.metricStatus = map[string]bool{}
	cpu.metricStatusDefault = "active"
	cpu.reportAllCPUs = true
	cpu.lastMetrics = cgm.Metrics{}

	cfgFile := path.Join(defaults.EtcPath, "cpu.json")
	if _, err := os.Stat(cfg); os.IsNotExist(err) {
		cfg = ""
	}

	if cfgFile != "" {
		f, err := os.Open(cfgFile)
		if err != nil {
			return nil, errors.Wrap(err, "config file")
		}
		defer f.Close()

		var cfg config
		dec := json.NewDecoder(f)
		if err := dec.Decode(&cfg); err != nil {
			return nil, errors.Wrapf(err, "parsing config file %s", cfgFile)
		}

		if cfg.ID != "" {
			cpu.id = cfg.ID
		}

		cpu.reportAllCPUs = cfg.AllCPU

		if len(cfg.Metrics) > 0 {
			cpu.metricStatus = cfg.Metrics
		}

		if ok, _ := regexp.MatchString(`^(active|disabled)$`, cfg.DefaultMetricStatus); ok {
			cpu.metricStatusDefault = cfg.DefaultMetricStatus
		}

		if cfg.RunTTL != "" {
			dur, err := time.ParseDuration(cfg.RunTTL)
			if err != nil {
				return nil, errors.Wrapf(err, "parsing config file %s", cfgFile)
			}
			cpu.runTTL = dur
		}
	}

	return &cpu, nil
}

// Flush returns the last metrics
func (c *CPU) Flush() cgm.Metrics {
	c.Lock()
	defer c.Unlock()
	return c.lastMetrics
}

// Collect metrics from the wmi resource
func (c *CPU) Collect() error {
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

	resetStatus := func(err error) {
		c.Lock()
		c.lastEnd = time.Now()
		c.lastRunDuration = time.Since(c.lastStart)
		c.lastError = err
		if err != nil {
			// on error, ensure metrics are reset
			// do not keep returning a stale set of metrics
			c.lastMetrics = cgm.Metrics{}
		}
		c.running = false
		c.Unlock()
	}

	c.running = true
	c.lastStart = time.Now()
	c.Unlock()

	var dst []Win32_PerfFormattedData_PerfOS_Processor
	if err := wmi.Query(c.query, &dst); err != nil {
		resetStatus(err)
		return errors.Wrap(err, "wmi.cpu")
	}

	addMetric := func(mname, mtype string, mval interface{}) {
		found, active := c.metricStatus[mname]
		if (found && active) || (!found && c.metricStatusDefault == "active") {
			metrics[mname] = cgm.Metric{Type: mtype, Value: mval}
		}
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

		addMetric(pfx+"PercentC1Time", "L", group.PercentC1Time)
		addMetric(pfx+"PercentC2Time", "L", group.PercentC2Time)
		addMetric(pfx+"PercentC3Time", "L", group.PercentC3Time)
		addMetric(pfx+"PercentIdleTime", "L", group.PercentIdleTime)
		addMetric(pfx+"PercentInterruptTime", "L", group.PercentInterruptTime)
		addMetric(pfx+"PercentDPCTime", "L", group.PercentDPCTime)
		addMetric(pfx+"PercentPrivilegedTime", "L", group.PercentPrivilegedTime)
		addMetric(pfx+"PercentUserTime", "L", group.PercentUserTime)
		addMetric(pfx+"PercentProcessorTime", "L", group.PercentProcessorTime)
		addMetric(pfx+"C1TransitionsPersec", "L", group.C1TransitionsPersec)
		addMetric(pfx+"C2TransitionsPersec", "L", group.C2TransitionsPersec)
		addMetric(pfx+"C3TransitionsPersec", "L", group.C3TransitionsPersec)
		addMetric(pfx+"InterruptsPersec", "L", group.InterruptsPersec)
		addMetric(pfx+"DPCsQueuedPersec", "L", group.DPCsQueuedPersec)
	}

	c.Lock()
	c.lastMetrics = metrics
	c.Unlock()

	resetStatus(nil)
	return nil
}
