package fake

import (
	"math/rand"
	"time"
)

const commonJitter = 3000 * time.Millisecond

// Now returns time around 3000ms based on the current time
func Now() time.Time {
	return dateWithJitter(time.Now(), commonJitter)
}

// Date returns time around 3000ms based on specified time
func Date(t time.Time) time.Time {
	return dateWithJitter(t, commonJitter)
}

// DateWithJitter returns time around 'jitter' based on specified time
func DateWithJitter(t time.Time, jitter time.Duration) time.Time {
	return dateWithJitter(t, jitter)
}

// DateRange returns time between min and max
func DateRange(min, max time.Time) time.Time {
	if min.Equal(max) || min.After(max) {
		return min
	}
	diff := max.Sub(min).Milliseconds()
	jitter := rand.Int63n(diff)
	return min.Add(time.Duration(jitter) * time.Millisecond)
}

// SleepMillis sleeps at random times.
func SleepMillis(min, max uint32) {
	if min >= max {
		<-time.After(time.Duration(min) * time.Millisecond)
	}
	<-time.After(time.Duration(Uint32Range(min, max)) * time.Millisecond)
}

// TimeMillisAfter waits for the duration between min and max
// to elapse and then sends the current time
// on the returned channel.
func TimeMillisAfter(min, max uint32) <-chan time.Time {
	if min >= max {
		return time.After(time.Duration(min) * time.Millisecond)
	}

	return time.After(time.Duration(Uint32Range(min, max)) * time.Millisecond)
}
