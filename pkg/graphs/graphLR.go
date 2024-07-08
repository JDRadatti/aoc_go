package graphs

import (
	"fmt"
	"regexp"
	"strings"
)

const allowedChars string = "[A-Z1-9]"
const allowedLength int = 3

type LRNode [2]string // LRNode[0] = left; LRNode[1] = right

type LRGraph map[string]LRNode

func (g LRGraph) GetNode(id string) LRNode { return g[id] }

func (g LRGraph) AddNode(node []byte) bool {
	re := regexp.MustCompile(fmt.Sprintf("%s{%d}", allowedChars, allowedLength))
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
	for !g.NodeMatch(start, end) {
		start = g.Next(instructions[i], start)
		i = (i + 1) % len(instructions)
		steps++
	}
	return steps
}

func (g LRGraph) Next(instruction int, start string) string {
	return g[start][instruction]
}

// NodeMatch determines if node N1 is the same as node N2
// NodeMatch allows * characters in N2 to match with any allowed characters
func (g LRGraph) NodeMatch(n1, n2 string) bool {
	s := strings.ReplaceAll(n2, "*", allowedChars)
	re, _ := regexp.Compile(s)
	return re.MatchString(n1)
}
