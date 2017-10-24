// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

import (
	"testing"
)

func TestNewCPUCollector(t *testing.T) {
	t.Log("Testing NewCPUCollector")

	t.Log("no config")
	{
		_, err := NewCPUCollector("")
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}
	}

	t.Log("config (missing)")
	{
		_, err := NewCPUCollector("testdata/missing.json")
		if err == nil {
			t.Fatal("expected error")
		}
	}

	t.Log("config (bad syntax)")
	{
		_, err := NewCPUCollector("testdata/bad_syntax.json")
		if err == nil {
			t.Fatal("expected error")
		}
	}

	t.Log("config (config.json)")
	{
		p, err := NewCPUCollector("testdata/config.json")
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		if p.(*CPU).reportAllCPUs {
			t.Fatal("expected false")
		}
	}
}

func TestCPUID(t *testing.T) {
	t.Log("Testing ID")

	t.Log("no config")
	{
		p, err := NewCPUCollector("")
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		expect := "cpu"
		id := p.ID()
		if id != expect {
			t.Fatalf("expected (%s) got (%s)", expect, id)
		}
	}

	t.Log("config (missing)")
	{
		p, err := NewCPUCollector("testdata/missing.json")
		if err == nil {
			t.Fatal("expected error")
		}

		if p != nil {
			t.Fatalf("expected nil, got (%#v)", p)
		}
	}

	t.Log("config (config.json)")
	{
		p, err := NewCPUCollector("testdata/config.json")
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		expect := "cpu"
		id := p.ID()
		if id != expect {
			t.Fatalf("expected (%s) got (%s)", expect, id)
		}
	}
}

func TestCPUFlush(t *testing.T) {
	t.Log("Testing Flush")

	t.Log("no config")
	{
		p, err := NewCPUCollector("")
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		metrics := p.Flush()
		if metrics == nil {
			t.Fatal("expected metrics")
		}
		if len(metrics) > 0 {
			t.Fatalf("expected empty metrics, got %v", metrics)
		}
	}

	t.Log("config (missing)")
	{
		p, err := NewCPUCollector("testdata/missing.json")
		if err == nil {
			t.Fatal("expected error")
		}
		if p != nil {
			t.Fatalf("expected nil got (%#v)", p)
		}
	}

	t.Log("config (config.json)")
	{
		p, err := NewCPUCollector("testdata/config.json")
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
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

func TestCPUCollect(t *testing.T) {
	t.Log("Testing Collect")

	t.Log("no config")
	{
		p, err := NewCPUCollector("")
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		if err := p.Collect(); err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}
	}

	t.Log("config (missing)")
	{
		p, err := NewCPUCollector("testdata/missing.json")
		if err == nil {
			t.Fatal("expected error")
		}
		if p != nil {
			t.Fatalf("expected nil got (%#v)", p)
		}
	}

	t.Log("config (config.json)")
	{
		p, err := NewCPUCollector("testdata/config.json")
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		if err := p.Collect(); err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		metrics := p.Flush()
		if metrics == nil {
			t.Fatal("expected error")
		}
		if len(metrics) == 0 {
			t.Fatalf("expected metrics, got %v", metrics)
		}
	}
}
