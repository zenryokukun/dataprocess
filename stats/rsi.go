package stats

//初回はcutler式
//A/(A+B)が基本
//n -> 日数.基本は14
//A -> 上昇分の合計/n
//B -> 下落分の合計/n
//差分が14個必要となるため、配列の長さは15必要。

//2回目以降は、前回RSIを元に計算する(wilder式)：
//prevA = A'*(n-1),prevB = B'*(n-1)
//diff = P[現在] -P'[１つ前]
//A = (prevA+diff)/n  ...diff > 0
//B = (prevB+diff)/n ... diff < 0
//ris = A/(A+B)

import "math"

type (
	//初期化するときに価格も一緒に渡すやつ
	Rsi struct {
		Rsi       []float64
		A         float64 //plus average
		B         float64 //minus average
		MaxLength int     //Rsiの最大保有数
		n         int     //日数 default:14
		last      float64 //最後の価格
	}
	//価格なしで初期化するやつ
	RsiEmp struct {
		Prices []float64
		Rsi
	}
)

func cutler(prices []float64, n int) (float64, float64, float64) {
	plus := 0.0
	minus := 0.0
	length := float64(n)
	prev := prices[0]
	for _, v := range prices[1:] {
		diff := v - prev
		if diff >= 0.0 {
			plus += diff
		} else {
			minus += diff
		}
		prev = v
	}

	A := plus / length
	B := math.Abs(minus) / length

	return A / (A + B), A, B
}

//************************************************
//constructors
//************************************************

//prices 価格の配列。n+1の長さである必要がある。
//n　何個分の平均とするか。ふつうは14だけど選べる。
//ml　最大何個のrsiを保存しておくか
func NewRsi(prices []float64, n, ml int) *Rsi {
	if len(prices) != n+1 {
		panic("NewRsi:len(prices) must be n+1 exact.")
	}
	R, A, B := cutler(prices, n)
	return &Rsi{
		Rsi:       []float64{R},
		A:         A,
		B:         B,
		n:         n,
		last:      prices[len(prices)-1],
		MaxLength: ml,
	}
}

func NewRsiEmp(n, ml int) *RsiEmp {
	return &RsiEmp{
		Rsi: Rsi{
			n:         n,
			MaxLength: ml,
		},
	}
}

//************************************************
//Rsi methods
//************************************************

//rsiを更新する関数
//vは現在価格を想定
func (r *Rsi) Update(v float64) {

	diff := v - r.last
	pA := r.A * float64(r.n-1)
	pB := r.B * float64(r.n-1)

	var nA, nB float64
	if diff > 0 {
		nA = (pA + diff) / float64(r.n)
		nB = pB / float64(r.n)
	} else {
		nB = (pB + math.Abs(diff)) / float64(r.n)
		nA = pA / float64(r.n)
	}

	rsi := nA / (nA + nB)
	r.Rsi = append(r.Rsi, rsi)
	r.A = nA
	r.B = nB
	r.last = v

	if len(r.Rsi) > r.MaxLength {
		st := len(r.Rsi) - r.MaxLength
		r.Rsi = r.Rsi[st:]
	}
}

//************************************************
//RsiEmp methods
//************************************************
func (r *RsiEmp) Update(v float64) {
	if len(r.Rsi.Rsi) > 0 {
		//２回目以降のrsi計算
		r.Rsi.Update(v)
	} else {
		//まだRsi計算する分のpriceが溜まっていない状態
		r.Prices = append(r.Prices, v)
		if len(r.Prices) == r.n+1 {
			//溜まったら初回の設定処理
			rsi, A, B := cutler(r.Prices, r.n)
			r.Rsi.Rsi = append(r.Rsi.Rsi, rsi)
			r.Rsi.A = A
			r.Rsi.B = B
			r.Rsi.last = v
			r.Prices = nil
		}
	}
}

func (r *RsiEmp) Last() float64 {
	if len(r.Rsi.Rsi) == 0 {
		return math.NaN()
	}
	return r.Rsi.Rsi[len(r.Rsi.Rsi)-1]
}
