package main

func sellGreedy(a *Algorithm, price float64) {
	a.Funds += (a.Bitcoins * price)
	a.Bitcoins = 0
}

func sellRandom(a *Algorithm, price float64) {
	numBitcoinsToSell := randomFloat64(0, a.Bitcoins)

	a.Funds += (numBitcoinsToSell * price)
	a.Bitcoins -= numBitcoinsToSell
}
