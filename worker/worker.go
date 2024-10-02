package worker

import (
	"FIRECALC_REV2/stock"
	"sync"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/xmlpath.v2"
)

type Worker interface {
	WorkData()
}

type DataGetter struct {
	Id          uint
	Ticker string
	Waitgroup   *sync.WaitGroup
	DataList    *[]stock.Stock
}

type DataWorker struct {
	Id        uint
	Ticker string
	Waitgroup *sync.WaitGroup
	Mutex     *sync.Mutex
	Stock     *stock.Stock
	DataList  *[][]stock.Stock
}

func (w *DataWorker) WorkData() {

}

func (g *DataGetter) WorkData() {

}

func urlTickerPage(ticker string) string {
	completeURL string = fmt.Sprintf("https://finance.yahoo.com/quote/%s/history/?period1=0&period2=9999999999&frequency=1mo", ticker)
	return completeURL
}
