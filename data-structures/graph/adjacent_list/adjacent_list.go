package adjacent_list

import (
	"data-structures-and-algorithms/data-structures/stack-queue/queue_array"
	"data-structures-and-algorithms/data-structures/stack-queue/stack_array"
	"errors"
	"fmt"
)

const infinity = int(^uint(0) >> 1)

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
// a connection may hold the connected vertex's value and
// a weight which may be associated to the connection
type edge struct {
	connections [][]int
}

// Table represents Dijkstra's algorithm result table
type Table struct {
	vertex   int
	visited  bool
	distance int
	previous int
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
			connections: [][]int{},
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

	g.appendEdge(v1.value, v2.value, weight)
	g.appendEdge(v2.value, v1.value, weight)

	return nil
}

func (g graph) appendEdge(v1, v2 int, weight *int) {
	conn := []int{}
	conn = append(conn, v2)

	if weight != nil {
		conn = append(conn, *weight)
	}

	g.adjacentList[v1].edge.connections = append(g.adjacentList[v1].edge.connections, conn)
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
			if visited[conn[0]] == false {
				queue.Enqueue(conn[0])
				visited[conn[0]] = true
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
			if visited[conn[0]] == false {
				visited[conn[0]] = true
				stack.Push(conn[0])
			}
		}
	}
	return result
}

func (g *graph) dijkstra(source int) []Table {
	visited := make(map[int]bool)
	distance := g.initDistance(source)
	previous := make(map[int]int)

	result := []Table{}

	for len(visited) < len(g.adjacentList) {
		next := getNext(distance, visited)

		visited[next] = true
		current := g.adjacentList[next]

		for _, conn := range current.edge.connections {

			neighbourDistance := conn[1] + distance[current.value]

			if distance[conn[0]] > neighbourDistance {
				distance[conn[0]] = neighbourDistance
				previous[conn[0]] = current.value
			}
		}

		result = append(result, Table{
			vertex:   current.value,
			previous: previous[current.value],
			distance: distance[current.value],
			visited:  visited[current.value],
		})
	}

	return result
}

func getNext(distance map[int]int, visited map[int]bool) int {
	smallest := infinity
	var next int

	for k, v := range distance {
		if v < smallest && !visited[k] {
			smallest = v
			next = k
		}
	}
	return next
}

func (g *graph) initDistance(source int) map[int]int {
	dst := make(map[int]int)
	for _, v := range g.adjacentList {
		if v.value == source {
			dst[v.value] = 0
		} else {
			dst[v.value] = infinity
		}
	}
	return dst
}
