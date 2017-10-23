// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package procfs

import (
	"bufio"
	"encoding/json"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
	cgm "github.com/circonus-labs/circonus-gometrics"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// CPU metrics from the Linux ProcFS
type CPU struct {
	pfscommon
	numCPU        float64
	clockHZ       float64 // getconf CLK_TCK, may be overriden in config file
	reportAllCPUs bool    // may be overriden in config file
}

// config defines what elements can be overriden in a config file
type config struct {
	File                string          `json:"proc_file"`
	ClockHZ             float64         `json:"clock_hz"`
	AllCPU              bool            `json:"report_all_cpus"`
	Metrics             map[string]bool `json:"metrics"`
	DefaultMetricStatus string          `json:"metric_default_status"`
}

// NewCPUCollector creates new procfs cpu collector
func NewCPUCollector(cfgFile string) (collector.Collector, error) {
	id := "cpu"
	cpu := CPU{}

	cpu.id = id
	cpu.file = "/proc/stat"
	cpu.logger = log.With().Str("pkg", "builtins.procfs.cpu").Logger()
	cpu.numCPU = float64(runtime.NumCPU())
	cpu.metricStatus = map[string]bool{}
	cpu.defaultStatus = "active"
	cpu.clockHZ = 100
	cpu.reportAllCPUs = true

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

		cpu.reportAllCPUs = cfg.AllCPU

		if cfg.File != "" {
			cpu.file = cfg.File
		}

		if cfg.ClockHZ > 0 {
			cpu.clockHZ = cfg.ClockHZ
		}

		if len(cfg.Metrics) > 0 {
			cpu.metricStatus = cfg.Metrics
		}

		if ok, _ := regexp.MatchString(`^(active|disabled)$`, cfg.DefaultMetricStatus); ok {
			cpu.metricDefaultStatus = cfg.DefaultMetricStatus
		}
	}

	if _, err := os.Stat(cpu.file); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "procfs")
	}

	return &cpu, nil
}

// Flush returns the last metrics
func (c *CPU) Flush() cgm.Metrics {
	c.Lock()
	defer c.Unlock()
	return c.lastMetrics
}

// Collect metrics from the procfs resource
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
			c.lastMetrics = &cgm.Metrics{}
		}
		c.running = false
		c.Unlock()
	}

	c.running = true
	c.lastStart = time.Now()
	c.Unlock()

	f, err := os.Open(c.file)
	if err != nil {
		resetStatus(err)
		return errors.Wrap(err, "procfs.cpu")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		line := scanner.Text()
		fields := strings.Fields(line)

		switch {
		case fields[0] == "processes":
			metricName := fields[0]
			found, active := c.metricStatus[metricName]
			if found && !active {
				continue
			}
			if !found && c.defaultStatus != "active" {
				continue
			}
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				resetStatus(err)
				return errors.Wrapf(err, "parsing %s", fields[0])
			}
			metrics[metricName] = cgm.Metric{
				Type:  "L",
				Value: v,
			}

		case fields[0] == "procs_running":
			metricName := "procs_runnable"
			found, active := c.metricStatus[metricName]
			if found && !active {
				continue
			}
			if !found && c.defaultStatus != "active" {
				continue
			}
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				resetStatus(err)
				return errors.Wrapf(err, "parsing %s", fields[0])
			}
			metrics[metricName] = cgm.Metric{
				Type:  "L",
				Value: v,
			}

		case fields[0] == "procs_blocked":
			metricName := fields[0]
			found, active := c.metricStatus[metricName]
			if found && !active {
				continue
			}
			if !found && c.defaultStatus != "active" {
				continue
			}
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				resetStatus(err)
				return errors.Wrapf(err, "parsing %s", fields[0])
			}
			metrics[metricName] = cgm.Metric{
				Type:  "L",
				Value: v,
			}

		case fields[0] == "ctx":
			metricName := "context_switch"
			found, active := c.metricStatus[metricName]
			if found && !active {
				continue
			}
			if !found && c.defaultStatus != "active" {
				continue
			}
			v, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				resetStatus(err)
				return errors.Wrapf(err, "parsing %s", fields[0])
			}
			metrics[metricName] = cgm.Metric{
				Type:  "L",
				Value: v,
			}

		case strings.HasPrefix(fields[0], "cpu"):
			if fields[0] != "cpu" && !c.reportAllCPUs {
				continue
			}
			cpuMetrics, err := c.parseCPU(fields)
			if err != nil {
				resetStatus(err)
				return errors.Wrapf(err, "parsing %s", fields[0])
			}
			for mn, mv := range *cpuMetrics {
				found, active := c.metricStatus[mn]
				if found && !active {
					continue
				}
				if !found && c.defaultStatus != "active" {
					continue
				}
				metrics[mn] = mv
			}
		}
	}

	if err := scanner.Err(); err != nil {
		resetStatus(err)
		return errors.Wrapf(err, "parsing %s", f.Name())
	}

	c.Lock()
	c.lastMetrics = metrics
	c.Unlock()

	resetStatus(nil)
	return nil
}

