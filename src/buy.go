package main

func buyGreedy(a *Algorithm, price float64) {
	numBitcoinsToBuy := a.Funds / price

	a.Funds -= (numBitcoinsToBuy * price)
	a.Bitcoins += numBitcoinsToBuy
}

func buyRandom(a *Algorithm, price float64) {
	numBitcoinsCanBuy := a.Funds / price
	numBitcoinsToBuy := randomFloat64(0, numBitcoinsCanBuy)

	a.Funds -= (numBitcoinsToBuy * price)
	a.Bitcoins += numBitcoinsToBuy
}
