package adjacent_list

import (
	"sort"
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
			if want[i] != got[i][0] {
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
	for i := 1; i <= 5; i++ {
		wg.addVertex(i)
	}

	t.Run("add edge weight in graph connections", func(t *testing.T) {
		v1 := wg.adjacentList[1] // = a
		v2 := wg.adjacentList[2] // = b
		v3 := wg.adjacentList[3] // = c
		v4 := wg.adjacentList[4] // = d
		v5 := wg.adjacentList[5] // = e

		/*
			undirected, weighted, cyclic graph
				(a) - - 6 - - (b)
				 |		    _ /	  \5
				 1	  _ 2 /	    	 (c)
				 |	/			  /5
				(d) - - 1 - - (e)
		*/

		weight := 6
		err := wg.addEdge(v1, v2, &weight)
		assertNoError(t, err)

		weight = 1
		err = wg.addEdge(v1, v4, &weight)
		assertNoError(t, err)

		weight = 2
		err = wg.addEdge(v4, v2, &weight)
		assertNoError(t, err)

		weight = 1
		err = wg.addEdge(v4, v5, &weight)
		assertNoError(t, err)

		weight = 2
		err = wg.addEdge(v2, v5, &weight)
		assertNoError(t, err)

		weight = 5
		err = wg.addEdge(v2, v3, &weight)
		assertNoError(t, err)

		weight = 5
		err = wg.addEdge(v5, v3, &weight)
		assertNoError(t, err)

		t.Run("weight total is correct", func(t *testing.T) {
			want := 22
			got := 0
			for _, v := range wg.adjacentList {
				for _, conn := range v.edge.connections {
					got += conn[1]
				}
			}

			if want != got/2 {
				t.Errorf("weighted graph total is invalid, expected: %d but got %d", want, got)
			}
		})

		t.Run("djikstra's algorithm - shortest path from start to anywhere", func(t *testing.T) {
			got := wg.dijkstra(1) // starting point is a;

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].vertex < got[j].vertex
			})

			// 1 = a, 2 = b, 3 = c, 4 = d, 5 = e
			want := []Table{
				{1, true, 0, 0},
				{2, true, 3, 4},
				{3, true, 7, 5},
				{4, true, 1, 1},
				{5, true, 2, 4},
			}

			t.Run("correct shortest path for each vertice is formed", func(t *testing.T) {
				wantString := "dead"
				gotString := ""

				for i := 1; i < len(want); i++ {
					asciiCode := 96 + want[i].previous
					gotString += string(rune(asciiCode))
				}

				if wantString != gotString {
					t.Errorf("invalid shortest path was formed for vertex")
				}
			})

			for i := range got {
				if got[i].vertex != want[i].vertex {
					t.Errorf("vertex value: %d does not match expected value: %d", got[i].vertex, want[i].vertex)
				}

				if got[i].visited != want[i].visited {
					t.Errorf("vertex value: %t does not match expected value: %t", got[i].visited, want[i].visited)
				}

				if got[i].distance != want[i].distance {
					t.Errorf("vertex value: %d does not match expected value: %d", got[i].distance, want[i].distance)
				}

				if got[i].previous != want[i].previous {
					t.Errorf("vertex value: %d does not match expected value: %d", got[i].previous, want[i].previous)
				}
			}
		})
	})
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error occured while adding edges %v", err)
	}
}
