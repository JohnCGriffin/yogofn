package yogofn

import "reflect"

// Filter takes a boolean predicate and a collection (i.e. an Array or Slice),
// resulting in a reduced output slice where each element passes the predicate.
// The predicate, the input collection and the output collection are all
// typed as interface{}. The result is almost always going to require a cast.
//
// For instance, filtering out names with "Z" might be:
//
//   nonZs := Filter(
//		func(name string) bool { return !strings.Contains(name,"Z"); },
//		names).([]string)
//
func Filter(predicate, collection interface{}) interface{} {

	// common optimizations
	{
		switch f1 := predicate.(type) {

		case func(string) bool:
			if ss, ok := collection.([]string); ok {
				result := make([]string, 0, len(ss))
				for _, s := range ss {
					if f1(s) {
						result = append(result, s)
					}
				}
				return result
			}

		case func(int) bool:
			if ns, ok := collection.([]int); ok {
				result := make([]int, 0, len(ns))
				for _, n := range ns {
					if f1(n) {
						result = append(result, n)
					}
				}
				return result
			}

		case func(float64) bool:
			if fs, ok := collection.([]float64); ok {
				result := make([]float64, 0, len(fs))
				for _, f := range fs {
					if f1(f) {
						result = append(result, f)
					}
				}
				return result
			}
		}
	}

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
	resultValue := reflect.Zero(resultType)

	for i := 0; i < length; i++ {
		v := collectionValue.Index(i)
		if predicateValue.Call([]reflect.Value{v})[0].Bool() {
			resultValue = reflect.Append(resultValue, v)
		}
	}

	return resultValue.Interface()
}
