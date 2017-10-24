// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build !windows

package wmi

import "github.com/circonus-labs/circonus-agent/internal/builtins/collector"

// NewCPUCollector creates new procfs cpu collector
func NewCPUCollector(cfgFile string) (collector.Collector, error) {
	return &wmicommon{
		id:        "not_implemented",
		lastError: collector.ErrNotImplemented,
	}, collector.ErrNotImplemented
}
