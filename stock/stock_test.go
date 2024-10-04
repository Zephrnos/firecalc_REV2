package stock

import (
	"fmt"
	"testing"
)

func testStruct(t *testing.T) {
	var test Stock
	test.Ticker = "RIVN"

	fmt.Println(test.Ticker)
}
