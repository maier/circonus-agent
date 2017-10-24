// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package wmi

import (
	"reflect"
	"testing"

	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
)

func TestNew(t *testing.T) {
	t.Log("Testing New")

	t.Log("no config")
	{
		p, err := New("")
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
		p, err := New("testdata/missing.json")
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

func TestCollect(t *testing.T) {
	t.Log("Testing Collect")

	t.Log("no config")
	{
		p, err := New("")
		if err == nil {
			t.Fatal("expected error")
		}
		if err != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
		if _, ok := p.(collector.Collector); !ok {
			t.Fatal("expected a collector.Collector interface")
		}

		cerr := p.Collect()
		if cerr == nil {
			t.Fatal("expected error")
		}
		if cerr != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
	}

	t.Log("config")
	{
		// does not matter if the file exists or not - this is a stub it should just react benignly
		p, err := New("testdata/missing.json")
		if err == nil {
			t.Fatal("expected error")
		}
		if err != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
		if _, ok := p.(collector.Collector); !ok {
			t.Fatal("expected a collector.Collector interface")
		}

		cerr := p.Collect()
		if cerr == nil {
			t.Fatal("expected error")
		}
		if cerr != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
	}
}

func TestFlush(t *testing.T) {
	t.Log("Testing Flush")

	t.Log("no config")
	{
		p, err := New("")
		if err == nil {
			t.Fatal("expected error")
		}
		if err != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
		if _, ok := p.(collector.Collector); !ok {
			t.Fatal("expected a collector.Collector interface")
		}

		metrics := p.Flush()
		if metrics == nil {
			t.Fatal("expected metrics")
		}
		if len(metrics) > 0 {
			t.Fatalf("expected empty metrics, got %v", metrics)
		}
	}

	t.Log("config")
	{
		// does not matter if the file exists or not - this is a stub it should just react benignly
		p, err := New("testdata/missing.json")
		if err == nil {
			t.Fatal("expected error")
		}
		if err != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
		if _, ok := p.(collector.Collector); !ok {
			t.Fatal("expected a collector.Collector interface")
		}

		metrics := p.Flush()
		if metrics == nil {
			t.Fatal("expected error")
		}
		if len(metrics) > 0 {
			t.Fatalf("expected empty metrics, got %v", metrics)
		}
	}
}

func TestID(t *testing.T) {
	t.Log("Testing ID")

	t.Log("no config")
	{
		p, err := New("")
		if err == nil {
			t.Fatal("expected error")
		}
		if err != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
		if _, ok := p.(collector.Collector); !ok {
			t.Fatal("expected a collector.Collector interface")
		}

		expect := "not_implemented"
		id := p.ID()
		if id != expect {
			t.Fatalf("expected (%s) got (%s)", expect, id)
		}
	}

	t.Log("config")
	{
		// does not matter if the file exists or not - this is a stub it should just react benignly
		p, err := New("testdata/missing.json")
		if err == nil {
			t.Fatal("expected error")
		}
		if err != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
		if _, ok := p.(collector.Collector); !ok {
			t.Fatal("expected a collector.Collector interface")
		}

		expect := "not_implemented"
		id := p.ID()
		if id != expect {
			t.Fatalf("expected (%s) got (%s)", expect, id)
		}
	}
}

func TestInventory(t *testing.T) {
	t.Log("Testing Inventory")

	t.Log("no config")
	{
		p, err := New("")
		if err == nil {
			t.Fatal("expected error")
		}
		if err != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
		if _, ok := p.(collector.Collector); !ok {
			t.Fatal("expected a collector.Collector interface")
		}

		expect := "InventoryStats"
		inventory := p.Inventory()
		if it := reflect.TypeOf(inventory).Name(); it != expect {
			t.Fatalf("expected (%s) got (%s)", expect, it)
		}
	}

	t.Log("config")
	{
		// does not matter if the file exists or not - this is a stub it should just react benignly
		p, err := New("testdata/missing.json")
		if err == nil {
			t.Fatal("expected error")
		}
		if err != collector.ErrNotImplemented {
			t.Fatalf("expected a collector.ErrNotImplemented error, got (%s)", err)
		}
		if _, ok := p.(collector.Collector); !ok {
			t.Fatal("expected a collector.Collector interface")
		}

		expect := "InventoryStats"
		inventory := p.Inventory()
		if it := reflect.TypeOf(inventory).Name(); it != expect {
			t.Fatalf("expected (%s) got (%s)", expect, it)
		}
	}
}
