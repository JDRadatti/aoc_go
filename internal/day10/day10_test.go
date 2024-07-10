package day10

import (
	"testing"
)

func TestDay10a(t *testing.T) {
	inputs := [][]byte{
		[]byte(`.....
.S-7.
.|.|.
.L-J.
.....`),
		[]byte(`7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`),
	}
	expected := [...]int{
		4, 8,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionA(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}

func TestDay10b(t *testing.T) {
	inputs := [][]byte{
		[]byte(``),
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

