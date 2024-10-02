package stock

type Stock struct {
	Ticker string
	Ratio  uint8
	Price  float32
	Change float32
	Date   string
	Div    float32
}
