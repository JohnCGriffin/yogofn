package yogofn

import (
	"testing"
)

func stringLengthsNormal() []int {
	regular := make([]int, 0)
	for _, t := range estados {
		regular = append(regular, len(t))
	}
	return regular
}

func stringLengthsYogoFn() []int {
	return Map(func(s string) int { return len(s) }, estados).([]int)
}

func mismoMap(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	return (a[0] == b[0]) && mismoMap(a[1:], b[1:])
}

func TestStringLength(t *testing.T) {
	// Empezemos normalmente, después con yogofn
	regular := stringLengthsNormal()
	yogo := stringLengthsYogoFn()
	if !mismoMap(regular, yogo) {
		t.Fail()
	}
}

func TestAddNaturalsUnevenLengths(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	doubled := Map(func(a, b int) int { return a + b }, nums, nums[:10]).([]int)
	for i := 0; i < len(nums) && i < len(doubled); i++ {
		if nums[i]*2 != doubled[i] {
			t.Fail()
		}
	}
}

func BenchmarkStringLengthPerformanceStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringLengthsNormal()
	}
}

func BenchmarkStringLengthPerformanceYogo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringLengthsYogoFn()
	}
}
