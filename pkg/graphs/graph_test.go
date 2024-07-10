package graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type InputConfig[K any] struct {
	Edges       map[Node][]Node
	Expected    K
    Start Node
}

func (config *InputConfig[int]) RunBFS(t *testing.T) {
    rangeSet := Graph{Edges: config.Edges}
    output := rangeSet.BFS(config.Start)
    assert.Equal(t, config.Expected, output, "bfs failed")
}

func TestBFS(t *testing.T) {
    //  0,0 -> 1,0
    //   |      |
    //  0,1 -> 1,1
    config := InputConfig[int]{
        Edges: map[Node][]Node {
            {0,0}: {{1,0}, {0, 1}},
            {0,1}: {{0,0}, {1, 1}},
            {1,0}: {{0,0}, {1, 1}},
            {1,1}: {{1,0}, {0, 1}},
        },
        Expected: 2,
        Start: Node{0,0},
    }
    config.RunBFS(t)
}
