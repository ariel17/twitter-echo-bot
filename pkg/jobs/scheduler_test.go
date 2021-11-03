package jobs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScheduler(t *testing.T) {
	f := func() error {
		time.Sleep(200 * time.Millisecond)
		return nil
	}

	s := NewScheduler(f, 100 * time.Millisecond)
	s.Start()
	assert.True(t, s.IsTicking())

	time.Sleep(time.Second)

	s.Stop()
	assert.False(t, s.IsTicking())
}