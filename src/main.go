package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Algorithm struct {
	Name     string
	Funds    float64
	Bitcoins float64
	Worth    float64
}

type DateInfo struct {
	Trial        int
	InitialFunds float64
	Small        int
	Large        int
	Date         string
	Ssma         float64
	Lsma         float64
	Algorithms   []Algorithm
}

func main() {
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

	// make sure that the simulated date is larger than the large MA size
	for trial := 1; trial <= 30; trial++ {

		dates := []DateInfo{}
		greedy := Algorithm{"Greedy", funds, numBitcoins, 0.0}
		random := Algorithm{"Random", funds, numBitcoins, 0.0}
		n := 30 // oldest date that I found to work for simulating based on historical data
		// n := 2425 // oldest date that I found to work for simulating based on historical data

		for n > 0 {

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

				buyGreedy(&greedy, price)
				buyRandom(&random, price)

				fmt.Println("Buying")
				fmt.Printf("GREEDY: Funds: %f, Bitcoins: %f\n", greedy.Funds, greedy.Bitcoins)
				fmt.Printf("RANDOM: Funds: %f, Bitcoins: %f\n", random.Funds, random.Bitcoins)

			} else if smaSmall > smaLarge {

				sellGreedy(&greedy, price)
				sellRandom(&random, price)

				fmt.Println("Selling")
				fmt.Printf("GREEDY: Funds: %f, Bitcoins: %f\n", greedy.Funds, greedy.Bitcoins)
				fmt.Printf("RANDOM: Funds: %f, Bitcoins: %f\n", random.Funds, random.Bitcoins)

			} else {
				fmt.Println("Nothing")
			}

			calculateWorth(&greedy, price)
			calculateWorth(&random, price)

			fmt.Printf("Greedy Current Worth: %f\n", greedy.Worth)
			fmt.Printf("Random Current Worth: %f\n", random.Worth)

			algoritms := []Algorithm{greedy, random}
			d := DateInfo{trial, funds, small, large, currentDate, smaSmall, smaLarge, algoritms}
			dates = append(dates, d)

			n--
		}
		writeToFile("../data/algorithm-data.csv", dates)
	}
}

func getDate(n int) string {
	today := time.Now().Local()
	return today.AddDate(0, 0, -n).Format("2006-01-02")
}

func writeToFile(file string, dates []DateInfo) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	// w.Write([]string{"Trial", "Initial_Funds", "Small", "Large", "Date", "Small_SMA", "Large_SMA", "A1_Name", "A1_Funds", "A1_Bitcoins", "A1_Worth", "A2_Name", "A2_Funds", "A2_Bitcoins", "A2_Worth"})
	for _, d := range dates {

		trial := strconv.Itoa(d.Trial)
		initialFunds := strconv.FormatFloat(d.InitialFunds, 'f', 2, 64)
		s := strconv.Itoa(d.Small)
		l := strconv.Itoa(d.Large)
		ssma := strconv.FormatFloat(d.Ssma, 'f', 6, 64)
		lsma := strconv.FormatFloat(d.Lsma, 'f', 6, 64)

		var arr = []string{trial, initialFunds, s, l, d.Date, ssma, lsma}

		for i := 0; i < len(d.Algorithms); i++ {
			funds := strconv.FormatFloat(d.Algorithms[i].Funds, 'f', 2, 64)
			bitcoins := strconv.FormatFloat(d.Algorithms[i].Bitcoins, 'f', 2, 64)
			worth := strconv.FormatFloat(d.Algorithms[i].Worth, 'f', 2, 64)

			arr = append(arr, d.Algorithms[i].Name, funds, bitcoins, worth)
		}

		err = w.Write(arr)
		if err != nil {
			panic(err)
		}
	}
	w.Flush()
}
