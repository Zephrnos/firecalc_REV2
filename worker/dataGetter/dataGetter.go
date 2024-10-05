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
	Stock     stock.Stock
	RWMutex   *sync.RWMutex
	Waitgroup *sync.WaitGroup
}

func (g *DataGetter) WorkData() {
	// Implement your getter logic here
	getHTML(g.Stock.Ticker)
}

func urlTickerPage(ticker string) string {
	return fmt.Sprintf("https://finance.yahoo.com/quote/%s/history/?period1=0&period2=9999999999&frequency=1mo", ticker)
}

func getHTML(ticker string) error {
	url := urlTickerPage(ticker)
	fmt.Printf("Fetching URL: %s\n", url)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// Set the User-Agent header to mimic a browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")

	// Perform the request using an HTTP client
	client := &http.Client{}
	response, err := client.Do(req)
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
