package day06

import (
	"github.com/JDRadatti/aoc_go/pkg/utils"
	"regexp"
    "bytes"
)

// https://adventofcode.com/2023/day/6
func SolutionA(input []byte) int {
	re := regexp.MustCompile(`\d+`)
	var digits []int
	for _, i := range re.FindAll(input, -1) {
		digits = append(digits, utils.BAtoI(i))
	}
	offset := len(digits) / 2
	velocity := 1 // millimeters per millisecond
	sum := 1
	for i := 0; i < offset; i++ {
		time := digits[i] // milliseconds
		distance := digits[i+offset]
		first, last := FirstAndLastTime(time, distance, velocity)
		sum *= (last - first + 1)
	}
	return sum
}


// https://adventofcode.com/2023/day/6
func SolutionB(input []byte) int {
    input = bytes.ReplaceAll(input, []byte(" "), []byte(""))
    return SolutionA(input)
}

func MaxDistance(holdtime, time, velocity int) int {
	return holdtime * velocity * (time - holdtime)
}

func FirstAndLastTime(time, distance, velocity int) (int, int) {
	var first, last int
	for i := 1; i < time; i++ {
		maxDist := MaxDistance(i, time, velocity)
		if maxDist > distance {
			first = i
			break
		}
	}
	for i := time - 1; i > 0; i-- {
		maxDist := MaxDistance(i, time, velocity)
		if maxDist > distance {
			last = i
			break
		}
	}
	return first, last
}
