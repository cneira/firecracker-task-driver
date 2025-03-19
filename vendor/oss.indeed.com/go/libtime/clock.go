package libtime // import "oss.indeed.com/go/libtime"

import "time"

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock -g -i Clock -o libtimetest/ -s _mock.go

// Clock provides a mock-able interface of some of the standard library time
// package functions, like time.Now and time.Since. A default implementation is
// provided by SystemClock which defers to the time package functions.
type Clock interface {
	// Now is the time on this Clock.
	Now() time.Time

	// Since returns the size of the interval between the time.Time of Clock.Now
	// and the provided time.Time.
	Since(time.Time) time.Duration

	// SinceMS returns the same thing as Since, but converted to milliseconds.
	SinceMS(time.Time) int
}

// SystemClock creates a Clock that is implemented by deferring to the
// standard library time package.
func SystemClock() Clock {
	return (*clock)(nil)
}

type clock struct {
	// nothing to see here
}

var _ Clock = (*clock)(nil)

func (*clock) Now() time.Time {
	return time.Now()
}

func (c *clock) Since(t time.Time) time.Duration {
	now := c.Now()
	return now.Sub(t)
}

func (c *clock) SinceMS(t time.Time) int {
	dur := c.Since(t)
	millis := DurationToMillis(dur)
	i := int(millis)
	return i
}
