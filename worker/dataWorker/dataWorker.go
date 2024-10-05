package dataworker

import (
	"FIRECALC_REV2/stock"
	"sync"
)

type DataWorker struct {
	Stock     *stock.Stock
	Waitgroup *sync.WaitGroup
	Mutex     *sync.Mutex
	DataList  *[][]stock.Stock
}

func (w *DataWorker) WorkData() {
	// Implement your worker logic here
}
