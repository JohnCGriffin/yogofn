package yogofn

import (
	"math"
	"testing"
)

func masLargoNormal() string {
	answer := ""
	for _, estado := range estados {
		if len(estado) > len(answer) {
			answer = estado
		}
	}
	return answer
}

func masLargoYogo() string {
	return Reduce(func(a, b string) string {
		if len(a) > len(b) {
			return a
		}
		return b
	}, estados).(string)
}

func mismoLargo(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	return (a[0] == b[0]) && mismoLargo(a[1:], b[1:])
}

func TestMismoTotal(t *testing.T) {
	total := 0
	for _, e := range estados {
		total += len(e)
	}
	if total != Reduce(func(a string, b int) int { return b + len(a) }, estados).(int) {
		t.Fail()
	}
}

func TestMasLargo(t *testing.T) {
	// Empezemos normalmente, después con yogofn
	regular := masLargoNormal()
	yogo := masLargoYogo()
	if len(regular) != len(yogo) {
		t.Fail()
	}
}

func expectPanic(t *testing.T, thunk func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	thunk()
}

func TestNoCollection(t *testing.T) {
	expectPanic(t,
		func() {
			Any(func(a bool) bool { return a })
		})
}

func TestWrongFilterArgs(t *testing.T) {
	expectPanic(t,
		func() {
			Filter(func() bool { return true }, []int{})
		})
}

func TestWrongFilterReturn(t *testing.T) {
	expectPanic(t,
		func() {
			Filter(func(a bool) {}, []int{})
		})
}

func TestMissingMapCollection(t *testing.T) {
	expectPanic(t,
		func() {
			Map(func(a bool) {})
		})
}

func TestReduceWithInitialValue(t *testing.T) {
	if Reduce(func(a, b int) int { return a + b }, []int{1, 2, 3}, 10) != 16 {
		t.Fail()
	}
}

func TestReduceWithTooManyArguments(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Reduce(func(a, b int) int { return a + b }, []int{1, 2, 3}, 10, 12345)
}

func TestBadReduceBinaryArity(t *testing.T) {
	expectPanic(t,
		func() {
			Reduce(func(a bool) bool { return true }, []bool{})
		})
}

func BenchmarkMasLargoStandard(b *testing.B) {
	for i := 0; i < 10000; i++ {
		masLargoNormal()
	}
}

func BenchmarkMasLargoYogo(b *testing.B) {
	for i := 0; i < 10000; i++ {
		masLargoYogo()
	}
}

func BenchmarkTotalFirst100Floats(b *testing.B) {
	nums := make([]float64, 0)
	for i := 1; i <= 100; i++ {
		nums = append(nums, float64(i))
	}
	for i := 0; i < b.N; i++ {
		total := Reduce(func(a, b float64) float64 { return a + b }, nums).(float64)
		if math.Abs(total-5050.0) > 0.1 {
			b.Fail()
		}
	}
}
