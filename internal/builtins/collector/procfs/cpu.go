// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build !linux

package procfs

import "github.com/circonus-labs/circonus-agent/internal/builtins/collector"

// NewCPUMetrics creates new procfs cpu collector
func NewCPUMetrics(cfgFile string) (collector.Collector, error) {
	return &pfscommon{
		id:        "not_implemented",
		lastError: collector.ErrNotImplemented,
	}, collector.ErrNotImplemented
}
