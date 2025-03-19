package libtime

import (
	"time"
)

// StopTimerFunc is used to stop a time.Timer.
//
// Calling StopTimerFunc prevents its time.Timer from firing. Returns true if the call
// stops the timer, false if the timer has already expired. or has been stopped.
//
// https://pkg.go.dev/time#Timer.Stop
type StopTimerFunc func() bool

// SafeTimer creates a time.Timer and a StopTimerFunc, forcing the caller to deal
// with the otherwise potential resource leak. Encourages safe use of a time.Timer
// in a select statement, but without the overhead of a context.Context.
//
// Typical usage:
//
//    t, stop := libtime.SafeTimer(interval)
//    defer stop()
//    for {
//      select {
//        case <- t.C:
//          foo()
//        case <- otherC :
//          return
//      }
//    }
//
// Does not panic if duration is <= 0, instead assuming the smallest positive value.
func SafeTimer(duration time.Duration) (*time.Timer, StopTimerFunc) {
	t := time.NewTimer(duration)
	return t, t.Stop
}
