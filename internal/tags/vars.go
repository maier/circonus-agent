// Copyright © 2018 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package tags

import "regexp"

const (
	Delimiter       = ":"
	Separator       = ","
	replacementChar = "_"
)

var (
	valid   = regexp.MustCompile(`^[^:,]+:[^:,]+(,[^:,]+:[^:,]+)*$`)
	cleaner = regexp.MustCompile(`[\[\]'"` + "`]")
)
