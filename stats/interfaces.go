package stats

type Indicator interface {
	Update(v float64)
	Last() float64
}
