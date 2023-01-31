package longest_common_prefix

import (
	"reflect"
	"testing"
)

func Test_testCase1(t *testing.T) {
	result := longestCommonPrefix([]string{"flower", "flow", "flight"})
	expected := "fl"

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_testCase2(t *testing.T) {
	result := longestCommonPrefix([]string{"dog", "racecar", "car"})
	expected := ""

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_testCase3(t *testing.T) {
	result := longestCommonPrefix([]string{"a"})
	expected := ""

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_testCase4(t *testing.T) {
	result := longestCommonPrefix([]string{"ab", "a"})
	expected := "a"

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}
