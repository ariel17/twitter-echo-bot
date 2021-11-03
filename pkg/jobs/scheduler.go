package jobs

import (
	"log"
	"time"

	"github.com/ariel17/twitter-echo-bot/pkg/configs"
)

// Scheduler manages the ticket object and fires/stops async jobs in background.
type Scheduler struct {
	ticker *time.Ticker
	done   chan bool
	f func() error
}

// NewScheduler creates a new instance of scheduler with an specific f function
// to tick.
func NewScheduler() *Scheduler {
	return &Scheduler{
		ticker: time.NewTicker(configs.GetJobSeconds()),
		done:   make(chan bool),
		f: tick,
	}
}

func (s *Scheduler) Start() {
	go func() {
		for {
			select {
			case <-s.done:
				return
			case <-s.ticker.C:
				if err := s.f(); err != nil {
					panic(err)
				}
			}
		}
	}()
}

func (s *Scheduler) Stop() {
	s.ticker.Stop()
	s.done <- true
}

func tick() error {
	if err := answer(); err != nil {
		log.Fatalf("answer method failed: %v", err)
		return err
	}
	return nil
}