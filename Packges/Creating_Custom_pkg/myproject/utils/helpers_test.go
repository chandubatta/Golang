package utils

import (
	"testing"
)

//this is the one way of unit testing
/*func TestAdd(t *testing.T) {
	result := Add(30, 40)
	expected := 70
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}*/

//func TestAdd(t *testing.T){

// }

//This is the another way of unit testing

func TestAdd(t *testing.T) {

	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{10, 20, 30},
		{5, 5, 10},
		{100, 200, 300},
	}

	for _, test := range tests {

		result := Add(test.a, test.b)

		if result != test.expected {
			t.Errorf(
				"Expected %d but got %d",
				test.expected,
				result,
			)
		}
	}
}
