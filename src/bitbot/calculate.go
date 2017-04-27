package main

import (
	"math/rand"
	"time"

	"github.com/Jeffail/gabs"
)

func calculateSMA(bpis []*gabs.Container) float64 {
	var sum float64
	for _, bpi := range bpis {
		sum += bpi.Data().(float64)
	}
	return (sum / float64(len(bpis)))
}

func calculateWorth(a *Algorithm, price float64) {
	a.Worth = a.Funds + (a.Bitcoins * price)
}

func randomFloat64(min, max float64) float64 {
	rand.Seed(time.Now().UTC().UnixNano()) // otherwise always same random values
	// return a random value between the min value and the max
	return (rand.Float64() * max) + min
}