func (c *CPU) parseCPU(fields []string) (*cgm.Metrics, error) {
	var numCPU float64
	var metricBase string

	if fields[0] == "cpu" {
		metricBase = c.id + "`all"
		numCPU = c.numCPU // aggregate cpu metrics
	} else {
		metricBase = c.id + "`" + fields[0]
		numCPU = 1 // individual cpu metrics
	}

	metricType := "n" // resmon double

	userNormal, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return nil, err
	}

	userNice, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return nil, err
	}

	sys, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return nil, err
	}

	idleNormal, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return nil, err
	}

	waitIO, err := strconv.ParseFloat(fields[5], 64)
	if err != nil {
		return nil, err
	}

	irq, err := strconv.ParseFloat(fields[6], 64)
	if err != nil {
		return nil, err
	}

	softIRQ, err := strconv.ParseFloat(fields[7], 64)
	if err != nil {
		return nil, err
	}

	steal := float64(0)
	if len(fields) > 8 {
		v, err := strconv.ParseFloat(fields[8], 64)
		if err != nil {
			return nil, err
		}
		steal = v
	}

	guest := float64(0)
	if len(fields) > 9 {
		v, err := strconv.ParseFloat(fields[9], 64)
		if err != nil {
			return nil, err
		}
		guest = v
	}

	guestNice := float64(0)
	if len(fields) > 10 {
		v, err := strconv.ParseFloat(fields[10], 64)
		if err != nil {
			return nil, err
		}
		guestNice = v
	}

	metrics := cgm.Metrics{
		metricBase + "`user":              cgm.Metric{Type: metricType, Value: ((userNormal + userNice) / numCPU) / c.clockHZ},
		metricBase + "`user`normal":       cgm.Metric{Type: metricType, Value: (userNormal / numCPU) / c.clockHZ},
		metricBase + "`user`nice":         cgm.Metric{Type: metricType, Value: (userNice / numCPU) / c.clockHZ},
		metricBase + "`kernel":            cgm.Metric{Type: metricType, Value: ((sys + guest + guestNice) / numCPU) / c.clockHZ},
		metricBase + "`kernel`sys":        cgm.Metric{Type: metricType, Value: (sys / numCPU) / c.clockHZ},
		metricBase + "`kernel`guest":      cgm.Metric{Type: metricType, Value: (guest / numCPU) / c.clockHZ},
		metricBase + "`kernel`guest_nice": cgm.Metric{Type: metricType, Value: (guestNice / numCPU) / c.clockHZ},
		metricBase + "`idle":              cgm.Metric{Type: metricType, Value: ((idleNormal + steal) / numCPU) / c.clockHZ},
		metricBase + "`idle`normal":       cgm.Metric{Type: metricType, Value: (idleNormal / numCPU) / c.clockHZ},
		metricBase + "`idle`steal":        cgm.Metric{Type: metricType, Value: (steal / numCPU) / c.clockHZ},
		metricBase + "`wait_io":           cgm.Metric{Type: metricType, Value: (waitIO / numCPU) / c.clockHZ},
		metricBase + "`intr":              cgm.Metric{Type: metricType, Value: ((irq + softIRQ) / numCPU) / c.clockHZ},
		metricBase + "`intr`soft":         cgm.Metric{Type: metricType, Value: (irq / numCPU) / c.clockHZ},
		metricBase + "`intr`hard":         cgm.Metric{Type: metricType, Value: (softIRQ / numCPU) / c.clockHZ},
	}

	return &metrics, nil
}
