package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	n := 2425 // oldest date that I found to work for simulating based on historical data
	var large, small int
	var funds float64

	flag.IntVar(&large, "large", 1234, "The date range containing more dates from the start date (e.g., 200 versus 50)")
	flag.IntVar(&small, "small", 1234, "The date range containing fewer dates from the start date (e.g., 50 versus 200)")
	flag.Float64Var(&funds, "funds", 1234.12, "The available funding.")
	flag.Parse()

	fmt.Printf("Starting Funds: %f\n", funds)
	fmt.Printf("Large Moving Average Using: %d\n", large)
	fmt.Printf("Small Moving Average Using: %d\n", small)

	numBitcoins := 0.0 // always start with zero bitcoins

	greedyFunds := funds
	randomFunds := funds
	greedyBitcoins := numBitcoins
	randomBitcoins := numBitcoins
	greedyWorth := 0.0
	randomWorth := 0.0

	// make sure that the simulated date is larger than the large MA size
	for n > large {

		// get current simulated date
		currentDate := getDate(n)
		price := getPriceOnDate(currentDate)

		fmt.Printf("\nCurrent Date: %s\n", currentDate)
		fmt.Printf("Current Price: %f\n", price)

		// go back 'm' number of days and get the date
		endLarge := getDate(n - large) // larger date (older)
		endSmall := getDate(n - small) // smaller date (not as old)

		// gather the data for that date range
		dataLarge := getHistoricalData(currentDate, endLarge) // larger date range
		dataSmall := getHistoricalData(currentDate, endSmall) // smaller date range

		// calculate the moving average for the given date range
		smaLarge := calculateSMA(dataLarge) // larger date range MA
		smaSmall := calculateSMA(dataSmall) // smaller date range MA

		fmt.Printf("SMA Large: %f\n", smaLarge)
		fmt.Printf("SMA Small: %f\n", smaSmall)

		if smaLarge > smaSmall {

			greedyFunds, greedyBitcoins = buyGreedy(greedyFunds, greedyBitcoins, price)
			randomFunds, randomBitcoins = buyRandom(randomFunds, randomBitcoins, price)

			fmt.Println("Buying")
			fmt.Printf("GREEDY: Funds: %f, Bitcoins: %f\n", greedyFunds, greedyBitcoins)
			fmt.Printf("RANDOM: Funds: %f, Bitcoins: %f\n", randomFunds, randomBitcoins)

		} else if smaSmall > smaLarge {

			greedyFunds, greedyBitcoins = sellGreedy(greedyFunds, greedyBitcoins, price)
			randomFunds, randomBitcoins = sellRandom(randomFunds, randomBitcoins, price)

			fmt.Println("Selling")
			fmt.Printf("GREEDY: Funds: %f, Bitcoins: %f\n", greedyFunds, greedyBitcoins)
			fmt.Printf("RANDOM: Funds: %f, Bitcoins: %f\n", randomFunds, randomBitcoins)

		} else {
			fmt.Println("Nothing")
		}

		greedyWorth = calculateWorth(greedyFunds, greedyBitcoins, price)
		randomWorth = calculateWorth(randomFunds, randomBitcoins, price)

		fmt.Printf("Greedy Current Worth: %f\n", greedyWorth)
		fmt.Printf("Random Current Worth: %f\n", randomWorth)

		n--
	}
}

func getDate(n int) string {
	today := time.Now().Local()
	return today.AddDate(0, 0, -n).Format("2006-01-02")
}
