package yogofn

import "reflect"

func typecheck(unknown interface{}, possibleKind ...reflect.Kind) {

	kind := reflect.TypeOf(unknown).Kind()

	for _, t := range possibleKind {
		if t == kind {
			return
		}
	}
	panic("wrong type")
}
