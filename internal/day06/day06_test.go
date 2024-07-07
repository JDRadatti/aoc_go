package day06

import "testing"

func TestDay06a(t *testing.T) {
	inputs := [][]byte{
		[]byte(`Time:      7  15   30
Distance:  9  40  200
`),
	}
	expected := [...]int{
        288,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionA(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}

func TestDay06b(t *testing.T) {
	inputs := [][]byte{
		[]byte(`Time:      7  15   30
Distance:  9  40  200
`),
	}
	expected := [...]int{
		0,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionB(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}
