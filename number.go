package aefire

import (
	"math/rand"
	"time"
)

func Round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}

func Random(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}
