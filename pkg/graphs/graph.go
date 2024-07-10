package graphs

import "github.com/JDRadatti/aoc_go/pkg/queue"

type Node [2]int // X, Y

type Graph struct {
	Edges map[Node][]Node
	Size  int
}

func (g *Graph) AddEdge(a, b Node) {
	if g.Size == 0 {
		g.Edges = map[Node][]Node{}
	}
	g.Edges[a] = append(g.Edges[a], b)
	g.Size++
}

func (g *Graph) Neighbors(a Node) []Node {
	return g.Edges[a]
}

// BFS runs a breath first search on graph g and returns number of levels
func (g Graph) BFS(start Node) int {
	visited := map[Node]struct{}{}
	nodeStack := queue.Queue[Node]{}
	nodeStack.Enqueue(start)
	levels := -1 
	for !nodeStack.Empty() {
		for i := 0; i < len(nodeStack); i++ {
			curr, ok := nodeStack.Dequeue()
			visited[curr] = struct{}{}
			for _, neighbor := range g.Neighbors(curr) {
				if _, ok = visited[neighbor]; ok {
					continue // skip already visited nodes
				}
				nodeStack.Enqueue(neighbor)
			}
		}
		levels++
	}
	return levels
}

// TODO: Update Schematic day to use new graph
//import "bytes"
//
// type Graph struct {
// 	Nodes map[Node]struct{}
// 	Edges []Edge
// }
//
// type Node struct {
// 	visited bool
// 	Value   int
// 	Id      int
// }
//
// type Edge struct {
// 	neighbors map[Node]struct{}
// 	IsGear    bool
// }
//
// func (e *Edge) GearRatio() int {
// 	if !e.IsGear {
// 		return 0
// 	}
//
// 	gearRatio := 1
// 	for neighbor := range e.neighbors {
// 		gearRatio *= neighbor.Value
// 	}
// 	return gearRatio
// }
//
// func (edge *Edge) addNeighbor(node Node) {
// 	if len(edge.neighbors) == 0 {
// 		edge.neighbors = map[Node]struct{}{}
// 	}
// 	edge.neighbors[node] = struct{}{}
// }
//
// func (node *Node) updateID(value byte) {
// 	node.Value = node.Value*10 + int(value) - 48
// }
//
// func (graph *Graph) addEdge(edge Edge) {
// 	graph.Edges = append(graph.Edges, edge)
// }
//
// func (graph *Graph) addNode(node Node) {
// 	if len(graph.Nodes) == 0 {
// 		graph.Nodes = map[Node]struct{}{}
// 	}
// 	graph.Nodes[node] = struct{}{}
// }
//
// var directions [][]int = [][]int{
// 	{1, 1},
// 	{-1, -1},
// 	{1, -1},
// 	{-1, 1},
// 	{1, 0},
// 	{0, 1},
// 	{-1, 0},
// 	{0, -1},
// }
//
// func (graph *Graph) InitFromSchematic(schematic []byte) {
// 	var split [][]byte = bytes.Split(schematic, []byte("\n"))
// 	var nodes [][]*Node
// 	var nilNode *Node = &Node{Value: 0}
// 	id := 0
//
// 	// First pass, init 2D array of Nodes
// 	// Edge is a node with id = -1
// 	for row := 0; row < len(split); row++ {
// 		if len(split[row]) == 0 {
// 			continue
// 		}
//
// 		nodes = append(nodes, []*Node{})
// 		currNode := &Node{}
// 		for col := 0; col < len(split[row]); col++ {
// 			curr := split[row][col]
// 			if curr >= 48 && curr <= 57 {
// 				currNode.updateID(curr)
// 				currNode.Id = id
// 				id += 1
// 				nodes[row] = append(nodes[row], currNode)
// 			} else if curr == 46 { // curr byte is a "."
// 				currNode = &Node{}
// 				nodes[row] = append(nodes[row], nilNode)
// 			} else { // other sympol / edge
// 				currNode = &Node{}
// 				nodes[row] = append(nodes[row], &Node{Value: -1})
// 			}
// 		}
// 	}
//
// 	for row := 0; row < len(nodes); row++ {
// 		for col := 0; col < len(nodes[row]); col++ {
//
// 			if nodes[row][col].Value >= 0 {
// 				continue
// 			}
//
// 			edge := Edge{}
// 			for dir := 0; dir < len(directions); dir++ {
// 				x := directions[dir][0]
// 				y := directions[dir][1]
// 				if row+y >= 0 &&
// 					col+x >= 0 &&
// 					row+y < len(nodes) &&
// 					col+x < len(nodes[row]) {
// 					neighbor := nodes[row+y][col+x]
// 					if neighbor.Value > 0 { // found node
// 						edge.addNeighbor(*neighbor)
// 						graph.addNode(*neighbor)
// 					}
// 				}
// 			}
// 			if split[row][col] == '*' && len(edge.neighbors) == 2 {
// 				edge.IsGear = true
// 			}
// 			if len(edge.neighbors) > 0 {
// 				graph.addEdge(edge)
// 			}
//
// 		}
// 	}
// }
