package day03

import "testing"

func TestDay03a(t *testing.T) {
	inputs := [][]byte{
		[]byte(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`),
		[]byte(`...
1.1
.*.
1.1
`),
	}
	expected := [...]int{
		4361,
        4,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionA(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}

func TestDay03b(t *testing.T) {
	inputs := [][]byte{
		[]byte(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`),
	}
	expected := [...]int{
		467835,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionB(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}
