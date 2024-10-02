package main

import (
	"fmt"
	"strings"
	// sv "sync"
	// c "github.com/gocolly/colly"
)

type Stock struct {
	ticker string
	ratio  uint8
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
				break // Allows you to proceede if your ratio of the stock is valid
			}
		}
		// Apply the ratio to the "Stock" object and incrament the portfolioRatioTotal counter
		// Finally, append the stock to the "stocks" slice
		stockList = append(stockList, newStock)
		if portfolioRatioTotal == 100 {
			break // Exits once we have enetred all our stock options
		}
	}
	return stockList
}

func main() {

	stocks := getStocks()

	for index, stock := range stocks {
		fmt.Printf("Index: %d, Stock %s, Ratio %d\n", index, stock.ticker, stock.ratio)
	}

}
