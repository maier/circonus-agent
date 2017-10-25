// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

import (
	"path"

	"github.com/StackExchange/wmi"
	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
	"github.com/circonus-labs/circonus-agent/internal/config/defaults"
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

	c, err := NewCPUCollector(path.Join(defaults.EtcPath, "cpu"))
	if err != nil {
		return []collector.Collector{}, errors.Wrap(err, "initializing wmi.cpu")
	}

	collectors = append(collectors, c)

	return collectors, nil
}
