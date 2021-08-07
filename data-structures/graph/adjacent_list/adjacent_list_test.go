package adjacent_list

import (
	"testing"
)

func TestUnweightedGraph(t *testing.T) {
	g := newGraph()

	t.Run("insert vertices in the graph", func(t *testing.T) {
		for i := 10; i <= 50; i += 10 {
			g.addVertex(i)
		}

		want := 5
		if got := g.graphSize; got != want {
			t.Errorf("unexpected graph size, got: %d but expected %d", got, want)
		}
	})

	t.Run("vertices currently have no connections", func(t *testing.T) {
		for k, v := range g.adjacentList {
			if len(v.edge.connections) > 0 {
				t.Errorf("unexpected vertice edge, vertex %d should not have any connections", k)
			}
		}
	})

	/*

		undirected, unweighted, cyclic graph

		    (20)	(40)
		   /	\	/
		(10)	(30)
			\	/
			 (50)
	*/

	t.Run("add edges to form graph connections", func(t *testing.T) {
		v1 := g.adjacentList[10]
		v2 := g.adjacentList[20]
		v3 := g.adjacentList[30]
		v4 := g.adjacentList[40]
		v5 := g.adjacentList[50]

		err := g.addEdge(v1, v2, nil)
		assertNoError(t, err)
		err = g.addEdge(v1, v5, nil)
		assertNoError(t, err)

		err = g.addEdge(v2, v3, nil)
		assertNoError(t, err)

		err = g.addEdge(v3, v4, nil)
		assertNoError(t, err)
		err = g.addEdge(v3, v5, nil)
		assertNoError(t, err)

		want := []int{20, 40, 50}
		got := g.adjacentList[30].edge.connections

		for i := 0; i < len(got); i++ {
			if want[i] != got[i] {
				t.Errorf("invalid vertice edge connection expected %v but got %v", want, got)
			}
		}
	})

	t.Run("breadth first search", func(t *testing.T) {
		got := g.traverseBreadthFirstSearch(10)
		want := []int{10, 20, 50, 30, 40}

		for i := range want {
			if got[i] != want[i] {
				t.Errorf("invalid traversal result, expected: %v but got %v", want, got)
			}
		}
	})

	t.Run("depth first search", func(t *testing.T) {
		got := g.traverseDepthFirstSearch(10)
		want := []int{10, 50, 30, 40, 20}

		for i := range want {
			if got[i] != want[i] {
				t.Errorf("invalid traversal result, expected: %v but got %v", want, got)
			}
		}
	})

	g.showConnections()
}

func TestWeightedGraph(t *testing.T) {

	wg := newGraph()
	for i := 10; i <= 50; i += 10 {
		wg.addVertex(i)
	}

	t.Run("add edge weight in graph connections", func(t *testing.T) {
		v1 := wg.adjacentList[10]
		v2 := wg.adjacentList[20]
		v3 := wg.adjacentList[30]

		weight := 10

		err := wg.addEdge(v1, v2, &weight)
		assertNoError(t, err)
		err = wg.addEdge(v1, v3, &weight)
		assertNoError(t, err)

		want := 20
		got := *v2.edge.weight + *v3.edge.weight
		if got != want {
			t.Errorf("unexpected edge weight calculated, want %d but got %d", want, got)
		}
	})
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error occured while adding edges %v", err)
	}
}
