package yogofn

import "reflect"

func Filter(predicate, collection interface{}) interface{} {

	typecheck(collection, reflect.Array, reflect.Slice)
	typecheck(predicate, reflect.Func)

	if pT := reflect.TypeOf(predicate); pT.NumIn() != 1 || pT.NumOut() != 1 || pT.Out(0).Kind() != reflect.Bool {
		panic("Filter expected predicate: (func(T) bool), received: (" + pT.String() + ")")
	}

	collectionValue := reflect.ValueOf(collection)
	predicateValue := reflect.ValueOf(predicate)

	length := collectionValue.Len()

	// Make empty result slice
	resultType := reflect.SliceOf(reflect.TypeOf(collection).Elem())
	var resultValue reflect.Value = reflect.Zero(resultType)

	for i := 0; i < length; i++ {
		v := collectionValue.Index(i)
		if predicateValue.Call([]reflect.Value{v})[0].Bool() {
			resultValue = reflect.Append(resultValue, v)
		}
	}

	return resultValue.Interface()
}
