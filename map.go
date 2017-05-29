package yogofn

import (
	"reflect"
)

// Map takes a conversion function (mapper) and a collection(s) (i.e. an Array or Slice),
// resulting in an output slice where each element is transformed.
// The mapping function, the input collection and the output collection are all
// typed as interface{}. The result is almost always going to require a cast.
//
// For instance, uppercasing names
//
//   upNames := Map(func(s string){ return strings.ToUpper(s) },names).([]string)
//
// Note that the number of collections in the uppercasing example is one.  However,
// any number of input lists may be used.  Each input is iterated until one of them
// is empty.  So, given and and low temperature lists, a temperature range might
// be output.
//
//   tempRanges := Map(func(a,b float64) float64 {return math.Abs(a-b) },
//						lows,highs).([]float64)
//
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
