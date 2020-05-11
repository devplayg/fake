package fake

import (
	"math/rand"
	"time"
)

const commonJitter = 3000 * time.Millisecond

func Now() time.Time {
	return dateWithJitter(time.Now(), commonJitter)
}

func Date(t time.Time) time.Time {
	return dateWithJitter(t, commonJitter)
}

func DateWithJitter(t time.Time, jitter time.Duration) time.Time {
	return dateWithJitter(t, jitter)
}

func DateRange(min, max time.Time) time.Time {
	if min.Equal(max) || min.After(max) {
		return min
	}
	diff := max.Sub(min).Milliseconds()
	jitter := rand.Int63n(diff)
	return min.Add(time.Duration(jitter) * time.Millisecond)
}
