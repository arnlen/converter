package exchangeratesapi

type Rate struct {
	Date   string
	Base   string
	Symbol string
	Rate   float64
}
