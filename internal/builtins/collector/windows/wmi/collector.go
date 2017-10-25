// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

import (
	"time"

	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
	cgm "github.com/circonus-labs/circonus-gometrics"
)

// Collect metrics
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

// ID returns id of collector
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
