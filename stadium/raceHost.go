package stadium

import (
	"sync"
)

type racer interface {
	Start()
	RegisterHost(stadium *Stadium)
}

type Stadium struct {
	racers []racer
	rank   Counter
	wg     sync.WaitGroup
}

func NewStadium() Stadium {
	return Stadium{
		racers: []racer{},
		rank:   Counter{},
	}
}

func (host *Stadium) AddRacer(racer racer) {
	racer.RegisterHost(host)
	host.racers = append(host.racers, racer)
}

func (host *Stadium) GetRank() int {
	host.wg.Done()
	return host.rank.GetRank()
}

func (host *Stadium) StartRace() {
	for _, racer := range host.racers {
		host.wg.Add(1)
		go racer.Start()
	}
	host.wg.Wait()
}
