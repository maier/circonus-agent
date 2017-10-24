// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build !windows

package wmi

import (
	"testing"

	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
)

func TestNewCPUCollector(t *testing.T) {
	t.Log("Testing NewCPUCollector")

	t.Log("no config")
	{
		p, err := NewCPUCollector("")
		if err == nil {
			t.Fatal("expected error")
		}
		if err != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
		if _, ok := p.(collector.Collector); !ok {
			t.Fatal("expected a collector.Collector interface")
		}
	}

	t.Log("config")
	{
		// does not matter if the file exists or not - this is a stub it should just react benignly
		p, err := NewCPUCollector("testdata/missing.json")
		if err == nil {
			t.Fatal("expected error")
		}
		if err != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
		if _, ok := p.(collector.Collector); !ok {
			t.Fatal("expected a collector.Collector interface")
		}
	}
}
