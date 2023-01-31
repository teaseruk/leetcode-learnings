package valid_parantheses

import (
	"reflect"
	"testing"
)

func Test_case1(t *testing.T) {
	result := isValid("()")
	expected := true

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case2(t *testing.T) {
	result := isValid("()[]{}")
	expected := true

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case3(t *testing.T) {
	result := isValid("(]")
	expected := false

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case4(t *testing.T) {
	result := isValid("{([()])}")
	expected := true

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case5(t *testing.T) {
	result := isValid("((")
	expected := false

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case6(t *testing.T) {
	result := isValid("()")
	expected := true

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}
