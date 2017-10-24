// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build !windows

package wmi

import "testing"

func TestConfigure(t *testing.T) {
	t.Log("Testing configure")

	err := configure()
	if err != nil {
		t.Fatalf("expected NO error, got (%s)", err)
	}
}
