package romaninteger

import (
	"reflect"
	"testing"
)

func Test_case1(t *testing.T) {
	result := romanToInt("III")
	expected := 3

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case2(t *testing.T) {
	result := romanToInt("LVIII")
	expected := 58

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case3(t *testing.T) {
	result := romanToInt("MCMXCIV")
	expected := 1994

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}
