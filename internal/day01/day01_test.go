package day01

import (
	"testing"
)

func TestDay01a(t *testing.T) {
	inputs := [][]byte{
		[]byte("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\n"),
		[]byte("abc2\n"),
	}
	expected := [...]int{
		142,
		22,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionA(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}

func TestDay01b(t *testing.T) {
	inputs := [][]byte{
		[]byte("sdfmjthreeldkm\n",), 
		[]byte("onetwo\n",), 
		[]byte("fiveight\n",), 
		[]byte("threeoneight\n",), 
        []byte(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`),
	}
	expected := [...]int{
        33,
        12,
        58,
        38,
		281,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionB(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}
