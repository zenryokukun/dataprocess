package stats

import "math"

type (
	Mva struct {
		Prices    []float64 //価格の配列
		Avg       []float64 //mvaを入れる配列
		Interval  int       //何足分の平均とするか
		MaxLength int       //mvaを最大何個保持するか
	}
)

//コンストラクタ
//itv 分母。何個分の平均とするか
//ml　Avgを最大何個保持するか
func NewMva(itv int, ml int) *Mva {
	return &Mva{Interval: itv, MaxLength: ml}
}

//平均を計算する。itvはlen(prices)と同じ。いらなかったかも
func average(prices []float64, itv int) float64 {
	sum := 0.0
	for _, v := range prices {
		sum += v
	}
	return sum / float64(itv)
}

//更新処理。priceは現在価格を想定
func (m *Mva) Update(price float64) {
	m.Prices = append(m.Prices, price)
	if len(m.Prices) >= m.Interval {
		st := len(m.Prices) - m.Interval
		avg := average(m.Prices[st:], m.Interval)
		m.Avg = append(m.Avg, avg)
		m.shift()
	}
}

func (m *Mva) Last() float64 {
	if len(m.Avg) == 0 {
		return math.NaN()
	}
	return m.Avg[len(m.Avg)-1]
}

//MaxLengthを超えたときに、古いデータを切り捨てる
func (m *Mva) shift() {
	if len(m.Prices) > m.MaxLength {
		st := len(m.Prices) - m.MaxLength
		m.Prices = m.Prices[st:]
	}
	if len(m.Avg) > m.MaxLength {
		st := len(m.Avg) - m.MaxLength
		m.Avg = m.Avg[st:]
	}
}
