package graphs

import "regexp"

const offset int = int('A')
const numNodes int = 26 // ID contsrained to Alphabet

type LRNode [2]int // LRNode[0] = left; LRNode[1] = right

type LRGraph [numNodes]LRNode

func (g *LRGraph) GetNode(id []byte) LRNode { return g[g.Index(id)] }

func (g *LRGraph) Index(id []byte) int {
	return int(id[0]) - offset
}

func (g *LRGraph) AddNode(node []byte) bool {
	re, _ := regexp.Compile("[A-Z]{3}")
	nodes := re.FindAll(node, -1)
	if len(nodes) == 3 {
		node := g.Index(nodes[0])
		left := g.Index(nodes[1])
		right := g.Index(nodes[2])
		g[node] = [2]int{left, right}
		return true
	}
	return false
}

func (g *LRGraph) Search(instructions []int, start, end []byte) int {
	i, steps := 0, 0
    currIndex, endIndex := g.Index(start), g.Index(end)
	for currIndex != endIndex {
		currIndex = g[currIndex][instructions[i]]
		i = (i + 1) % len(instructions)
		steps++
	}
	return steps
}
