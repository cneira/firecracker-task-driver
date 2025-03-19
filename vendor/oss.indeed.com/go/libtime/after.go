package libtime

import (
	"context"
	"time"
)

// After is a safer, but more expensive alternative to time.After.
//
// After waits for the duration to elapse and then sends the current time
// on the returned channel. If the context gets canceled before the duration
// elapses, no message is sent on the returned channel.
//
// The returned channel is never closed.
func After(ctx context.Context, duration time.Duration) <-chan time.Time {
	durationElapsedCh := make(chan time.Time, 1)
	go func() {
		t := time.NewTimer(duration)
		defer t.Stop()

		select {
		case tick := <-t.C:
			durationElapsedCh <- tick
		case <-ctx.Done():
		}
	}()
	return durationElapsedCh
}
