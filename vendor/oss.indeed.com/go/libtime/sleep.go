package libtime

import (
	"time"
)

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock -g -i Sleeper -o libtimetest/ -s _mock.go

// A Sleeper is a useful way for calling time.Sleep
// in a mockable way for tests.
type Sleeper interface {
	// Sleep for the specified amount of time.
	Sleep(time.Duration)
}

type sleeper struct{}

var _ Sleeper = (*sleeper)(nil)

// NewSleeper creates a Sleeper that will actually call
// time.Sleep under the hood, causing the program to sleep.
func NewSleeper() Sleeper {
	return (*sleeper)(nil) // lolz
}

func (s *sleeper) Sleep(duration time.Duration) {
	time.Sleep(duration)
}
