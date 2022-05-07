# Simple statistic calculation for trading bots.
## stats package
Supports mva,ema,rsi,bollinger band,and macd
## extrema package
Use Search func to search local min max.
> import "github.com/zenryokukun/dataprocess/extrema"
> 
> maxima,minima := extrema.Search([]int{1,2,3,...},[]float64{5.1,4.6,3.0,...},0.02)
