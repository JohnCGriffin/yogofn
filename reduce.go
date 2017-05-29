package yogofn

import (
	"fmt"
	"reflect"
)

// Reduce takes an binary reducer function, i.e. 2-arity,
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

	typecheck(collection, reflect.Array, reflect.Slice)
	typecheck(binary, reflect.Func)
	binaryType := reflect.TypeOf(binary)
	if binaryType.NumIn() != 2 || binaryType.NumOut() != 1 {
		panic("Reduce expected arity-2 reducer: (func(T1,T2) T3), received (" + binaryType.String() + ")")
	}

	collectionValue := reflect.ValueOf(collection)
	binaryValue := reflect.ValueOf(binary)
	length := collectionValue.Len()

	resultType := reflect.TypeOf(binary).Out(0)
	accum := reflect.Zero(resultType)

	switch len(init) {
	case 0:
	case 1:
		accum = reflect.ValueOf(init[0])
	default:
		panic(fmt.Sprintf("Reduce(collection,binaryFunction [,initValue]) received %d arguments", 2+len(init)))
	}

	for i := 0; i < length; i++ {
		v := collectionValue.Index(i)
		accum = binaryValue.Call([]reflect.Value{v, accum})[0]
	}

	return accum.Interface()
}
