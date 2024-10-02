package worker

import (
	"FIRECALC_REV2/stock"
	"sync"
)

type Getter struct {
	id uint
}

type Worker struct {
	id        uint
	waitgroup *sync.WaitGroup
	mutex     *sync.Mutex
	stock     *stock.Stock
	dataList  *[][]stock.Stock
}
