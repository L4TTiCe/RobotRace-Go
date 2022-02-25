package stadium

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type racer interface {
	Start()
	RegisterHost(stadium *Stadium)
	GetCommandString() string
	GetDelay() time.Duration
	Announce()
}

type Stadium struct {
	racers   []racer
	finished []racer

	rank Counter
	wg   sync.WaitGroup
}

func (stadium *Stadium) String() string {
	return fmt.Sprintf("Stadium")
}

func NewStadium() *Stadium {
	stadium := Stadium{
		racers:   []racer{},
		finished: []racer{},
		rank:     Counter{},
	}
	stadium.init()
	return &stadium
}

func (stadium *Stadium) init() {
	log.Printf("[%s]\tReady to Host\n",
		stadium,
	)
}

func (stadium *Stadium) done() {
	log.Printf("[%s]\tDone\n",
		stadium,
	)
}

func (stadium *Stadium) AddRacer(racer racer) {
	racer.RegisterHost(stadium)
	stadium.racers = append(stadium.racers, racer)
	log.Printf("[%s]\t%s enters with Commands:\t%s\t(Delay: %dms)\n",
		stadium,
		racer,
		racer.GetCommandString(),
		racer.GetDelay().Milliseconds(),
	)
}

func (stadium *Stadium) GetRank(racer racer) int {
	stadium.wg.Done()
	stadium.finished = append(stadium.finished, racer)
	return stadium.rank.GetRank()
}

func (stadium *Stadium) Announce() {
	log.Printf("[%s]\tAnnouncing Results\n", stadium)
	for _, racer := range stadium.finished {
		racer.Announce()
	}
}

func (stadium *Stadium) StartRace() {
	log.Printf("[%s]\tStarting Race\n", stadium)
	timestamp := time.Now()
	defer stadium.done()

	for _, racer := range stadium.racers {
		stadium.wg.Add(1)
		go racer.Start()
	}
	stadium.wg.Wait()
	log.Printf("[%s]\tThe Race has Concluded\t(took %dms)\n",
		stadium,
		time.Since(timestamp).Milliseconds(),
	)
	stadium.Announce()
}
