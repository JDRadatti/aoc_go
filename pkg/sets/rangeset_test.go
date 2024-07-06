package rangeset

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//[]rangeset.Range([]rangeset.Range{rangeset.Range{0, 5}})
//rangeset.Ranges(rangeset.Ranges{rangeset.Range{0, 5}})

type InputConfig[K any] struct {
	Ranges      []Range
	Input       [][]int
	InputRanges []Ranges
	Expected    []K
}

func (config *InputConfig[int]) RunSearch(t *testing.T) {
	rangeSet := NewRangeSet()
	rangeSet.Ranges = config.Ranges
	for i, input := range config.Input {

		assert.Equal(t, config.Expected[i], rangeSet.Search(input[0]),
			fmt.Sprintf("failed at test %d", i))
	}
}

func (config *InputConfig[Ranges]) RunAdd(t *testing.T) {
	rangeSet := NewRangeSet()
	rangeSet.Ranges = config.Ranges
	for i, input := range config.Input {
		start, end := input[0], input[1]
		rangeSet.Add(start, end)
		assert.Equal(t, config.Expected[i], rangeSet.Ranges,
			fmt.Sprintf("failed at test %d", i))
	}
}

func (config *InputConfig[Ranges]) RunNewFromRanges(t *testing.T) {
	for i := range config.InputRanges {
		rangeSet := NewRangeSetFromRanges(config.InputRanges[i])
		assert.Equal(t, config.Expected[i], rangeSet.Ranges,
			fmt.Sprintf("failed at test %d", i))
	}
}

// Note: Uses InputRange i and i+1 as input for all even i in InputRange
func (config *InputConfig[Ranges]) RunIntersection(t *testing.T) {
	j := 0
	for i := 1; i < len(config.InputRanges); i += 2 {
		rangeSet0 := NewRangeSetFromRanges(config.InputRanges[i-1])
		rangeSet1 := NewRangeSetFromRanges(config.InputRanges[i])
		outputSet := rangeSet0.Intersection(rangeSet1)
		assert.Equal(t, config.Expected[j], outputSet.Ranges,
			fmt.Sprintf("failed at test %d", i))
		j++
	}
}

func TestRangeSetSearch(t *testing.T) {
	// Test assumes non-decreasing first terms in ranges
	config := InputConfig[int]{
		Ranges:   []Range{{0, 1}, {4, 8}, {12, 15}},
		Input:    [][]int{{20}, {-1}, {0}, {5}, {3}, {9}, {8}, {13}},
		Expected: []int{2, 0, 0, 1, 1, 2, 1, 2},
	}
	config.RunSearch(t)
	config = InputConfig[int]{
		Ranges:   []Range{{0, 1}},
		Input:    [][]int{{0}, {1}, {-2}, {13}},
		Expected: []int{0, 0, 0, 0},
	}
	config.RunSearch(t)

	config = InputConfig[int]{
		Ranges:   []Range{},
		Input:    [][]int{{0}, {13}, {-13}},
		Expected: []int{0, 0, 0, 0},
	}
	config.RunSearch(t)
}

func TestRangeSetAddBasic(t *testing.T) {
	config := InputConfig[Ranges]{
		Ranges: []Range{},
		Input: [][]int{
			{0, 5}, // start, end
			{7, 9},
			{2, 4},  // shouldn't do anything
			{7, 9},  // shouldn't do anything
			{4, 6},  // extend range 1
			{-1, 8}, // connect ranges
			{15, 17},
			{20, 37},
		},

		Expected: []Ranges{
			{{0, 5}},
			{{0, 5}, {7, 9}},
			{{0, 5}, {7, 9}},
			{{0, 5}, {7, 9}},
			{{0, 6}, {7, 9}},
			{{-1, 9}},
			{{-1, 9}, {15, 17}},
			{{-1, 9}, {15, 17}, {20, 37}},
		},
	}
	config.RunAdd(t)

	config = InputConfig[Ranges]{
		Ranges: []Range{},
		Input: [][]int{
			{50, 55},
			{40, 49}, // ax
			{49, 50}, // merge
			{56, 60}, // cz
			{30, 35},
			{38, 38},
			{64, 65},
			{36, 63}, // cx
			{80, 85},
			{29, 79}, // ax
			{0, 4},
			{5, 85}, // cy
			{90, 95},
			{100, 105},
			{107, 150},
			{1, 107}, // by
		},

		Expected: []Ranges{
			{{50, 55}},
			{{40, 49}, {50, 55}},
			{{40, 55}},
			{{40, 55}, {56, 60}},
			{{30, 35}, {40, 55}, {56, 60}},
			{{30, 35}, {38, 38}, {40, 55}, {56, 60}},
			{{30, 35}, {38, 38}, {40, 55}, {56, 60}, {64, 65}},
			{{30, 35}, {36, 63}, {64, 65}},
			{{30, 35}, {36, 63}, {64, 65}, {80, 85}},
			{{29, 79}, {80, 85}},
			{{0, 4}, {29, 79}, {80, 85}},
			{{0, 4}, {5, 85}},
			{{0, 4}, {5, 85}, {90, 95}},
			{{0, 4}, {5, 85}, {90, 95}, {100, 105}},
			{{0, 4}, {5, 85}, {90, 95}, {100, 105}, {107, 150}},
			{{0, 150}},
		},
	}
	config.RunAdd(t)

}

func TestNewRangeSetFromRanges(t *testing.T) {
	config := InputConfig[Ranges]{
		Ranges: []Range{},
		InputRanges: []Ranges{
			{{0, 4}, {5, 5}, {10, 15}},            // Basic use case
			{{0, 4}, {5, 50}, {10, 15}, {10, 20}}, // Overlapping ranges
			{{5, 50}, {10, 15}, {0, 4}, {6, 20}},  // Out of order
			{{5, 50}},                             // Single
			{{5, 50}, {10, 20}, {20, 30}},         // First dwarfs all others
			{{50, 5}, {15, 10}, {4, 0}, {6, 20}},  // Really out of order
		},

		Expected: []Ranges{
			{{0, 4}, {5, 5}, {10, 15}},
			{{0, 4}, {5, 50}},
			{{0, 4}, {5, 50}},
			{{5, 50}},
			{{5, 50}},
			{{0, 4}, {5, 50}},
		},
	}
	config.RunNewFromRanges(t)
}

func TestIntersecton(t *testing.T) {
	config := InputConfig[Ranges]{
		InputRanges: []Ranges{
			{{0, 10}}, // Test 1
			{{1, 4}},  // Test 1

			{{0, 10}, {20, 30}, {40, 50}, {60, 70}}, // Test 2
			{{80, 90}},                              // Test 2

			{{0, 10}, {20, 30}, {40, 50}, {60, 70}}, // Test 2
			{{5, 45}},                               // Test 2

			{{5, 45}},
			{{0, 10}, {20, 30}, {40, 50}, {60, 70}},

			{{5, 45}, {46, 100}},
			{{5, 45}, {46, 100}},

            {}, 
            {},
		},

		Expected: []Ranges{
			{{1, 4}},                      // test 1 expected
			nil,                           // test 2
			{{5, 10}, {20, 30}, {40, 45}}, // test 3
			{{5, 10}, {20, 30}, {40, 45}},
			{{5, 45}, {46, 100}},
            nil,
		},
	}
	config.RunIntersection(t)
}
