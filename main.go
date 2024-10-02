package main

import (
	//"FIRECALC_REV2/worker"
	"FIRECALC_REV2/stock"
	"fmt"
	"strings"
	"sync"
)

func getStocks() []stock.Stock {
	stockList := []stock.Stock{}
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
		newStock := stock.Stock{
			Ticker: userStock,
		}
		// Get the ratio of the stock to the portfolio
		for {
			var userRatio uint8
			fmt.Println("Now enter a uint8 portfolio ratio for the stock:")
			fmt.Scanln(&userRatio)
			if (portfolioRatioTotal+userRatio) > 100 || userRatio == 0 {
				fmt.Printf("Not a valid ratio, you have %d percent ratio left. Please try again.", 100-portfolioRatioTotal)
			} else {
				newStock.Ratio = userRatio
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

func getStockData(stockType *stock.Stock) []stock.Stock {
	stockDataList := []stock.Stock{}
	var wg sync.WaitGroup

	for index, temp := range temp {
		workedStock := worker.DataGetter

	}

	return stockDataList
}

func workStockData(id int, wg *sync.WaitGroup, mu *sync.Mutex, operatedStock *stock.Stock, listOfData [][]stock.Stock) {
	defer mu.Unlock()
	defer wg.Done()
	// Fetch the stock data.
	fetchedStockData := getStockData(operatedStock)
	// Lock and Unlock data for appending
	mu.Lock()
	listOfData = append(listOfData, fetchedStockData)
}

func main() {
	// get user selected stocks
	stocks := getStocks()
	collectedStockData := [][]stock.Stock{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for index, stock := range stocks {
		wg.Add(1)
		go workStockData(index, &wg, &mu, &stock, collectedStockData)
	}

}
