package utils

import "testing"



func TestLCM(t *testing.T) {
    inputs := [][]int{
        {2, 3},
        {2, 201, 5, 10},
    }
    expected := []int{
        6, 2010,
    }
    for i, input := range inputs {
        result := LCM(input)
        if result != expected[i] {
            t.Errorf("LCM test failed on input %d. Expected %d but got %d", i, expected[i], result)
        }
    }
}

func TestGCD(t *testing.T) {
    inputs := [][]int{
        {2, 201, 5, 10},
        {2, 3},
        {2, 4, 6, 8, 10},
        {0, 4, 6, 8, 10},
        {48, 18},
        {48, 18, 48 * 6},
        {11, 11, 11, 11, 11},
        {100, 200, 10, 2, 4},
    }
    expected := []int{
        1, 1, 2, 2, 6, 6, 11, 2,
    }
    for i, input := range inputs {
        if GCD(input) != expected[i] {
            t.Errorf("GCD test failed on input %d", i)
        }
    }
}
