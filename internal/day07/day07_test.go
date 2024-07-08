package day07

import "testing"

func TestDay07a(t *testing.T) {
	inputs := [][]byte{
		[]byte(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`),
	}
	expected := [...]int{
		6440,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionA(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}

func TestDay07b(t *testing.T) {
	inputs := [][]byte{
		[]byte(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`),
	}
	expected := [...]int{
		5905,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionB(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}
