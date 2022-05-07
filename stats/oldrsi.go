package stats

import (
	"math"
)

//RSI
//A/(A+B)が基本
//m -> 日数.基本は14
//A -> 上昇分の合計/n
//B -> 下落分の合計/m
//差分が14個必要となるため、配列の長さは15必要。

var (
	defaultDays = 14
	prevPlus    float64 //＋の移動平均
	prevMinus   float64 //-の移動平均
)

func updateGlobal(plusAvg, minusAvg float64) {
	prevPlus = plusAvg
	prevMinus = minusAvg
}

func changeDefaultDays(newDays int) {
	defaultDays = newDays
}

//初回RSI計算。Cutler式
//nのrsiを計算するためには、pricesの長さはn+1である必要がある点に注意
//nとn-1の差分がn個必要なため。
//nはdefaultDays
func _cutler(prices []float64) float64 {
	if len(prices) != defaultDays+1 {
		panic("len of prices must be defaultDays+1")
	}
	plus := 0.0
	minus := 0.0
	length := float64(defaultDays)
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
	if A == B {
		return 0.5
	}
	updateGlobal(A, B)
	return A / (A + B)
}

//２回目以降RSI計算。Wilder式
//nのrsiを計算するためには、pricesの長さはn+1である必要がある点に注意
//nとn-1の差分がn個必要なため。
func wilder(prices []float64) float64 {
	if len(prices) != defaultDays+1 {
		//panic("len of prices must be defaultDays+1")
	}
	length := defaultDays
	last := prices[len(prices)-1]  //last
	last2 := prices[len(prices)-2] //second to last
	diff := last - last2
	prevA := prevPlus * float64(length-1)
	prevB := prevMinus * float64(length-1)
	var A, B float64
	if diff > 0 {
		A = (prevA + diff) / float64(length)
		B = prevB / float64(length)
	} else {
		B = (prevB + math.Abs(diff)) / float64(length)
		A = prevA / float64(length)
	}
	updateGlobal(A, B)
	return A / (A + B)
}

func CalcRsi(prices []float64) float64 {
	if prevPlus == 0.0 && prevMinus == 0.0 {
		//初回
		return _cutler(prices)
	} else {
		//2回目以降
		return wilder(prices)
	}
}
