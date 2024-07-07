package day07


import "testing"

func TestDay07a(t *testing.T) {
	inputs := [][]byte{
		[]byte(`
`),
	}
	expected := [...]int{
        0,
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
		[]byte(`
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
