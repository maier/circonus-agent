// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package reverse

import (
	"context"
	"sync"
	"time"

	"github.com/circonus-labs/circonus-agent/internal/check"
	"github.com/circonus-labs/circonus-agent/internal/config"
	"github.com/circonus-labs/circonus-agent/internal/reverse/connection"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Reverse struct {
	agentAddress string
	configs      *check.ReverseConfigs
	checkBundle  *check.Check
	enabled      bool
	logger       zerolog.Logger
}

func New(parentLogger zerolog.Logger, checkBundle *check.Check, agentAddress string) (*Reverse, error) {
	if checkBundle == nil {
		return nil, errors.New("invalid checkBundle (nil")
	}
	if agentAddress == "" {
		return nil, errors.New("invalid agent address (empty)")
	}

	r := &Reverse{
		agentAddress: agentAddress,
		checkBundle:  checkBundle,
		enabled:      viper.GetBool(config.KeyReverse),
	}

	cfgs, err := r.checkBundle.GetReverseConfigs()
	if err != nil {
		return nil, errors.Wrap(err, "getting reverse configurations")
	}
	r.configs = cfgs

	r.logger = parentLogger.With().Str("pkg", "reverse").Str("cid", viper.GetString(config.KeyCheckBundleID)).Logger()

	return r, nil
}

// Start reverse connection(s) to the broker(s)
func (r *Reverse) Start(ctx context.Context) error {
	if !r.enabled {
		r.logger.Info().Msg("disabled, not starting")
		return nil
	}
	if r.configs == nil {
		return errors.New("invalid reverse configurations (nil)")
	}
	if len(*r.configs) == 0 {
		return errors.New("invalid reverse configurations (zero)")
	}

	lastRefresh := time.Now()
	refreshCheck := false
	rctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for {
		select {
		case <-rctx.Done():
			return nil
		default:
		}

		if time.Since(lastRefresh) > 5*time.Minute {
			refreshCheck = true
		}

		if refreshCheck {
			if err := r.checkBundle.RefreshCheckConfig(); err != nil {
				if cberr, ok := errors.Cause(err).(*check.BundleNotActiveError); ok {
					r.logger.Error().Err(cberr).Msg("exiting reverse")
					cancel()
					return err
				}
				r.logger.Error().Err(err).Msg("refreshing check")
				continue
			}

			cfgs, err := r.checkBundle.GetReverseConfigs()
			if err != nil {
				cancel()
				return errors.Wrap(err, "getting reverse configurations")
			}
			r.configs = cfgs
			refreshCheck = false
		}

		primaryCN, err := r.checkBundle.FindPrimaryBrokerInstance(r.configs)
		if err != nil {
			if nferr, ok := errors.Cause(err).(*check.NoOwnerFoundError); ok {
				r.logger.Warn().Err(nferr).Msg("refreshing check bundle configuration")
				refreshCheck = true
				continue
			}
			return err
		}

		cfg, ok := (*r.configs)[primaryCN]
		if !ok {
			r.logger.Warn().Str("primary", primaryCN).Msg("primary broker cn not found, refreshing check")
			refreshCheck = true
			continue
		}

		r.logger.Debug().Interface("config", cfg).Msg("reverse broker")
		rc, err := connection.New(r.logger, r.agentAddress, &cfg)
		if err != nil {
			cancel()
			return err
		}

		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			if err := rc.Start(rctx); err != nil {
				r.logger.Warn().Err(err).Msg("reverse connection")
				if cerr, ok := err.(*connection.OpError); ok {
					if cerr.Fatal {
						cancel()
					} else if cerr.RefreshCheck {
						refreshCheck = true
					}
				} // else just let it wrap around and find the owner again
			}
			wg.Done()
			return
		}()

		wg.Wait()
	}
}
