package adjacent_matrix

import (
	"testing"
)

func TestAdjacentMatrix(t *testing.T) {
	g := newGraph()

	t.Run("populate matrix with connections between vertices", func(t *testing.T) {

		/*
			         1
				  /     \
				0	     3	-  4
				  \		/
					 2
		*/

		g.addEdge(0, 1)
		assertConnection(t, g.matrix[0], []int{0, 1, 0, 0, 0})

		g.addEdge(0, 2)
		assertConnection(t, g.matrix[0], []int{0, 1, 1, 0, 0})

		g.addEdge(1, 3)
		assertConnection(t, g.matrix[1], []int{1, 0, 0, 1, 0})

		g.addEdge(2, 3)
		assertConnection(t, g.matrix[2], []int{1, 0, 0, 1, 0})

		g.addEdge(3, 4)
		assertConnection(t, g.matrix[3], []int{0, 1, 1, 0, 1})

		assertConnection(t, g.matrix[4], []int{0, 0, 0, 1, 0})
	})

	t.Run("add new vertex in matrix", func(t *testing.T) {

		oldLength := len(g.matrix)

		t.Run("length of matrix should increase", func(t *testing.T) {

			want := oldLength + 1
			g.addVertex()

			if len(g.matrix) <= oldLength {
				t.Errorf("matrix length should increase but did not, want: %d, got: %d", want, oldLength)
			}
		})

		t.Run("connections length should increase", func(t *testing.T) {
			for i := range g.matrix {
				got := len(g.matrix[i])
				if got <= oldLength {
					t.Errorf("connections length should increas but did not, want %d, got: %d", oldLength+1, got)
				}
			}
		})

		t.Run("connections last element defaults to 0", func(t *testing.T) {
			for i := 0; i < len(g.matrix); i++ {
				if g.matrix[i][len(g.matrix)-1] != 0 {
					t.Errorf("invalid connection was formed, the new vertix connection should default to 0")
				}
			}
		})

		t.Run("remove a vertex in matrix", func(t *testing.T) {

			oldLength = len(g.matrix)
			want := oldLength - 1
			removed := g.removeVertex(5)

			t.Run("correct vertex is removed", func(t *testing.T) {
				expected := []int{0, 0, 0, 0, 0, 0}
				for i := range removed {
					if removed[i] != expected[i] {
						t.Errorf("unexpexted removal, expected to remove: %v, removed: %v", expected, removed)
					}
				}
			})

			t.Run("length of matrix should decrease", func(t *testing.T) {
				if len(g.matrix) >= oldLength {
					t.Errorf("matrix length should decrease but did not, want: %d, got: %d", want, oldLength)
				}
			})

			g.showConnections()
		})
	})
}

func assertConnection(t *testing.T, got, want []int) {
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("unexpected graph connection formed. want: %v got: %v", want, got)
		}
	}

}
