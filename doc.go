// Package yogofn offers Filter/Map/Reduce capabilities to Go.
// All arguments are typed as interface{} and reflection is utilized
// internally. Failure to supply good arguments fails with panic.
//
// All return values must be cast to the actual type.  For instance,
// here's code that takes the numbers from one to ten, filters out only
// the even ones, squares the result and then totals the squares.
//
//	nums := []int {1,2,3,4,5,6,7,8,9,10}
//	evens := Filter(func(n int) bool { return n % 2 == 0 }, nums).([]int)
//	squared := Map(func(n int) int { return n*n }, evens).([]int)
//	total := Reduce(func(a,b int) int { return a+b }, squared).(int)
//
// So why?  Because sometimes, I'd prefer to just do:
//
//	evens := Filter(func(n int) bool { return n % 2 == 0 }, nums).([]int)
//
// rather than:
//
//      results := make([]int,0)
//      for _,n := range evens {
//         if n % 2 == 0 {
//            result = append(results,n)
//         }
//      }
//
// Now, of course, Python programmers get to do the Filter/Map/Reduce succinctly at once:
//
//      sum([ x*x for x in nums if x%2 == 0 ])
//
// That's for later...
//
// A few handy reducers are added for common math reductions: AddF64, SubF64, AddInt, SubInt
package yogofn

// let's not forget
// go test -coverprofile=/tmp/coverage.out && go tool cover -html=/tmp/coverage.out
