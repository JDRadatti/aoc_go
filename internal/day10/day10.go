package day10

import (
	"bytes"
	"github.com/JDRadatti/aoc_go/pkg/graphs"
	"slices"
)

// https://adventofcode.com/2023/day/
func SolutionA(input []byte) int {
	graph, start := InitGraph(input)
	return graph.BFS(start)
}

// https://adventofcode.com/2023/day/
func SolutionB(input []byte) int {
	return 0
}

type Direction [2]int // X, Y
var Directions = struct {
	East  Direction
	West  Direction
	North Direction
	South Direction
}{
	East:  Direction{1, 0},
	West:  Direction{-1, 0},
	North: Direction{0, -1},
	South: Direction{0, 1},
}

var Pipes map[byte][]Direction = map[byte][]Direction{
	'|': {Directions.North, Directions.South},
	'-': {Directions.East, Directions.West},
	'L': {Directions.North, Directions.East},
	'J': {Directions.North, Directions.West},
	'7': {Directions.South, Directions.West},
	'F': {Directions.South, Directions.East},
	'S': {Directions.South, Directions.East, Directions.West, Directions.North},
}

func OppositeDirection(d Direction) Direction {
	switch d {
	case Directions.South:
		return Directions.North
	case Directions.North:
		return Directions.South
	case Directions.East:
		return Directions.West
	case Directions.West:
		return Directions.East
	}
	return Direction{}

}

// ConnectingPipe checks the Node in Direction and determines
// if it is a valid connecting pipe.
func ConnectingPipe(d Direction, symbol byte) bool {
	oppositeDirection := OppositeDirection(d)
	if directions, ok := Pipes[symbol]; ok {
		if slices.Contains(directions, oppositeDirection) {
			return true
		}
	}
	return false
}

func InitGraph(input []byte) (graphs.Graph, graphs.Node) {
	split := bytes.Split(input, []byte("\n"))
	graph := graphs.Graph{}
	var start graphs.Node
	for y := 0; y < len(split); y++ {
		for x := 0; x < len(split[y]); x++ {
			symbol := split[y][x]
			if symbol == '.' {
				continue
			}
			currNode := graphs.Node{x, y}
			if symbol == 'S' {
				start = currNode
			}
			if dirs, ok := Pipes[symbol]; ok {
				for _, d := range dirs {
					nextX := x + d[0]
					nextY := y + d[1]
					if nextX >= 0 && nextY >= 0 &&
						nextX < len(split[nextY]) && nextY < len(split) {
						if ConnectingPipe(d, split[nextY][nextX]) {
							nextNode := graphs.Node{nextX, nextY}
							graph.AddEdge(currNode, nextNode)
						}
					}
				}
			}
		}
	}
	return graph, start
}
