// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package procfs

import (
	"sync"
	"time"

	cgm "github.com/circonus-labs/circonus-gometrics"
	"github.com/rs/zerolog"
)

// pfscommon defines ProcFS metrics common elements
type pfscommon struct {
	file                string // may be overriden in config file
	id                  string
	lastEnd             time.Time
	lastError           error
	lastMetrics         cgm.Metrics
	lastRunDuration     time.Duration
	lastStart           time.Time
	logger              zerolog.Logger
	metricStatus        map[string]bool // list of metrics and whether they should be collected or not, may be overriden in config file
	metricStatusDefault string          // default collection status for metrics NOT explicitly in metricStatus, may be overriden in config file
	running             bool
	runTTL              time.Duration // may be overriden in config file (default is for every request)
	sync.Mutex
}
