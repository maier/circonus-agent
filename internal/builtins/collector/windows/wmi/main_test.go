// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Log("Testing New")

	c, err := New()
	if err != nil {
		t.Fatalf("expected NO error, got (%s)", err)
	}
	if len(c) == 0 {
		t.Fatal("expected at least 1 collector.Collector")
	}
}
