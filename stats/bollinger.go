package stats

import (
	"fmt"
	"math"
)

type (
	Bollinger struct {
		Mva *Mva
		Std []float64
	}
)

//itv 何足分か
//ml 最大何こもつか
func NewBollinger(itv, ml int) *Bollinger {
	mva := NewMva(itv, ml)
	return &Bollinger{
		Mva: mva,
	}
}

func (b *Bollinger) Update(v float64) {
	b.Mva.Update(v)
	if len(b.Mva.Avg) > 0 {
		b.setStd()
	}
}

//標準偏差を計算する関数
func (b *Bollinger) setStd() {
	fmt.Println(len(b.Mva.Avg))
	st := len(b.Mva.Prices) - b.Mva.Interval
	prices := b.Mva.Prices[st:]        //平均計算用の配列
	mva := b.Mva.Avg[len(b.Mva.Avg)-1] //直近の移動平均
	diffSum := 0.0
	for _, p := range prices {
		squaredDiff := math.Pow((p - mva), 2)
		diffSum += squaredDiff
	}

	variance := diffSum / float64(len(prices)) //分散
	std := math.Sqrt(variance)                 //標準偏差

	b.Std = append(b.Std, std)

	if len(b.Std) > b.Mva.MaxLength {
		st := len(b.Std) - b.Mva.MaxLength
		b.Std = b.Std[st:]
	}

}
