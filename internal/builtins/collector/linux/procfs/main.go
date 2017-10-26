// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package procfs

import (
	"path"
	"runtime"

	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
	"github.com/circonus-labs/circonus-agent/internal/config/defaults"
	"github.com/pkg/errors"
)

// New creates new ProcFS collector
func New() ([]collector.Collector, error) {
	if runtime.GOOS != "linux" {
		return []collector.Collector{}, nil
	}

	collectors := make([]collector.Collector, 10)

	c, err := NewCPUCollector(path.Join(defaults.EtcPath, "cpu"))
	if err != nil {
		return []collector.Collector{}, errors.Wrap(err, "initializing procfs.cpu")
	}

	collectors = append(collectors, c)

	return collectors, nil
}
