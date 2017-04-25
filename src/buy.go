package main

func buyGreedy(funds, numBitcoins, price float64) (float64, float64) {
	numBitcoinsToBuy := funds / price

	funds -= (numBitcoinsToBuy * price)
	numBitcoins += numBitcoinsToBuy

	return funds, numBitcoins
}

func buyRandom(funds, numBitcoins, price float64) (float64, float64) {
	numBitcoinsCanBuy := funds / price
	numBitcoinsToBuy := randomFloat64(0, numBitcoinsCanBuy)

	funds -= (numBitcoinsToBuy * price)
	numBitcoins += numBitcoinsToBuy

	return funds, numBitcoins
}
