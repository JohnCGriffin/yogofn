package yogofn

import (
	"reflect"
)

func Map(mapper interface{}, collections ...interface{}) interface{} {

	typecheck(mapper, reflect.Func)
	mapperValue := reflect.ValueOf(mapper)

	if len(collections) < 1 {
		panic("Map requires at least one collection")
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

	// empty output slice
	resultType := reflect.SliceOf(reflect.TypeOf(mapper).Out(0))
	var resultValue reflect.Value = reflect.Zero(resultType)

	for i := 0; i < length; i++ {
		vals := make([]reflect.Value, 0)
		for _, c := range collectionValues {
			vals = append(vals, c.Index(i))
		}
		elem := mapperValue.Call(vals)[0]
		resultValue = reflect.Append(resultValue, elem)
	}

	return resultValue.Interface()
}
