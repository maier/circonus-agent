// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build linux

package procfs

import (
	"fmt"
	"path"
	"regexp"
	"runtime"
	"strings"

	"github.com/circonus-labs/circonus-agent/internal/builtins/collector"
	"github.com/circonus-labs/circonus-agent/internal/config"
	"github.com/circonus-labs/circonus-agent/internal/config/defaults"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	CollectorPrefix     = "procfs/"
	PackageName         = "builtins.linux.procfs"
	ProcFSPath          = "/proc"
	NameCPU             = "cpu"
	NameDisk            = "disk"
	NameNetInterface    = "if"
	NameNetProto        = "proto"
	NameNetSocket       = "socket"
	NameLoad            = "load"
	NameVM              = "vm"
	metricNameSeparator = "`"        // character used to separate parts of metric names
	regexPat            = `^(?:%s)$` // fmt pattern used compile include/exclude regular expressions
)

var (
	defaultExcludeRegex = regexp.MustCompile(fmt.Sprintf(regexPat, ""))
	defaultIncludeRegex = regexp.MustCompile(fmt.Sprintf(regexPat, ".+"))
)

// New creates new ProcFS collector
func New() ([]collector.Collector, error) {
	none := []collector.Collector{}

	if runtime.GOOS != "linux" {
		return none, nil
	}

	l := log.With().Str("pkg", "builtins.procfs").Logger()

	enbledCollectors := viper.GetStringSlice(config.KeyCollectors)
	if len(enbledCollectors) == 0 {
		l.Info().Msg("no builtin collectors enabled")
		return none, nil
	}

	collectors := make([]collector.Collector, 0, len(enbledCollectors))
	initErrMsg := "initializing builtin collector"
	for _, name := range enbledCollectors {
		if !strings.HasPrefix(name, CollectorPrefix) {
			continue
		}
		name = strings.Replace(name, CollectorPrefix, "", -1)
		cfgBase := "procfs_" + name + "_collector"
		switch name {
		case NameCPU:
			c, err := NewCPUCollector(path.Join(defaults.EtcPath, cfgBase), ProcFSPath)
			if err != nil {
				l.Error().Str("name", name).Err(err).Msg(initErrMsg)
				continue
			}
			collectors = append(collectors, c)

		case NameDisk:
			c, err := NewDiskCollector(path.Join(defaults.EtcPath, cfgBase), ProcFSPath)
			if err != nil {
				l.Error().Str("name", name).Err(err).Msg(initErrMsg)
				continue
			}
			collectors = append(collectors, c)

		case NameNetInterface:
			c, err := NewNetIFCollector(path.Join(defaults.EtcPath, cfgBase), ProcFSPath)
			if err != nil {
				l.Error().Str("name", name).Err(err).Msg(initErrMsg)
				continue
			}
			collectors = append(collectors, c)

		case NameNetProto:
			c, err := NewNetProtoCollector(path.Join(defaults.EtcPath, cfgBase), ProcFSPath)
			if err != nil {
				l.Error().Str("name", name).Err(err).Msg(initErrMsg)
				continue
			}
			collectors = append(collectors, c)

		case NameNetSocket:
			c, err := NewNetSocketCollector(path.Join(defaults.EtcPath, cfgBase), ProcFSPath)
			if err != nil {
				l.Error().Str("name", name).Err(err).Msg(initErrMsg)
				continue
			}
			collectors = append(collectors, c)

		case NameLoad:
			c, err := NewLoadCollector(path.Join(defaults.EtcPath, cfgBase), ProcFSPath)
			if err != nil {
				l.Error().Str("name", name).Err(err).Msg(initErrMsg)
				continue
			}
			collectors = append(collectors, c)

		case NameVM:
			c, err := NewVMCollector(path.Join(defaults.EtcPath, cfgBase), ProcFSPath)
			if err != nil {
				l.Error().Str("name", name).Err(err).Msg(initErrMsg)
				continue
			}
			collectors = append(collectors, c)

		default:
			l.Warn().Str("name", name).Msg("unknown builtin collector, ignoring")
		}
	}

	return collectors, nil
}