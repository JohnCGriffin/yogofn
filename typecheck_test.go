package yogofn

import (
	"reflect"
	"testing"
)

func TestTypeCheckGood1(t *testing.T) {
	typecheck("some text", reflect.String)
}

func TestTypeCheckGood2(t *testing.T) {
	typecheck("some text", reflect.String, reflect.Float64)
}

func TestTypeCheckBad0(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	typecheck(123)
}

func TestTypeCheckBad2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	typecheck(123, reflect.Slice, reflect.Float32)
}
