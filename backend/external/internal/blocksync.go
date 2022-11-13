package internal

import (
	"go.uber.org/zap"
	"ueckoken/plarail2022-external/pkg/synccontroller"
	"ueckoken/plarail2022-external/spec"
)

// StartStationSync starts sync controller for station state.
func startBlockSync(logger *zap.Logger, syncInput chan synccontroller.KV[spec.Blocks_BlockId, spec.NotifyStateRequest_State], syncOutput chan<- synccontroller.KV[spec.Blocks_BlockId, spec.NotifyStateRequest_State]) {
	s := synccontroller.NewSyncController(logger, syncInput, syncOutput)

	go s.Run()
}
