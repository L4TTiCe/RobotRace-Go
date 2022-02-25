package stadium

import (
	"sync"
)

type racer interface {
	Start()
	RegisterHost(stadium *Stadium)
	Announce()
}

type Stadium struct {
	racers   []racer
	finished []racer

	rank Counter
	wg   sync.WaitGroup
}

func NewStadium() Stadium {
	return Stadium{
		racers:   []racer{},
		finished: []racer{},
		rank:     Counter{},
	}
}

func (host *Stadium) AddRacer(racer racer) {
	racer.RegisterHost(host)
	host.racers = append(host.racers, racer)
}

func (host *Stadium) GetRank(racer racer) int {
	host.wg.Done()
	host.finished = append(host.finished, racer)
	return host.rank.GetRank()
}

func (host *Stadium) Announce() {
	for _, racer := range host.finished {
		racer.Announce()
	}
}

func (host *Stadium) StartRace() {
	for _, racer := range host.racers {
		host.wg.Add(1)
		go racer.Start()
	}
	host.wg.Wait()
	host.Announce()
}
