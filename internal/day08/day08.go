package day08

import (
	"bytes"
	"github.com/JDRadatti/aoc_go/pkg/graphs"
)

// https://adventofcode.com/2023/day/8
func SolutionA(input []byte) int {
	nodes := bytes.Split(input, []byte("\n"))
	graph := graphs.LRGraph{}
	instructions := Instructions(nodes[0])

	for _, node := range nodes {
		graph.AddNode(node)
	}
	return graph.Search(instructions, "AAA", "ZZZ")
}

// https://adventofcode.com/2023/day/8
func SolutionB(input []byte) int {
	return 0
}

func Instructions(input []byte) []int {
	instructions := []int{}
	for _, b := range input {
		switch b {
		case 'L':
			instructions = append(instructions, 0)
		case 'R':
			instructions = append(instructions, 1)
		default:
			panic("invalid instruction")
		}
	}
	return instructions
}
