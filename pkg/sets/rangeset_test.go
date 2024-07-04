package rangeset

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//[]rangeset.Range([]rangeset.Range{rangeset.Range{0, 5}})
//rangeset.Ranges(rangeset.Ranges{rangeset.Range{0, 5}})

type InputConfig[K any] struct {
	Ranges   []Range
	Input    [][]int
	Expected []K
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

// func TestRangeSetAddSingleRanges(t *testing.T) {
// 	rangeSet := NewRangeSet()
// 	// add range 0, 5
// 	// add range 7,9
// 	rangeSet.Add(5, 5)
// 	assert.Equal(t, 0, 0, "Equality of ranges")
//  what if both end and start are in the same gutter?
//  what if they one is in the first gutter and last is in the last
// }
//

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
