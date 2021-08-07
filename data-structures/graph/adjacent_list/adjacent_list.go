package adjacent_list

import (
	"data-structures-and-algorithms/data-structures/stack-queue/queue_array"
	"data-structures-and-algorithms/data-structures/stack-queue/stack_array"
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

func (g *graph) traverseBreadthFirstSearch(source int) []int {
	result := []int{}
	visited := make(map[int]bool)
	queue := queue_array.Queue{}

	if _, ok := g.adjacentList[source]; !ok {
		return result
	}


	queue.Enqueue(source)
	visited[source] = true

	for queue.Len > 0 {
		current := queue.Dequeue()
		vtx := g.adjacentList[current]

		result = append(result, vtx.value)

		for _, conn := range vtx.edge.connections {
			if visited[conn] == false {
				queue.Enqueue(conn)
				visited[conn] = true
			}
		}
	}

	return result
}

func (g *graph) traverseDepthFirstSearch(source int) []int {
	result := []int{}
	visited := make(map[int]bool)
	stack := stack_array.Stack{}

	if _, ok := g.adjacentList[source]; !ok {
		return result
	}

	stack.Push(source)
	visited[source] = true

	for stack.Len > 0 {
		current := stack.Pop()
		vtx := g.adjacentList[current]

		result = append(result, current)

		for _, conn := range vtx.edge.connections {
			if visited[conn] == false {
				visited[conn] = true
				stack.Push(conn)
			}
		}
	}
	return result
}