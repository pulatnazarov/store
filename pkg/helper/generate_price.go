package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomPrice(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}
