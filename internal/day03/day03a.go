package day03

import (
    "github.com/JDRadatti/aoc_go/pkg/graphs"
)

// https://adventofcode.com/2023/day/3
func SolutionA(schematic []byte) int {
    graph := graphs.Graph{}
    graph.InitFromSchematic(schematic)


    sum := 0
    for n := range graph.Nodes {
        sum += n.Value
    }
	return sum
}

// https://adventofcode.com/2023/day/3
func SolutionB(schematic []byte) int {
    graph := graphs.Graph{}
    graph.InitFromSchematic(schematic)


    sum := 0
    for e := range graph.Edges {
        sum += graph.Edges[e].GearRatio()
    }
	return sum
}
