package yogofn

import (
	"math"
	"strings"
	"testing"
)

func letterCNormal() []string {
	result := make([]string, 0)
	for _, t := range estados {
		if strings.Index(t, "C") == 0 {
			result = append(result, t)
		}
	}
	return result
}

func letterCYogo() []string {
	return Filter(func(s string) bool { return strings.Index(s, "C") == 0 }, estados).([]string)
}

func mismo(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	return (a[0] == b[0]) && mismo(a[1:], b[1:])
}

func TestFilterSmallInts(t *testing.T) {
	if len(Filter(func(n int) bool { return 5 > n }, nums).([]int)) != 4 {
		t.Fail()
	}
}

func TestFilterSmallFloats(t *testing.T) {
	floats := []float64{10.0, 11.0, 12.0, 20.0, 30.0}
	if len(Filter(func(f float64) bool { return 20 > f }, floats).([]float64)) != 3 {
		t.Fail()
	}
}

func TestFilterNinos(t *testing.T) {
	var ninos = []persona{
		{"Becky", 12},
		{"Javier", 8},
		{"Shanika", 11},
		{"Camila", 7},
		{"Diego", 5},
		{"Alejandro", 17},
	}
	if len(Filter(func(p persona) bool { return p.edad > 10 }, ninos).([]persona)) != 3 {
		t.Fail()
	}
}

func TestAbsoluteValueMap(t *testing.T) {
	nums := make([]float64, 0)
	for x := 0.0; x < 100.0; x++ {
		nums = append(nums, x-30.0)
	}
	absValues := Map(math.Abs, nums).([]float64)
	for _, x := range absValues {
		if x < 0.0 {
			t.Fail()
		}
	}
}

func TestSelection(t *testing.T) {
	// Empezemos normalmente, después con yogofn
	regular := letterCNormal()
	yogo := letterCYogo()
	if !mismo(regular, yogo) {
		t.Fail()
	}
}

func BenchmarkFilterPerformanceStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCNormal()
	}
}

func BenchmarkFilterPerformanceYogo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCYogo()
	}
}
