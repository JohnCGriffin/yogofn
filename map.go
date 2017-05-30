package yogofn

import (
	"reflect"
)

// Map takes a mapper function (func(T1,..Tn) Tresult) and one or more
// collection. (i.e. Array or Slice), resulting in an output slice where
// each element is transformed. The mapping function, the input collection
// and the output collection are all typed as interface{}. The result
// requires a cast.
//
// For instance, converting a collection of ints to string representations
//
//   nums := []int {1,2,3,4}
//   numStrings := Map(fmt.Sprint,nums).([]string)
//
// or uppercasing names
//
//   upNames := Map(strings.ToUpper,names).([]string)
//
// Note that the number of collections in the uppercasing example is one.  However,
// any number of input lists may be used.  Each input is iterated until one of them
// is empty.  Thus, given low and high temperature lists, a temperature range might
// be output.
//
//   ranges := Map(func(l,h float64) float64 {return h-l },lows,highs).([]float64)
//
func Map(mapper interface{}, collections ...interface{}) interface{} {

	typecheck(mapper, reflect.Func)
	mapperValue := reflect.ValueOf(mapper)

	if len(collections) < 1 {
		panic("Map requires at least one collection")
	}

	// common optimizations
	if len(collections) == 1 {

		collection := collections[0]

		switch f1 := mapper.(type) {

		case func(string) int:
			if ss, ok := collection.([]string); ok {
				result := make([]int, 0, len(ss))
				for _, s := range ss {
					result = append(result, f1(s))
				}
				return result
			}

		case func(int) int:
			if ns, ok := collection.([]int); ok {
				result := make([]int, 0, len(ns))
				for _, n := range ns {
					result = append(result, f1(n))
				}
				return result
			}

		case func(float64) float64:
			if fs, ok := collection.([]float64); ok {
				result := make([]float64, 0, len(fs))
				for _, f := range fs {
					result = append(result, f1(f))
				}
				return result
			}
		}
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
	resultValue := reflect.Zero(resultType)

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
