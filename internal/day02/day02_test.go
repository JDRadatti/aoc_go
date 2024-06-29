package day02

import "testing"

func TestDay02a(t *testing.T) {
	inputs := [][]byte{
		[]byte(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`),
		[]byte(`Game 1: 3 blue, 4 red; 8 red; 2 green
`),
		[]byte(`Game 1: 3 blue, 4 red; 8 red; 2 green
Game 2: 1 blue, 2 green; 3 green, 15 blue, 1 red; 1 green, 1 blue
`),
        []byte(`Game 1: 10 blue, 10 blue
`),
        []byte(`Game 1: 4 blue, 10 blue; 13 red; 12 red
`),
	}
	expected := [...]int{
		8,
		1,
        1,
        0,
        1, 
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionA(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}
