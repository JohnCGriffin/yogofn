package yogofn

import "reflect"

// float64, int and string slices are common, so
// avoid reflecting on every element if possible.
func genericSlice(collection interface{}) []interface{} {

	result := make([]interface{}, 0)

	if fs, ok := collection.([]float64); ok {
		for _, f := range fs {
			result = append(result, f)
		}
		return result
	}

	if ns, ok := collection.([]int); ok {
		for _, n := range ns {
			result = append(result, n)
		}
		return result
	}

	if ss, ok := collection.([]string); ok {
		for _, s := range ss {
			result = append(result, s)
		}
		return result
	}

	// resort to per-element reflection
	cVal := reflect.ValueOf(collection)
	length := cVal.Len()

	for i := 0; i < length; i++ {
		result = append(result, cVal.Index(i).Interface())
	}

	return result
}

func genericSlices(collections []interface{}) [][]interface{} {
	result := make([][]interface{}, 0)
	for i := 0; i < len(result); i++ {
		result = append(result, genericSlice(collections[i]))
	}
	return result
}
