package internal

import (
	"time"
	"ueckoken/plarail2022-external/pkg/envStore"
	"ueckoken/plarail2022-external/pkg/httphandler"
	"ueckoken/plarail2022-external/pkg/synccontroller"
	"ueckoken/plarail2022-external/spec"

	"go.uber.org/zap"
)

// Run runs external server.
func Run(logger *zap.Logger) {
	synccontrollerInput := make(chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State])
	synccontrollerOutput := make(chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State])
	grpcHandlerInput := make(chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State])
	main2grpcHandler := make(chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State])
	httpInputKV := make(chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State])
	httpInput := make(chan *spec.Command2InternalRequest)
	httpOutput := make(chan *spec.Command2InternalRequest)

	go func() {
		for c := range synccontrollerOutput {
			select {
			case main2grpcHandler <- c:
			default:
				logger.Info("buffer full", zap.String("buffer", "main2grpcHandler"))
			}
			select {
			case httpInputKV <- c:
			default:
				logger.Info("buffer full", zap.String("buffer", "httpInputKV"))
			}
		}
	}()
	go func() {
		for c := range httpInputKV {
			httpInput <- &spec.Command2InternalRequest{Station: &spec.Stations{StationId: c.Key}, State: c.Value}
		}
	}()

	go func() {
		for c := range httpOutput {
			synccontrollerInput <- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]{Key: c.GetStation().GetStationId(), Value: c.GetState()}
		}
	}()

	envVal := envStore.GetEnv()

	httpServer := httphandler.NewHTTPServer(
		logger.Named("sync-controller"),
		httpOutput,
		httpInput,
		envVal,
	)
	logger.Info("1")
	StartStationSync(logger.Named("station-sync"), synccontrollerInput, synccontrollerOutput)
	logger.Info("2")
	grpcHandler := NewGrpcHandler(logger.Named("grpc-handler"), envVal, main2grpcHandler, grpcHandlerInput)

	logger.Info("3")

	client2blocksync := make(chan synccontroller.KV[spec.Blocks_BlockId, spec.NotifyStateRequest_State])
	blocksync2client := make(chan synccontroller.KV[spec.Blocks_BlockId, spec.NotifyStateRequest_State])
	startBlockSync(logger.Named("blocksync"), client2blocksync, blocksync2client)
	logger.Info("4")
	grpcBlockHandl := NewGrpcBlockHandler(logger.Named("grpc-block-handler"), envVal, client2blocksync, blocksync2client)
	logger.Info("5")

	go GRPCListenAndServe(logger, uint(envVal.ClientSideServer.GrpcPort), grpcHandler, grpcBlockHandl)
	go httpServer.StartServer()
	for {
		time.Sleep(100 * time.Second)
	}
}
