package functions

import (
	"reflect"
	"testing"
)

func TestAddTwo(t *testing.T) {
	got := addtwo(100, 300)
	if !reflect.DeepEqual(got, 400) {
		t.Fail()
	}
}
func TestAddStrings(t *testing.T) {
	got := addstrings("a", "b")
	if !reflect.DeepEqual(got, "ab") {
		t.Fail()
	}
}
func TestMultiply(t *testing.T) {
	got := multiply(50, 3)
	if !reflect.DeepEqual(got, 150) {
		t.Fail()
	}
}
func TestSubtract(t *testing.T) {
	got := subtract(10, 5)
	if !reflect.DeepEqual(got, 5) {
		t.Fail()
	}
}
func TestDivide(t *testing.T) {
	got := divide(10, 5)
	if !reflect.DeepEqual(got, 2) {
		t.Fail()
	}
}
