package libtime

import (
	"time"
)

const nanosToMillisDenominator = int64(time.Millisecond / time.Nanosecond)
const millisToSecsDenominator = int64(time.Second / time.Millisecond)

// ToMilliseconds returns time in milliseconds since epoch
func ToMilliseconds(t time.Time) int64 {
	return t.UnixNano() / nanosToMillisDenominator
}

// FromMilliseconds converts time in milliseconds since epoch to time.Time
func FromMilliseconds(ms int64) time.Time {
	seconds := ms / millisToSecsDenominator

	milliseconds := ms % millisToSecsDenominator
	nanoseconds := milliseconds * nanosToMillisDenominator

	return time.Unix(seconds, nanoseconds)
}

// DurationToMillis returns d in terms of milliseconds.
func DurationToMillis(d time.Duration) int64 {
	return d.Nanoseconds() / nanosToMillisDenominator
}
