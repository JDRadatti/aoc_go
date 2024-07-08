package day08

import "testing"

func TestDay08a(t *testing.T) {
	inputs := [][]byte{
		[]byte(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`),
        []byte(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`),
    }
	expected := [...]int{
		2, 
        6,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionA(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}

func TestDay08b(t *testing.T) {
	inputs := [][]byte{
		[]byte(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`),
        []byte(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`),
    }
	expected := [...]int{
		0, 
        0,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionB(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}
