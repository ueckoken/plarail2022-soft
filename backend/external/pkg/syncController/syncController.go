package syncController

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"ueckoken/plarail2021-soft-external/pkg/envStore"
	"ueckoken/plarail2021-soft-external/pkg/servo"
)

type StationState struct {
	servo.StationState
}

type stationKVS struct {
	stations []StationState
	mtx      sync.Mutex
}

func (skvs *stationKVS) update(u StationState) {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	for i, s := range skvs.stations {
		if s.StationID == u.StationID {
			skvs.stations[i].State = u.State
			return
		}
	}
	skvs.stations = append(skvs.stations, u)
	return
}
func (skvs *stationKVS) get(stationID int32) (station StationState, err error) {
	skvs.mtx.Lock()
	defer skvs.mtx.Unlock()
	for _, s := range skvs.stations {
		if s.StationID == stationID {
			return s, nil
		}
	}
	return StationState{}, errors.New("Not found")
}

func (skvs *stationKVS) retrieve() []StationState {
	return skvs.stations
}

type SyncController struct {
	ClientHandler2syncController chan StationState
	SyncController2clientHandler chan StationState
	Environment                  *envStore.Env
}

func (s *SyncController) StartSyncController() {
	var kvs stationKVS
	go s.periodicallySync(&kvs)
	go s.triggeredSync(s.Environment, &kvs)
	for {
	}
}

func (s *SyncController) triggeredSync(e *envStore.Env, kvs *stationKVS) {
	for c := range s.ClientHandler2syncController {
		kvs.update(c)
		c2i := servo.NewCommand2Internal(c.StationState, e)
		err := c2i.Send()
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@:", err)
		s.SyncController2clientHandler <- c
	}
}

func (s *SyncController) periodicallySync(kvs *stationKVS) {
	ch := time.Tick(2 * time.Second)
	for range ch {
		fmt.Println("lockig")
		kvs.mtx.Lock()
		fmt.Println("locked")
		for _, st := range kvs.retrieve() {
			s.SyncController2clientHandler <- st
		}
		kvs.mtx.Unlock()
	}
}
