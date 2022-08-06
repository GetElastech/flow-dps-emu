package engine

import (
	"fmt"
	"github.com/onflow/flow-dps/api/dps"
	"github.com/rs/zerolog"
)

// New returns a new RPC engine.
func NewDPS(log zerolog.Logger, config Config, dpsServer dps.APIServer) (*RPC, error) {
	if dpsServer == nil {
		return nil, fmt.Errorf("proxy not set")
	}

	eng, err := New(log, config)
	if err != nil {
		return nil, err
	}

	dps.RegisterAPIServer(eng.server, dpsServer)

	return eng, nil
}
