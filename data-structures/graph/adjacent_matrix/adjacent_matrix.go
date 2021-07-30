package adjacent_matrix

import (
	"fmt"
	"log"
)

// graph implementation using a matrix
type graph struct {
	matrix [][]int
}

func newGraph() *graph {
	mtx := make([][]int, 5)
	for i := range mtx {
		mtx[i] = make([]int, 5)
	}

	return &graph{matrix: mtx}
}

// addVertex adds a new vertex and its connections in the matrix
func (g *graph) addVertex() {

	for i := range g.matrix {
		g.matrix[i] = append(g.matrix[i], 0)
	}

	connections := make([]int, len(g.matrix)+1)
	g.matrix = append(g.matrix, connections)
}

func (g *graph) removeVertex(idx int) []int {
	removed := g.matrix[idx]
	g2 := g.matrix[:idx]
	g2 = append(g2, g.matrix[idx+1:]...)
	g.matrix = g2
	return removed
}

func (g *graph) addEdge(idx1, idx2 int) {

	if idx1 > len(g.matrix) || idx2 > len(g.matrix) {
		log.Fatalf("invalid vertex provided, should be less than %d", len(g.matrix))
	}

	g.matrix[idx1][idx2] = 1
	g.matrix[idx2][idx1] = 1
}

func (g *graph) showConnections() {
	for i := 0; i < len(g.matrix); i++ {
		fmt.Printf("vertex (%d) connections: %v\n", i, g.matrix[i])
	}
}
