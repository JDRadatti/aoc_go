package day03

import (
    "github.com/JDRadatti/aoc_go/pkg/graphs"
)

// https://adventofcode.com/2023/day/3
func SolutionA(schematic []byte) int {
    // convert to graph where symbol is edge and number is node
    // add nodes for each graph with an edge
    // adjaceny matrix
    // 
    graph := graphs.Graph{}
    graph.InitFromSchematic(schematic)


    sum := 0
    for n := range graph.Nodes {
        sum += n.Value
    }
	return sum
}
