package generators

import (
	"math/rand"
	"time"
)

// GenerateRandomFutureDate creates a new date in the future between min and max days from now
func GenerateRandomFutureDate(minDays int, maxDays int) time.Time {
	randDays := rand.Intn(maxDays-minDays) + minDays

	return time.Now().AddDate(0, 0, randDays)
}

// GenerateRandomPastDate creates a new date in the past between min and max days from now
func GenerateRandomPastDate(minDays int, maxDays int) time.Time {
	randDays := -(rand.Intn(maxDays-minDays) + minDays)

	return time.Now().AddDate(0, 0, randDays)
}
