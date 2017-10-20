// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package procfs

import (
	"testing"

	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
)

func TestNewCPUMetrics(t *testing.T) {
	t.Log("Testing NewCPUMetrics")

	t.Log("no config")
	{
		p, err := NewCPUMetrics("")
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
		p, err := NewCPUMetrics("testdata/missing.json")
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
