package stats

type (
	//初期化時に価格の配列が必要なほう
	Ema struct {
		Interval  int       //分母。何個分の平均にするか
		MaxLength int       //Emaの最大保持数
		Avg       []float64 //ema
		alpha     float64   // weight/(Interval+1)
	}

	//初期化時は価格は空
	EmaEmp struct {
		Prices    []float64
		Interval  int       //分母。何個分の平均にするか
		MaxLength int       //Emaの最大保持数
		Avg       []float64 //ema
		alpha     float64   // weight/(Interval+1)
	}
)

//**********************************************************************
//constructor
//**********************************************************************

//prices 価格配列
//itv　分母
//ml 最大保持数
//wieght ウェイト。2が一般的
func NewEma(prices []float64, itv, ml, weight int) *Ema {
	if len(prices) < itv {
		panic("len(prices)>=itv is not fullfilled")
	}
	ema := &Ema{
		Interval:  itv,
		MaxLength: ml,
		alpha:     alpha(itv, weight),
	}
	//初回emaセット
	first := ema.calcFirstEma(prices)
	//残りのema（配列がitvより長ければ）セット
	ema.Avg = append(ema.Avg, first)
	for _, v := range prices[itv:] {
		ema.Update(v)
	}
	return ema
}

//itv　分母
//ml 最大保持数
//wieght ウェイト。2が一般的
func NewEmaEmp(itv, ml, weight int) *EmaEmp {
	return &EmaEmp{
		Interval: itv, MaxLength: ml, alpha: alpha(itv, weight),
	}
}

//**********************************************************************
//[Ema] methods
//**********************************************************************

//更新処理
func (e *Ema) Update(v float64) {
	ema := e.calcEma(v)
	e.Avg = append(e.Avg, ema)
	e.shift()
}

//直近のemaを返す
func (e *Ema) Last() float64 {
	return e.Avg[len(e.Avg)-1]
}

//MaxLengthを超えたときに、古いデータを切り捨てる
func (e *Ema) shift() {
	if len(e.Avg) > e.MaxLength {
		st := len(e.Avg) - e.MaxLength
		e.Avg = e.Avg[st:]
	}
}

//初回ema計算
func (e *Ema) calcFirstEma(prices []float64) float64 {
	var sum float64
	for _, v := range prices[:e.Interval] {
		sum += v
	}
	ema := sum / float64(e.Interval)
	return ema
}

//2回目以降のema計算
//alpha = weight/(interval+1)
//ema = 現在価格v * alpha + 前ema*(1-alpha)
func (e *Ema) calcEma(v float64) float64 {
	last := e.Avg[len(e.Avg)-1]
	ema := v*e.alpha + last*(1-e.alpha)
	//fmt.Printf("current:%v,last:%v current:%v\n", v, last, ema)
	return ema
}

//**********************************************************************
//[EmaEmp] methods
//**********************************************************************

func (e *EmaEmp) Update(v float64) {
	if len(e.Avg) > 0 {
		//既に前回emaがある場合
		e.updateEma(v)
	} else {
		//まだemaが未設定なら価格更新
		e.Prices = append(e.Prices, v)
		if len(e.Prices) == e.Interval {
			//価格の長さが分母と等しくなったら初回ema計算
			e.updateFirstEma()
			//ema設定後は使わないのでクリアする
			e.Prices = nil
		}
	}
}

func (e *EmaEmp) Last() float64 {
	return e.Avg[len(e.Avg)-1]
}

func (e *EmaEmp) updateEma(v float64) {
	last := e.Avg[len(e.Avg)-1]
	ema := v*e.alpha + last*(1-e.alpha)
	e.Avg = append(e.Avg, ema)
	e.shift()
}

func (e *EmaEmp) updateFirstEma() {
	ema := average(e.Prices, e.Interval)
	e.Avg = append(e.Avg, ema)
}

func (e *EmaEmp) shift() {
	if len(e.Avg) > e.MaxLength {
		st := len(e.Avg) - e.MaxLength
		e.Avg = e.Avg[st:]
	}
}

//internal functions

//w weight,itv interval
func alpha(itv, w int) float64 {
	return float64(w) / (float64(itv) + 1.0)
}
