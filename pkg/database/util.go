package database

import (
	"math"
	"time"
)

func getExponentialDelay(attemptNumber uint8) time.Duration {
	return time.Second * time.Duration(math.Pow(float64(attemptNumber), 2))
}
