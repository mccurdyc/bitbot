package main

func sellGreedy(funds, numBitcoins, price float64) (float64, float64) {
	funds += (numBitcoins * price)
	numBitcoins = 0

	return funds, numBitcoins
}

func sellRandom(funds, numBitcoins, price float64) (float64, float64) {
	numBitcoinsToSell := randomFloat64(0, numBitcoins)

	funds += (numBitcoinsToSell * price)
	numBitcoins -= numBitcoinsToSell
	return funds, numBitcoins
}
