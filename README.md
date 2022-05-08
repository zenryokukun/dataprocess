# Simple statistic calculation for trading bots.
## stats package
Supports mva,ema,rsi,bollinger band,and macd
## extrema package
Use Search func to search local min max.
> import "github.com/zenryokukun/dataprocess/extrema"  
> maxima,minima := extrema.Search(*X,Y,Ratio*)
> *X* is a slice of *int* representing time series.
> *Y* is a slice of *float64* representing prices.
> Ratio is a *float64* for detecting local min max. 

Fields of maxima and minima.
