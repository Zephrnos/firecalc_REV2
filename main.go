package main

import (
	"FIRECALC_REV2/stock"
	datagetter "FIRECALC_REV2/worker/dataGetter"
	"fmt"
	"strings"
	"sync"
	"time"
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

func main() {
	// get user selected stocks
	stocklist := getStocks()
	//collectedStockData := [][]stock.Stock{}

	var stocklistWG sync.WaitGroup
	var stocklistRWMU sync.RWMutex

	for index := range stocklist {
		stocklistWG.Add(1)
		go func(stock stock.Stock) {
			defer stocklistWG.Done()
			stocklistRWMU.RLock()
			defer stocklistRWMU.RUnlock()
			fmt.Printf("Goroutine: %v\n", time.Now())
			stockToGetData := datagetter.DataGetter{
				Stock:     stock,
				RWMutex:   &stocklistRWMU,
				Waitgroup: &stocklistWG,
			}

			stockToGetData.WorkData()
		}(stocklist[index])
	}

	stocklistWG.Wait()

}
