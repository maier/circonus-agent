// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

import (
	"time"

	"github.com/StackExchange/wmi"
	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
	cgm "github.com/circonus-labs/circonus-gometrics"
	"github.com/pkg/errors"
)

func init() {
	// This initialization prevents a memory leak on WMF 5+. See
	// https://github.com/martinlindhe/wmi_exporter/issues/77 and linked issues
	// for details.
	s, err := wmi.InitializeSWbemServices(wmi.DefaultClient)
	if err != nil {
		return err
	}
	wmi.DefaultClient.SWbemServicesClient = s
}

// New creates new WMI collector
func New(cfgFile string) ([]collector.Collector, error) {

	collectors := make([]collector.Collector, 10)

	c, err := NewCPUCollector()
	if err != nil {
		return errors.Wrap(err, "initializing wmi.cpu")
	}
	collectors = append(collectors, c)

	return collectors, nil
}

// Collect returns collector metrics
func (p *wmicommon) Collect() error {
	p.Lock()
	defer p.Unlock()
	return collector.ErrNotImplemented
}

// Flush returns last metrics collected
func (p *wmicommon) Flush() cgm.Metrics {
	p.Lock()
	defer p.Unlock()
	return cgm.Metrics{}
}

// ID returns the id of the instance
func (p *wmicommon) ID() string {
	p.Lock()
	defer p.Unlock()
	return p.id
}

// Inventory returns collector stats for /inventory endpoint
func (p *wmicommon) Inventory() collector.InventoryStats {
	p.Lock()
	defer p.Unlock()
	return collector.InventoryStats{
		ID:              p.id,
		LastRunStart:    p.lastStart.Format(time.RFC3339Nano),
		LastRunEnd:      p.lastEnd.Format(time.RFC3339Nano),
		LastRunDuration: p.lastRunDuration.String(),
		LastError:       p.lastError.Error(),
	}
}
