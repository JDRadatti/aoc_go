package day09

import (
	"testing"
)

func TestDay09a(t *testing.T) {
	inputs := [][]byte{
		[]byte(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`),
	}
	expected := [...]int{
		114,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionA(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}

func TestDay09b(t *testing.T) {
	inputs := [][]byte{
		[]byte(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`),
	}
	expected := [...]int{
		2,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionB(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}
