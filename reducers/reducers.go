// Package reducers contains common functions used by yogofn.Reduce.
//
// For instance, to sum numbers
//
//   total := Reduce(AddF64, someFloat64s).(float64)
//
// Or to find the maximum daily temperature range given daily high and
// low temperatures.
//
//   widestDailyRange := Reduce(MaxF64, Map(SubF64, highs, lows)).(float64)
//
// Or to find a bank balance with an initial balance of $1250
//
//   balance := Reduce(AddF64, transactions, 1250).(float64)
//
package reducers

func AddF64(a, b float64) float64 { return a + b }

func AddInt(a, b int) int { return a + b }

func SubF64(a, b float64) float64 { return a - b }

func SubInt(a, b int) int { return a - b }

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinF64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func MaxF64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
