# yogofn
Functional Map/Reduce/Filter for Golang

### install
go get github.com/johncgriffin/yogofn

Yogofn offers a few slice generating operations as alternatives to Go's for-loops.

For instance, let's find the max daily temperature range given two parallel lists of low and high daily temperatures.
Here's standard Go code given two float64 slices.

```
var maxDailyRange float64
for i:=0; i<len(highs) && i<len(lows); i++ {
    maxDailyRange = math.Max(highs[i] - lows[i], maxDailyRange)
}
```
And here's using yogofn:
```
maxDailyRange := Reduce(MaxF64, Map(SubF64, highs, lows)).(float64)
```

Notice that the final outside operation (Reduce in this instance) required a cast to inform Go of the final type.  Inputs are either arrays or slices of any type.  Yogofn offers:

- ``Map(f,slice(s))`` -> slice projection
- ``Filter(f,slice)`` -> slice selection
- ``Reduce(f,slice(s))`` -> scalar
- ``Every(f,slice)`` -> bool 
- ``Any(f,slice)`` -> bool

Notice that because Reduce and Map can take more than one list, Zip is effected via

```
zipped := Map(func(x,y string) []string { return []string{x,y} }, xs, ys).([][]string)
```
A companion package is yogofn/reducers which contains simple but common scalar reductions for int and float64 numbers.

#### Reflection Performance

Not suprisingly, reflection slows down standard Go about 200 times.  Using Go's type switch, a few optimizations 
were placed into Map/Filter/Reduce 
for common data types float64, int, and string.  Those operations are implemented as normal typed Go loops.




