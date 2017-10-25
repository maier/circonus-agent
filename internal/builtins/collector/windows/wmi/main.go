// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

import (
	"path"
	"time"

	"github.com/StackExchange/wmi"
	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
	"github.com/circonus-labs/circonus-agent/internal/config/defaults"
	cgm "github.com/circonus-labs/circonus-gometrics"
	"github.com/pkg/errors"
)

func initialize() error {
	// This initialization prevents a memory leak on WMF 5+. See
	// https://github.com/martinlindhe/wmi_exporter/issues/77 and linked issues
	// for details.
	s, err := wmi.InitializeSWbemServices(wmi.DefaultClient)
	if err != nil {
		return err
	}
	wmi.DefaultClient.SWbemServicesClient = s
	return nil
}

// New creates new WMI collector
func New() ([]collector.Collector, error) {

	if err := initialize(); err != nil {
		return []collector.Collector{}, err
	}

	collectors := make([]collector.Collector, 10)

	c, err := NewCPUCollector(path.Join(defaults.EtcPath, "cpu.json"))
	if err != nil {
		return []collector.Collector{}, errors.Wrap(err, "initializing wmi.cpu")
	}
	collectors = append(collectors, c)

	return collectors, nil
}

// Collect returns collector metrics
func (c *wmicommon) Collect() error {
	c.Lock()
	defer c.Unlock()
	return collector.ErrNotImplemented
}

// Flush returns last metrics collected
func (c *wmicommon) Flush() cgm.Metrics {
	c.Lock()
	defer c.Unlock()
	return cgm.Metrics{}
}

// ID returns the id of the instance
func (c *wmicommon) ID() string {
	c.Lock()
	defer c.Unlock()
	return c.id
}

// Inventory returns collector stats for /inventory endpoint
func (c *wmicommon) Inventory() collector.InventoryStats {
	c.Lock()
	defer c.Unlock()
	return collector.InventoryStats{
		ID:              c.id,
		LastRunStart:    c.lastStart.Format(time.RFC3339Nano),
		LastRunEnd:      c.lastEnd.Format(time.RFC3339Nano),
		LastRunDuration: c.lastRunDuration.String(),
		LastError:       c.lastError.Error(),
	}
}
