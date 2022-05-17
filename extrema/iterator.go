package extrema

type (
	//Updateで価格をため込むやつ
	Iterator struct {
		TimeIndices []int
		Highs       []float64
		Lows        []float64
		ratio       float64
	}
	//ある時点における高値・底値情報
	localInf struct {
		Price     float64 //価格
		Time      string  //時間
		TimeIndex int     //時間をindexで表現
		Distance  int     //現在時刻との距離
	}
	HighLow struct {
		High *localInf
		Low  *localInf
	}
)

func NewLocalInf() *HighLow {
	return &HighLow{
		High: &localInf{}, Low: &localInf{},
	}
}

func NewIterator(ratio float64) *Iterator {
	return &Iterator{
		ratio: ratio,
	}
}

func (i *Iterator) Update(x int, h, l float64) *HighLow {
	i.TimeIndices = append(i.TimeIndices, x)
	i.Highs = append(i.Highs, h)
	i.Lows = append(i.Lows, l)
	return i.getLocal()
}

func (i *Iterator) getLocal() *HighLow {
	mx, _ := Search(i.TimeIndices, i.Highs, i.ratio)
	_, mn := Search(i.TimeIndices, i.Lows, i.ratio)
	now := i.TimeIndices[len(i.TimeIndices)-1]
	mxInf := getLocalInf(mx, now)
	mnInf := getLocalInf(mn, now)
	return &HighLow{
		High: mxInf,
		Low:  mnInf,
	}
}

func getLocalInf(e *Extrema, currentX int) *localInf {
	if len(e.Time) == 0 {
		return nil
	}
	lastV := e.Val[len(e.Val)-1]
	lastX := e.Time[len(e.Time)-1]
	return &localInf{
		Price:     lastV,
		TimeIndex: lastX,
		Distance:  currentX - lastX,
	}
}

/*
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
*/
