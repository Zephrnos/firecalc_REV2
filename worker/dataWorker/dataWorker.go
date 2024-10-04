package dataworker

import (
	"FIRECALC_REV2/stock"
	"sync"
)

type DataWorker struct {
	Id        uint
	Ticker    string
	Waitgroup *sync.WaitGroup
	Mutex     *sync.Mutex
	Stock     *stock.Stock
	DataList  *[][]stock.Stock
}

func (w *DataWorker) WorkData() {
	// Implement your worker logic here
}
