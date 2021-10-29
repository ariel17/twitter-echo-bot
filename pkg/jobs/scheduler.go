package jobs

import (
	"log"
	"time"

	"github.com/ariel17/twitter-echo-bot/pkg/configs"
)

// Scheduler manages the ticket object and fires/stops async jobs in background.
type Scheduler interface {
	Start()
	Stop()
}

// NewScheduler creates a new instance of concrete implementation of scheduler.
func NewScheduler() Scheduler {
	return &scheduler{
		ticker: time.NewTicker(configs.GetJobSeconds()),
		done:   make(chan bool),
	}
}

type scheduler struct {
	ticker *time.Ticker
	done   chan bool
}

func (s *scheduler) Start() {
	go func() {
		for {
			select {
			case <-s.done:
				return
			case <-s.ticker.C:
				log.Printf("")
				if err := answer(); err != nil {
					log.Fatalf("answer method failed: %v", err)
					panic(err)
				}
			}
		}
	}()
}

func (s *scheduler) Stop() {
	s.ticker.Stop()
	s.done <- true
}
