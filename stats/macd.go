package stats

import "math"

type (
	Macd struct {
		Short    *EmaEmp //短期EMA default:12
		Long     *EmaEmp //長期EMA default:26
		Macd     []float64
		Signal   []float64
		Interval int //signalの分母
	}
)

//12ema,26ema,signal 9mvaの標準的なmacdのコンストラクタ
func NewMacdDefault() *Macd {
	return &Macd{
		Short:    NewEmaEmp(12, 20, 2),
		Long:     NewEmaEmp(26, 30, 2),
		Interval: 9,
	}
}

func (m *Macd) Update(v float64) {
	m.Short.Update(v)
	m.Long.Update(v)
	if m.isMacdReady() {
		m.updateMacd(v)
	}
	if m.isSignalReady() {
		m.updateSignal()
	}
}

//returns the most recent MACD
func (m *Macd) Last() float64 {
	if len(m.Macd) == 0 {
		return math.NaN()
	}
	return m.Macd[len(m.Macd)-1]
}

//returns the most recent Signal
func (m *Macd) LastSignal() float64 {
	if len(m.Signal) == 0 {
		return math.NaN()
	}
	return m.Signal[len(m.Signal)-1]
}

func (m *Macd) updateMacd(v float64) {
	macd := m.Short.Last() - m.Long.Last()
	m.Macd = append(m.Macd, macd)
	//macdがLongの最大個数を超える場合、切り捨て。
	i := len(m.Macd) - m.Long.MaxLength
	if i > 0 {
		m.Macd = m.Macd[i:]
	}
}

//signal更新。isSignalReadyでチェックしてから呼ぶこと
//m.updateMacdの後に呼ぶこと
func (m *Macd) updateSignal() {
	i := len(m.Macd) - m.Interval
	if i < 0 {
		i = 0
	}
	targ := m.Macd[i:]
	sig := average(targ, m.Interval)
	m.Signal = append(m.Signal, sig)

	//SignalがLongの最大個数を超える場合、切り捨て。
	j := len(m.Signal) - m.Long.MaxLength
	if j > 0 {
		m.Signal = m.Signal[j:]
	}

}

//macdを更新処理が出来るだけのデータがたまっているか。
//Longのemaが１つ以上あれば準備OK状態
func (m *Macd) isMacdReady() bool {
	if len(m.Long.Avg) > 0 {
		return true
	}
	return false
}

//signalを更新処理が出来るだけのデータがたまっているか。
//Macdの長さがInterval以上なら準備OK状態
func (m *Macd) isSignalReady() bool {
	if len(m.Macd) >= m.Interval {
		return true
	}
	return false
}
