# Simple statistic calculation for trading bots.
## stats package
Supports mva,ema,rsi,bollinger band,and macd
## extrema package
Use Search func to search local min max.
> import "github.com/zenryokukun/dataprocess/extrema"  
> maxima,minima := extrema.Search(*slice of int,slice of float64,ratio to detect localMinMax*)

Fields of maxima and minima.
