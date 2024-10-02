package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type Stock struct {
	ticker string
	ratio  uint8
	price  float32
	change float32
	date   string
}

func getStocks() []Stock {
	stockList := []Stock{}
	var userStock string
	var portfolioRatioTotal uint8 = 0
	for {
		fmt.Println("Enter a ticker or type 'done' to stop:")
		fmt.Scanln(&userStock)
		userStock = strings.TrimSpace(userStock)
		if userStock == "done" {
			break // Exit loop when done entering stocks and the ratios
		}
		// Apply the stock ticker to a "Stock" object
		newStock := Stock{
			ticker: userStock,
		}
		// Get the ratio of the stock to the portfolio
		for {
			var userRatio uint8
			fmt.Println("Now enter a uint8 portfolio ratio for the stock:")
			fmt.Scanln(&userRatio)
			if (portfolioRatioTotal+userRatio) > 100 || userRatio == 0 {
				fmt.Printf("Not a valid ratio, you have %d percent ratio left. Please try again.", 100-portfolioRatioTotal)
			} else {
				newStock.ratio = userRatio
				portfolioRatioTotal += userRatio
				break // Allows you to procede if your ratio of the stock is valid
			}
		}
		// Apply the ratio to the "Stock" object and increment the portfolioRatioTotal counter
		// Finally, append the stock to the "stocks" slice
		stockList = append(stockList, newStock)
		if portfolioRatioTotal == 100 {
			break // Exits once we have entered all our stock options
		}
	}
	return stockList
}

func getStockData(stockType *Stock) []Stock {
	// Use the colly package to fetch the stock data here.
	dataCollector := colly.NewCollector()
	stockDataList := []Stock{}
	/* You'll need to fill in the details based on how you're fetching the data.
	Below will be a loop to process all the data from collectedData and transform
	it into something useable by stockDataList. Process will probably be something like

	dataCollector.OnHTML("div#quote-header-info", func(e *colly.HTMLElement){
		var firststepStockData Stock
		firststepStockData.ticker =
		firststepStockData.price =
		firststepStockData.change =
		firststepStockData.date =

	})

	for index,

	where we have to get the stocks oldest data, when get the differnce between oldest
	time and now, and then get all the other monthly data in a loop and then append
	that data to stockDataList.
	*/
	// Once you've fetched the data, you can return the stockDataList
	return stockDataList
}

func worker(id int, wg *sync.WaitGroup, mu *sync.Mutex, operatedStock *Stock, listOfData [][]Stock) {
	defer wg.Done()
	// Fetch the stock data.
	fetchedStockData := getStockData(operatedStock)
	// Lock and Unlock data for appending
	mu.Lock()
	listOfData = append(listOfData, fetchedStockData)
	mu.Unlock()
}

func main() {
	// get user selected stocks
	stocks := getStocks()
	collectedStockData := [][]Stock{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for index, stock := range stocks {
		wg.Add(1)
		go worker(index, &wg, &mu, &stock, collectedStockData)
	}

}
