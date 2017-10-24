// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build linux

package builtins

import (
	"os"
	"path"

	"github.com/circonus-labs/circonus-agent/internal/builtins/collector/linux/procfs"
	"github.com/circonus-labs/circonus-agent/internal/config/defaults"
	"github.com/pkg/errors"
)

func (b *Builtins) configure() error {
	// CPU metrics
	{
		cfg := path.Join(defaults.EtcPath, "cpu.json")
		if _, err := os.Stat(cfg); os.IsNotExist(err) {
			cfg = ""
		}

		cpu, err := procfs.NewCPUCollector(cfg)
		if err != nil {
			return errors.Wrap(err, "procfs.cpu")
		}

		b.collectors[cpu.ID()] = cpu
	}

	return nil
}
