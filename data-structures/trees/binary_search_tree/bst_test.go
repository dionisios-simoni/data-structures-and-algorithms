package binary_search_tree

import "testing"

func TestBinarySearchTree(t *testing.T)  {

	bst := newBinarySearchTree()

	t.Run("inserting a root node", func(t *testing.T) {
		bst.insert(100)

		want := 100
		if got := bst.root.value; got != want {
			t.Errorf("invalid root node value, got: %d but want: %d: ", got, want)
		}
	})

	t.Run("inserting a right child", func(t *testing.T) {
		bst.insert(150)

		want := 150
		if got := bst.root.right.value; got != want {
			t.Errorf("invalid right child node value, got: %d but want: %d: ", got, want)
		}
	})

	t.Run("inserting a left child", func(t *testing.T) {
		bst.insert(50)

		want := 50
		if got := bst.root.left.value; got != want {
			t.Errorf("invalid right child node value, got: %d but want: %d: ", got, want)
		}
	})

	t.Run("forming a subtree", func(t *testing.T) {
		bst.insert(170)
		bst.insert(130)

		/*
					100
			50				 150
						130		  170
		*/

		wantRight := 170
		wantLeft := 130

		if got := bst.root.right.right.value; got != wantRight {
			t.Errorf("invalid right subtree child node value, got: %d but want: %d: ", got, wantRight)
		}

		if got := bst.root.right.left.value; got != wantLeft {
			t.Errorf("invalid left subtree child node value, got: %d but want: %d: ", got, wantLeft)
		}
	})

	t.Run("search for a node", func(t *testing.T) {
		got := bst.search(130)
		want := 130
		if got.value != want {
			t.Errorf("invalid node value was found, got: %d but want: %d: ", got.value, want)
		}
	})

	t.Run("search with invalid input", func(t *testing.T) {
		got := bst.search(007)
		if got != nil {
			t.Errorf("search should return nil, but got: %v", got)
		}

	})
}
