package day08

import (
	"bytes"
	"github.com/JDRadatti/aoc_go/pkg/graphs"
	"github.com/JDRadatti/aoc_go/pkg/utils"
)

// https://adventofcode.com/2023/day/8
func SolutionA(input []byte) int {
	instructions, graph := Init(input)
	return graph.Search(instructions, "AAA", "ZZZ")
}

// https://adventofcode.com/2023/day/8
func SolutionB(input []byte) int {
	instructions, graph := Init(input)
	starts := []string{}
	for node := range graph {
		if node[len(node)-1] == 'A' {
			starts = append(starts, node)
		}
	}

	steps := []int{}
	for _, node := range starts {
		steps = append(steps, graph.Search(instructions, node, "**Z"))
	}

	return utils.LCM(steps)
}

func Init(input []byte) ([]int, graphs.LRGraph) {
	nodes := bytes.Split(input, []byte("\n"))
	graph := graphs.LRGraph{}
	instructions := Instructions(nodes[0])

	for _, node := range nodes {
		graph.AddNode(node)
	}
	return instructions, graph
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
