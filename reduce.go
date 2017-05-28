package yogofn

import (
	"fmt"
	"reflect"
)

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
