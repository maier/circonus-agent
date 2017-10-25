// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package builtins

import (
	"github.com/circonus-labs/circonus-agent/internal/builtins/collector/windows/wmi"
)

func (b *Builtins) configure() error {
	collectors, err := wmi.New()
	for _, c := range collectors {
		b.collectors[c.ID()] = c
	}
	return nil
}
