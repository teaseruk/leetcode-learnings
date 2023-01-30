package palindrome

import (
	"reflect"
	"testing"
)

func Test_case1(t *testing.T) {
	result := isPalindrome(121)
	expected := true

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case2(t *testing.T) {
	result := isPalindrome(-121)
	expected := false

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case3(t *testing.T) {
	result := isPalindrome(10)
	expected := false

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case4(t *testing.T) {
	result := isPalindrome(1234567890987654321)
	expected := true

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case5(t *testing.T) {
	result := isPalindrome(123456789987654321)
	expected := true

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case6(t *testing.T) {
	result := isPalindrome(12345678998765432)
	expected := false

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}
}

func Test_case7(t *testing.T) {

	result := isPalindrome(1000030001)
	expected := false

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wanted %v got %v", expected, result)
	}

}
