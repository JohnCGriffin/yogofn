package yogofn

import (
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

func TestSelection(t *testing.T) {
	// Empezemos normalmente, después con yogofn
	regular := letterCNormal()
	yogo := letterCYogo()
	if !mismo(regular, yogo) {
		t.Fail()
	}
}

func BenchmarkFilterPerformanceStandard(b *testing.B) {
	for i := 0; i < 20000; i++ {
		letterCNormal()
	}
}

func BenchmarkFilterPerformanceYogo(b *testing.B) {
	for i := 0; i < 20000; i++ {
		letterCYogo()
	}
}
