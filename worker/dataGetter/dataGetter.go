package datagetter

import (
	"FIRECALC_REV2/stock"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type DataGetter struct {
	Id        uint
	Ticker    string
	RWMutex   *sync.RWMutex
	Waitgroup *sync.WaitGroup
	DataList  *[]stock.Stock
}

func (g *DataGetter) WorkData() {
	// Implement your getter logic here
	getHTML(g.Ticker)
}

func urlTickerPage(ticker string) string {
	return fmt.Sprintf("https://finance.yahoo.com/quote/%s/history/?period1=0&period2=9999999999&frequency=1mo", ticker)
}

func getHTML(ticker string) error {
	url := urlTickerPage(ticker)
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching URL: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("status code error: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	filename := fmt.Sprintf("%s.html", ticker)
	err = os.WriteFile(filename, body, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	fmt.Printf("HTML saved to %s\n", filename)
	return nil
}
