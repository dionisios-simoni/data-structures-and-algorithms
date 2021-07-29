package adjacent_list

import (
	"errors"
	"fmt"
)

// graph is a graph representation implemented using a map
type graph struct {
	graphSize    int
	adjacentList map[int]*vertex
}

// vertex represents each vertex in the graph
// holds a vertex value and an edge
type vertex struct {
	value int
	edge  *edge
}

// edge holds connections reference to all vertices
// which are connected with a given vertex
// the field weight may or may not be populated
// depending on the type of the graph.
type edge struct {
	connections []int
	weight      *int
}

func newGraph() *graph {
	return &graph{
		graphSize:    0,
		adjacentList: make(map[int]*vertex),
	}
}

// addVertex adds a new vertex if it does not already exist
func (g *graph) addVertex(value int) bool {
	newVertex := &vertex{
		value: value,
		edge: &edge{
			connections: []int{},
			weight:      nil,
		},
	}

	if _, ok := g.adjacentList[value]; !ok {
		g.adjacentList[value] = newVertex
		g.graphSize++
		return true
	}

	return false
}

//addEdge adds a connection between two vertices.
func (g *graph) addEdge(v1, v2 *vertex, weight *int) error {

	_, ok1 := g.adjacentList[v1.value]
	_, ok2 := g.adjacentList[v2.value]

	if !ok1 || !ok2 {
		return errors.New("invalid vertex provided, does not exist in the graph")
	}

	if weight != nil {
		v2.edge.weight = weight
	}

	g.appendEdge(v1.value, v2.value)
	g.appendEdge(v2.value, v1.value)

	return nil
}

func (g graph) appendEdge(v1, v2 int) {
	g.adjacentList[v1].edge.connections = append(g.adjacentList[v1].edge.connections, v2)
}

// showConnections produces a user friendly output demonstrating the graph connections.
func (g *graph) showConnections() {
	for _, v := range g.adjacentList {
		fmt.Println("vertex", v.value, "edges connect with vertices:", v.edge.connections)
	}
}
