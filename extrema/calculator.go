package extrema

/*
毎フレーム価格を更新して直近のExtremaを返すオブジェクト
*/

type (
	Calculator struct {
		X         []int
		Highs     []float64
		Lows      []float64
		Maxima    *Extrema
		Minima    *Extrema
		ratio     float64
		maxLength int
	}

	Inf struct {
		Time  int
		Val   float64
		Which string //"high","low"
	}

	Infs struct {
		Last *Inf //直近の高値or底値
		Prev *Inf //Last１つ前の高値or底値
	}
)

func NewCalculator(ml int, ratio float64) *Calculator {
	return &Calculator{ratio: ratio, maxLength: ml}
}

func (h *Calculator) Update(x int, high, low float64) {
	h.Highs = append(h.Highs, high)
	h.Lows = append(h.Lows, low)
	h.X = append(h.X, x)
	h.shift()
}

//call this after h.Update
func (h *Calculator) Last() *Infs {
	maxima, _ := Search(h.X, h.Highs, h.ratio)
	_, minima := Search(h.X, h.Lows, h.ratio)
	maxT, maxV := lastFields(maxima)
	minT, minV := lastFields(minima)
	maxInf := &Inf{Time: maxT, Val: maxV, Which: "HIGH"}
	minInf := &Inf{Time: minT, Val: minV, Which: "LOW"}
	if maxT > minT {
		return &Infs{Last: maxInf, Prev: minInf}
	} else {
		return &Infs{Last: minInf, Prev: maxInf}
	}
	/*
		if maxT > minT {
			return &Inf{Time: maxT, Val: maxV, Which: "HIGH"}
		} else {
			return &Inf{Time: minT, Val: minV, Which: "LOW"}
		}
	*/

}

func lastFields(e *Extrema) (int, float64) {
	if len(e.Time) == 0 {
		return -1, -1
	}
	lt := e.Time[len(e.Time)-1]
	lv := e.Val[len(e.Val)-1]
	return lt, lv
}

func (h *Calculator) shift() {
	//h.X,h.Highs,h.Lowsは同じ長さ
	i := len(h.Highs) - h.maxLength
	if i > 0 {
		h.Highs = h.Highs[i:]
		h.Lows = h.Lows[i:]
		h.X = h.X[i:]
	}
}

/*
type (
	Holder struct {
		Prices    []float64
		X         []int
		Y         []float64
		ratio     float64
		maxLength int
	}
)


func NewHolder(ratio float64, ml int) *Holder {
	return &Holder{
		ratio: ratio, maxLength: ml,
	}
}

func (h *Holder) Update(v float64) {
	h.Prices = append(h.Prices, v)
	h.shift()
	if len(h.Prices) >= h.maxLength {
		h.setXY()
	}
}

func (h *Holder) Last() (int, float64) {
	if len(h.X) == 0 {
		return -1, -1
	}
	lx := h.X[len(h.X)-1]
	ly := h.Y[len(h.Y)-1]
	return lx, ly
}

func (h *Holder) shift() {
	st := len(h.Prices) - h.maxLength
	if st > 0 {
		h.Prices = h.Prices[st:]
	}
}

func (h *Holder) setXY() {
	Search(h.X, h.Y, h.ratio)
}
*/
