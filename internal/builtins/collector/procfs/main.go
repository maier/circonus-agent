// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package procfs

import (
	"time"

	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
	cgm "github.com/circonus-labs/circonus-gometrics"
)

// New creates new procfs collector
func New(cfgFile string) (collector.Collector, error) {
	return &pfscommon{id: "not_implemented"}, collector.ErrNotImplemented
}

// Collect returns collector metrics
func (p *pfscommon) Collect() error {
	return collector.ErrNotImplemented
}

// Flush returns last metrics collected
func (p *pfscommon) Flush() cgm.Metrics {
	return cgm.Metrics{}
}

// ID returns the id of the instance
func (p *pfscommon) ID() (string, error) {
	return p.id, nil
}

// Inventory returns collector stats for /inventory endpoint
func (p *pfscommon) Inventory() (collector.InventoryStats, error) {
	return collector.InventoryStats{
		ID:              p.id,
		LastRunStart:    p.lastStart.Format(time.RFC3339Nano),
		LastRunEnd:      p.lastEnd.Format(time.RFC3339Nano),
		LastRunDuration: p.lastRunDuration.String(),
		LastError:       p.lastError.Error(),
	}, nil
}
