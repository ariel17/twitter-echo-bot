package jobs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScheduler(t *testing.T) {
	count := 0
	f := func() error {
		time.Sleep(200 * time.Millisecond)
		count += 1
		return nil
	}

	s := NewScheduler(f, 100 * time.Millisecond)
	s.Start()
	assert.True(t, s.IsTicking())

	time.Sleep(time.Second)
	assert.Equal(t, 4, count)

	s.Stop()
	assert.False(t, s.IsTicking())
}