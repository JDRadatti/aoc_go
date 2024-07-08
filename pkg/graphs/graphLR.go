package graphs

import "regexp"

type LRNode [2]string // LRNode[0] = left; LRNode[1] = right

type LRGraph map[string]LRNode

func (g LRGraph) GetNode(id string) LRNode { return g[id] }

func (g LRGraph) AddNode(node []byte) bool {
	re, _ := regexp.Compile("[A-Z]{3}")
	nodes := re.FindAll(node, -1)
	if len(nodes) == 3 {
		node := string(nodes[0])
		left := string(nodes[1])
		right := string(nodes[2])
		g[node] = [2]string{left, right}
		return true
	}
	return false
}

func (g LRGraph) Search(instructions []int, start, end string) int {
	i, steps := 0, 0
    currIndex, endIndex := string(start), string(end)
	for currIndex != endIndex {
		currIndex = g[currIndex][instructions[i]]
		i = (i + 1) % len(instructions)
		steps++
	}
	return steps
}
