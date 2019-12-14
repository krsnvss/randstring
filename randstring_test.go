package randstring

import (
	"reflect"
	"testing"
)

func TestMakeRandomString(t *testing.T) {
	r := MakeRandomString(8)
	if len(r) != 8 {
		t.Error("String length is not enough!")
	}
	if reflect.TypeOf(r).Kind() != reflect.String {
		t.Error("Not a string!")
	}
	r1 := MakeRandomString(8)
	if r == r1 {
		t.Error("String is not random!")
	}
}
