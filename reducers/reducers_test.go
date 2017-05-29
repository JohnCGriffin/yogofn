package reducers

import (
	"testing"
)

func TestInt(t *testing.T) {

	intTest := func(t *testing.T, fName string, f func(int, int) int, a, b, expect int) {
		if tmp := f(a, b); tmp != expect {
			t.Errorf("(%v,%v) expect %v received %v", a, b, expect, tmp)
		}
	}

	table := []struct {
		name         string
		f            func(int, int) int
		a, b, expect int
	}{
		{"AddInt", AddInt, 100, 200, 300},
		{"AddInt", AddInt, 10, -20, -10},
		{"SubInt", SubInt, 10, 6, 4},
		{"SubInt", SubInt, 10, -20, 30},
		{"MaxInt", MaxInt, 1, 2, 2},
		{"MaxInt", MaxInt, 2, 1, 2},
		{"MinInt", MinInt, 1, 2, 1},
		{"MinInt", MinInt, 2, 1, 1},
	}

	for _, entry := range table {
		intTest(t, entry.name, entry.f, entry.a, entry.b, entry.expect)
	}
}

func TestF64(t *testing.T) {

	f64Test := func(t *testing.T, fName string, f func(float64, float64) float64, a, b, expect float64) {
		if tmp := f(a, b); tmp != expect {
			t.Errorf("(%v,%v) expect %v received %v", a, b, expect, tmp)
		}
	}

	table := []struct {
		name         string
		f            func(float64, float64) float64
		a, b, expect float64
	}{
		{"AddF64", AddF64, 100.0, 200.0, 300.0},
		{"AddF64", AddF64, 10.0, -20.0, -10.0},
		{"SubF64", SubF64, 10.0, 6.0, 4.0},
		{"SubF64", SubF64, 10.0, -20.0, 30.0},
		{"MaxF64", MaxF64, 1.0, 2.0, 2.0},
		{"MaxF64", MaxF64, 2.0, 1.0, 2.0},
		{"MinF64", MinF64, 1.0, 2.0, 1.0},
		{"MinF64", MinF64, 2.0, 1.0, 1.0},
	}

	for _, entry := range table {
		f64Test(t, entry.name, entry.f, entry.a, entry.b, entry.expect)

	}

}
