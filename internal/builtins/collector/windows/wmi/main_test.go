// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Log("Testing New")

	c, err := New()
	if err == nil {
		t.Fatal("expected error")
	}
	if len(c) == 0 {
		t.Fatal("expected at least 1 collector.Collector")
	}
}

func TestCollect(t *testing.T) {
	t.Log("Testing Collect")

	c, err := New()
	if err == nil {
		t.Fatal("expected error")
	}
	if len(c) == 0 {
		t.Fatal("expected at least 1 collector.Collector")
	}

	cerr := c[0].Collect()
	if cerr == nil {
		t.Fatal("expected error")
	}
}

func TestFlush(t *testing.T) {
	t.Log("Testing Flush")

	c, err := New()
	if err == nil {
		t.Fatal("expected error")
	}
	if len(c) == 0 {
		t.Fatal("expected at least 1 collector.Collector")
	}

	metrics := c[0].Flush()
	if metrics == nil {
		t.Fatal("expected metrics")
	}
	if len(metrics) > 0 {
		t.Fatalf("expected empty metrics, got %v", metrics)
	}

}

func TestID(t *testing.T) {
	t.Log("Testing ID")

	c, err := New()
	if err == nil {
		t.Fatal("expected error")
	}
	if len(c) == 0 {
		t.Fatal("expected at least 1 collector.Collector")
	}

	expect := "not_implemented"
	id := c[0].ID()
	if id != expect {
		t.Fatalf("expected (%s) got (%s)", expect, id)
	}
}

func TestInventory(t *testing.T) {
	t.Log("Testing Inventory")

	c, err := New()
	if err == nil {
		t.Fatal("expected error")
	}
	if len(c) == 0 {
		t.Fatal("expected at least 1 collector.Collector")
	}

	expect := "InventoryStats"
	inventory := c[0].Inventory()
	if it := reflect.TypeOf(inventory).Name(); it != expect {
		t.Fatalf("expected (%s) got (%s)", expect, it)
	}

}
