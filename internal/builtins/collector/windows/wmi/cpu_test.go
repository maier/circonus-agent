// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

import (
	"path/filepath"
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
		_, err := NewCPUCollector(filepath.Join("testdata", "missing"))
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}
	}

	t.Log("config (bad syntax)")
	{
		_, err := NewCPUCollector(filepath.Join("testdata", "bad_syntax"))
		if err == nil {
			t.Fatal("expected error")
		}
	}

	t.Log("config (config no settings)")
	{
		c, err := NewCPUCollector(filepath.Join("testdata", "config_no_settings"))
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}
		if c == nil {
			t.Fatal("expected no nil")
		}
	}

	t.Log("config (id setting)")
	{
		c, err := NewCPUCollector(filepath.Join("testdata", "config_id_setting"))
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}
		if c.(*CPU).id != "foo" {
			t.Fatalf("expected foo, got (%s)", c.ID())
		}
	}

	t.Log("config (report all cpus setting)")
	{
		c, err := NewCPUCollector(filepath.Join("testdata", "config_report_all_cpus_setting"))
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}
		if c.(*CPU).reportAllCPUs {
			t.Fatal("expected false")
		}
	}

	t.Log("config (metrics enabled setting)")
	{
		c, err := NewCPUCollector(filepath.Join("testdata", "config_metrics_enabled_setting"))
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}
		if len(c.(*CPU).metricStatus) == 0 {
			t.Fatalf("expected >0 metric status settings, got (%#v)", c.(*CPU).metricStatus)
		}
		enabled, ok := c.(*CPU).metricStatus["foo"]
		if !ok {
			t.Fatalf("expected 'foo' key in metric status settings, got (%#v)", c.(*CPU).metricStatus)
		}
		if !enabled {
			t.Fatalf("expected 'foo' to be enabled in metric status settings, got (%#v)", c.(*CPU).metricStatus)
		}
	}

	t.Log("config (metrics disabled setting)")
	{
		c, err := NewCPUCollector(filepath.Join("testdata", "config_metrics_disabled_setting"))
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}
		if len(c.(*CPU).metricStatus) == 0 {
			t.Fatalf("expected >0 metric status settings, got (%#v)", c.(*CPU).metricStatus)
		}
		enabled, ok := c.(*CPU).metricStatus["foo"]
		if !ok {
			t.Fatalf("expected 'foo' key in metric status settings, got (%#v)", c.(*CPU).metricStatus)
		}
		if enabled {
			t.Fatalf("expected 'foo' to be disabled in metric status settings, got (%#v)", c.(*CPU).metricStatus)
		}
	}

	t.Log("config (metrics disabled setting)")
	{
		c, err := NewCPUCollector(filepath.Join("testdata", "config_metrics_disabled_setting"))
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}
		if len(c.(*CPU).metricStatus) == 0 {
			t.Fatalf("expected >0 metric status settings, got (%#v)", c.(*CPU).metricStatus)
		}
		enabled, ok := c.(*CPU).metricStatus["foo"]
		if !ok {
			t.Fatalf("expected 'foo' key in metric status settings, got (%#v)", c.(*CPU).metricStatus)
		}
		if enabled {
			t.Fatalf("expected 'foo' to be disabled in metric status settings, got (%#v)", c.(*CPU).metricStatus)
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

	t.Log("config (config id setting)")
	{
		c, err := NewCPUCollector(filepath.Join("testdata", "config_id_setting"))
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		expect := "foo"
		id := c.ID()
		if id != expect {
			t.Fatalf("expected (%s) got (%s)", expect, id)
		}
	}
}

func TestCPUFlush(t *testing.T) {
	t.Log("Testing Flush")

	t.Log("no config")
	{
		c, err := NewCPUCollector("")
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		metrics := c.Flush()
		if metrics == nil {
			t.Fatal("expected metrics")
		}
		if len(metrics) > 0 {
			t.Fatalf("expected empty metrics, got %v", metrics)
		}
	}

	t.Log("config (config.json)")
	{
		c, err := NewCPUCollector(filepath.Join("testdata", "config"))
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		metrics := c.Flush()
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
		c, err := NewCPUCollector("")
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		if err := c.Collect(); err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}
	}

	t.Log("config (config.json)")
	{
		c, err := NewCPUCollector(filepath.Join("testdata", "config"))
		if err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		if err := c.Collect(); err != nil {
			t.Fatalf("expected NO error, got (%s)", err)
		}

		metrics := c.Flush()
		if metrics == nil {
			t.Fatal("expected error")
		}
		if len(metrics) == 0 {
			t.Fatalf("expected metrics, got %v", metrics)
		}
	}
}
