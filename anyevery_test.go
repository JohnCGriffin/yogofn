package yogofn

import (
	"testing"
)

func TestEveryTrue(t *testing.T) {
	if !Every(func(n int) bool { return n > 0 }, nums) {
		t.Fail()
	}
}

func TestAnyTrue(t *testing.T) {
	if !Any(func(n int) bool { return n > 0 }, nums) {
		t.Fail()
	}
}

func TestEveryFalse(t *testing.T) {
	if Every(func(n int) bool { return n%2 == 0 }, nums) {
		t.Fail()
	}
}

func TestAnyFalse(t *testing.T) {
	if Any(func(n int) bool { return n < 0 }, nums) {
		t.Fail()
	}
}

func TestEveryTrueTwoLists(t *testing.T) {
	squares := Map(func(n int) int { return n * n }, nums).([]int)
	if !Every(func(a, b int) bool { return a <= b }, nums, squares) {
		t.Fail()
	}
}

func TestAnyTrueTwoLists(t *testing.T) {
	squares := Map(func(n int) int { return n * n }, nums).([]int)
	if !Any(func(a, b int) bool { return a <= b }, nums, squares) {
		t.Fail()
	}
}

func TestEveryTrueTwoListsUnequalLength(t *testing.T) {
	squares := Map(func(n int) int { return n * n }, nums).([]int)[:5]
	if !Every(func(a, b int) bool { return a <= b }, nums, squares) {
		t.Fail()
	}
}

func TestBuscarJaliscoEveryFalse(t *testing.T) {
	if Every(func(s string) bool { return s == "Jalisco" }, estados) {
		t.Fail()
	}
}

func TestBuscarJaliscoAnyTrue(t *testing.T) {
	if !Any(func(s string) bool { return s == "Jalisco" }, estados) {
		t.Fail()
	}
}

func TestAdultosEveryTrue(t *testing.T) {
	if !Every(func(p persona) bool { return p.edad >= 18 }, adultos) {
		t.Fail()
	}
}

func TestEdadesEveryTrue(t *testing.T) {
	f := func(nino, adulto persona) bool { return nino.edad < adulto.edad }
	if !Every(f, ninos, adultos) {
		t.Fail()
	}

}
