package jobs

import (
	"log"
	"sync"
	"time"

	"github.com/ariel17/twitter-echo-bot/pkg/configs"
)

// Scheduler manages the ticket object and fires/stops async jobs in background.
type Scheduler struct {
	ticker *time.Ticker
	done   chan bool
	f func() error
	isTicking bool
	m  sync.Mutex
}

// NewDefaultScheduler creates a new instance of scheduler with an specific f function
// to tick.
func NewDefaultScheduler() *Scheduler {
	return &Scheduler{
		ticker: time.NewTicker(configs.GetJobSeconds()),
		done:   make(chan bool),
		f: tick,
	}
}

// NewScheduler creates a generic scheduler with ticker. What to do and in how
// many times is indicated through parameters.
func NewScheduler(f func() error, tickDuration time.Duration) *Scheduler {
	return &Scheduler{
		ticker: time.NewTicker(tickDuration),
		done:   make(chan bool),
		f: f,
	}
}

func (s *Scheduler) Start() {
	s.m.Lock()
	s.isTicking = true
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
	s.m.Unlock()
}

func (s *Scheduler) Stop() {
	s.m.Lock()
	s.ticker.Stop()
	s.done <- true
	s.isTicking = false
	s.m.Unlock()
}

func (s *Scheduler) IsTicking() bool {
	return s.isTicking
}

func tick() error {
	if err := answer(); err != nil {
		log.Fatalf("answer method failed: %v", err)
		return err
	}
	return nil
}