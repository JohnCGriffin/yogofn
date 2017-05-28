package yogofn

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Filter(func(i int) bool {
		return i%2 == 1
	}, [5]int{10, 100, 200, 23, 55}))
	fmt.Println("start")
	stuff := make([]int, 0)
	for i := 0; i < 1000; i++ {
		stuff = append(stuff, i)
	}
	for i := 0; i < 1000; i++ {
		Filter(func(i int) bool {
			return i%2 == 1
		}, stuff)
	}
	fmt.Println("done")
	x := Map(func(n int) float64 { return float64(n * n) }, []int{1, 2, 3}).([]float64)
	fmt.Println(Map(func(n int) string { return fmt.Sprintf("%d bottles of beer", n) }, []int{1, 2, 3}))
	fmt.Printf("type(x) = %T\n", x)
	z := Filter(func(i int) bool { return i%2 == 1 }, stuff).([]int)
	fmt.Printf("type(z) = %T\n", z)

	total := Reduce(func(a, b int) int { return a + b }, stuff[1:101]).(int)
	fmt.Println("HOWDY", total)
	total = Reduce(func(a, b int) int { return a + b }, stuff[1:101], 1).(int)
	fmt.Println("HOWDY", total)
	temps := []float64{44.5, 67.8, 66.4, 61.2, 44.3}
	fmt.Printf("max of %v is %v\n", temps, Reduce(math.Max, temps).(float64))

	words := []string{"one", "two", "three", "four"}
	longerOf := func(a, b string) string {
		if len(a) > len(b) {
			return a
		}
		return b
	}
	fmt.Printf("longest word in %v is %v\n", words, Reduce(longerOf, words))

	smalls := []int{1, 2, 3, 4}
	bigs := []int{20, 20, 30}

	fmt.Println(Map(func(a, b int) int { return a + b }, smalls, bigs))
	fmt.Println(Map(func(a, b int) []int { return []int{a, b} }, smalls, bigs))

	Filter(func(a int) bool { return a > 1 }, []int{1, 2, 3})
	//Filter(math.Max, []int{1, 2, 3})
	Reduce(math.Abs, []int{1, 2, 3})
}
