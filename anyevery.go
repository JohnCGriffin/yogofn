package yogofn

import (
	"reflect"
)

type _mode int

const (
	any _mode = iota
	every
)

func _anyEveryWorker(mode _mode, predicate interface{}, collections []interface{}) bool {

	typecheck(predicate, reflect.Func)
	predicateValue := reflect.ValueOf(predicate)

	if len(collections) < 1 {
		panic("requires at least one collection")
	}

	// make reflect.Value array to hold collection values
	collectionValues := func() []reflect.Value {
		tmp := make([]reflect.Value, 0)
		for _, c := range collections {
			typecheck(c, reflect.Array, reflect.Slice)
			tmp = append(tmp, reflect.ValueOf(c))
		}
		return tmp
	}()

	// establish common length
	length := collectionValues[0].Len()
	for _, v := range collectionValues {
		if length > v.Len() {
			length = v.Len()
		}
	}

	for i := 0; i < length; i++ {
		vals := make([]reflect.Value, 0)
		for _, c := range collectionValues {
			vals = append(vals, c.Index(i))
		}
		if predicateValue.Call(vals)[0].Interface().(bool) {
			if mode == any {
				return true
			}
		} else {
			if mode == every {
				return false
			}
		}
	}

	if mode == any {
		return false
	}
	return true
}

// Any takes a predicate function and a collection(s) (i.e. an Array or Slice),
// returning true if the predicate is satisfied on an element.  The predicate is
// passed a number of elements matching the number of collections given.
//
// For instance, identifying negative numbers
//
//   if Any(func(n int){ return n < 0 },nums) {
//	    ....
//
// Note that the number of collections in the example is one.  However,
// any number of input lists may be used.  Each input is iterated until one of them
// is empty.  So, given low and high temperature lists, identify if a range above 50
// degrees exists.
//
//   bigRange := func(a,b float64){ return math.Abs(a-b) > 50 }
//   if Any(bigRange,lows,highs){
//       ...
//
func Any(predicate interface{}, collections ...interface{}) bool {
	return _anyEveryWorker(any, predicate, collections)
}

// Every takes a predicate function and a collection(s) (i.e. an Array or Slice),
// returning true if the predicate is satisfied on each element.  The predicate is
// passed a number of elements matching the number of collections given.
//
// For instance, ensuring positive numbers
//
//   if !Every(func(n int){ return n > 0 },nums) {
//	    ....
//
// Note that the number of collections in the example is one.  However,
// any number of input lists may be used.  Each input is iterated until one of them
// is empty.  So, given low and high temperature lists, ensure that all have
// low <= high.
//
//   wellOrdered := func(a,b float64){ return a <= b }
//   if ! Every(wellOrdered,lows,highs){
//       ...
//
func Every(predicate interface{}, collections ...interface{}) bool {
	return _anyEveryWorker(every, predicate, collections)
}
