package day09

import (
	"bytes"
	"github.com/JDRadatti/aoc_go/pkg/utils"
)

// https://adventofcode.com/2023/day/
func SolutionA(input []byte) int {
	history := Init(input)
	sum := 0
	for _, h := range history {
		lastRow := Base(h)
		nextRow := Next(lastRow)
		sum += nextRow[0]
	}
	return sum
}

// https://adventofcode.com/2023/day/
func SolutionB(input []byte) int {
	return 0
}

func Init(input []byte) [][]int {
	levels := bytes.Split(input, []byte("\n"))
	history := [][]int{}
	for _, level := range levels {
		if len(level) <= 1 {
			continue
		}
		var levelHistory []int
		for _, i := range bytes.Split(level, []byte(" ")) {
			levelHistory = append(levelHistory, utils.BAtoI(i))
		}
		history = append(history, levelHistory)
	}
	return history
}

// RightMost finds the base of the history pyramid
// After finding the base, returns the rightmost column
// a b c d -- history
//
//	e f g
//	 h i -- base when all are 0
//
// Return [d,g,i]
func Base(history []int) []int {
	levels := [][]int{history}
	var all0 bool
	for !all0 {
		all0 = true
		var currLevel []int
		for i := 1; i < len(levels[len(levels)-1]); i++ {
			l := levels[len(levels)-1]
			currLevel = append(currLevel, l[i]-l[i-1])
			if l[i]-l[i-1] != 0 {
				all0 = false
			}
		}
		levels = append(levels, currLevel)
	}

	var lastRow []int
	for _, l := range levels {
		if len(l) == 0 {
			lastRow = append(lastRow, 0)
		} else {
			lastRow = append(lastRow, l[len(l)-1])
		}
	}
	return lastRow
}

// Next returns the next LastRow of the pyramid
// a b c d
//
//	e f g
//	 h i
//
// If input = [c, d, h] -> return [d, g, i]
func Next(lastRow []int) []int {
	nextRow := make([]int, len(lastRow))
	copy(nextRow, lastRow)
	for i := len(lastRow) - 3; i >= 0; i-- {
		nextRow[i] = lastRow[i] + nextRow[i+1]
	}
	return nextRow
}
