package yogofn

import (
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

func BenchmarkMasLargoStandard(b *testing.B) {
	for i := 0; i < 20000; i++ {
		masLargoNormal()
	}
}

func BenchmarkMasLargoYogo(b *testing.B) {
	for i := 0; i < 20000; i++ {
		masLargoYogo()
	}
}
