package yogofn

import (
	"fmt"
	"reflect"
)

// Reduce takes an binary reducer function, i.e. (func(T1,T1) T2),
// and a collection(s) (i.e. an Array or Slice),
// resulting in scalar output the is the result of repeatedly applying the function to
// the previous output and the next input element. The binary function, the input
// collection and the output collection are all typed as interface{}. The result is
// almost always going to require a cast.
//
// For instance, summing a list of integers:
//
//   sum := Reduce(func(a,b int) int { return a+b }, numbers).(int)
//
// The initial value may be given as the third argument, but otherwise is the zero value
// of the return type of the reducer function.  If the input collection is empty, the
// return value is that zero value.
//
func Reduce(binary, collection interface{}, init ...interface{}) interface{} {

	if len(init) > 1 {
		panic(fmt.Sprintf("Reduce(collection,binaryFunction [,initValue]) received %d arguments", 2+len(init)))
	}

	// Common optimizations
	{
		switch t := binary.(type) {
		case func(float64, float64) float64:
			var accum float64
			if len(init) == 1 {
				accum = init[0].(float64)
			}
			for _, f := range collection.([]float64) {
				accum = t(f, accum)
			}
			return accum

		case func(int, int) int:
			var accum int
			if len(init) == 1 {
				accum = init[0].(int)
			}
			for _, n := range collection.([]int) {
				accum = t(n, accum)
			}
			return accum

		case func(string, string) string:
			var accum string
			if len(init) == 1 {
				accum = init[0].(string)
			}
			for _, s := range collection.([]string) {
				accum = t(s, accum)
			}
			return accum
		}
	}

	typecheck(collection, reflect.Array, reflect.Slice)
	typecheck(binary, reflect.Func)
	binaryType := reflect.TypeOf(binary)

	if binaryType.NumIn() != 2 ||
		binaryType.NumOut() != 1 {
		panic("Reduce expected arity-2 reducer: (func(T1,T2) T3), received (" + binaryType.String() + ")")
	}

	collectionValue := reflect.ValueOf(collection)
	binaryValue := reflect.ValueOf(binary)
	length := collectionValue.Len()

	resultType := reflect.TypeOf(binary).Out(0)
	accum := reflect.Zero(resultType)

	if len(init) == 1 {
		accum = reflect.ValueOf(init[0])
	}

	for i := 0; i < length; i++ {
		v := collectionValue.Index(i)
		accum = binaryValue.Call([]reflect.Value{v, accum})[0]
	}

	return accum.Interface()
}
