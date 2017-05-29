# yogofn
Functional Map/Reduce/Filter for Golang

Given that Golang is typed and verbose, functional-style operations are unidiomatic.  

A few of the common ones are here.  They allow 
such actions as finding the max daily temperature range given two parallel lists of low and high daily temperatures.
Here's standard Go code given two float64 slices.

```
var maxDiff float64
for i:=0; i<len(highs) && i<len(lows); i++){
    maxDiff := math.Max(highs[i] - lows[i], maxDiff)
}
```
And here's using yogofn:
```
widestDailyRange := Reduce(MaxF64, Map(SubF64, highs, lows)).(float64)
```

Notice that the final outside operation (Reduce in this instance) required a cast to inform Go of the final type.  Inputs are either arrays or slices of any type.
