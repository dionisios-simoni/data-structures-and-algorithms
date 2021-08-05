package binary_search_tree

import (
	"testing"
)

func TestInsert(t *testing.T) {

	bst := newBinarySearchTree()

	t.Run("inserting a Root node", func(t *testing.T) {
		bst.Insert(100)

		want := 100
		if got := bst.Root.Value; got != want {
			t.Errorf("invalid Root node Value, got: %d but want: %d: ", got, want)
		}
	})

	t.Run("inserting a Right child", func(t *testing.T) {
		bst.Insert(150)

		want := 150
		if got := bst.Root.Right.Value; got != want {
			t.Errorf("invalid Right child node Value, got: %d but want: %d: ", got, want)
		}
	})

	t.Run("inserting a Left child", func(t *testing.T) {
		bst.Insert(50)

		want := 50
		if got := bst.Root.Left.Value; got != want {
			t.Errorf("invalid Right child node Value, got: %d but want: %d: ", got, want)
		}
	})

	t.Run("forming a subtree", func(t *testing.T) {
		bst.Insert(170)
		bst.Insert(130)

		/*
					100
			50				 150
						130		  170
		*/

		wantRight := 170
		wantLeft := 130

		if got := bst.Root.Right.Right.Value; got != wantRight {
			t.Errorf("invalid Right subtree child node Value, got: %d but want: %d: ", got, wantRight)
		}

		if got := bst.Root.Right.Left.Value; got != wantLeft {
			t.Errorf("invalid Left subtree child node Value, got: %d but want: %d: ", got, wantLeft)
		}
	})
}

func TestSearch(t *testing.T) {

	bst := newBinarySearchTree()
	bst.Insert(100)
	bst.Insert(150)
	bst.Insert(50)
	bst.Insert(170)
	bst.Insert(130)

	t.Run("Search for a node", func(t *testing.T) {
		got, _ := bst.Search(130)
		want := 130

		if got == nil {
			t.Errorf("unexpected nil return, should return node")
		} else {
			if got.Value != want {
				t.Errorf("invalid node Value was found, got: %d but want: %d: ", got.Value, want)
			}
		}
	})

	t.Run("Search with invalid input", func(t *testing.T) {
		got, _ := bst.Search(007)
		if got != nil {
			t.Errorf("Search should return nil, but got: %v", got)
		}
	})
}

func TestRemove(t *testing.T) {

	bst := newBinarySearchTree()

	bst.Insert(8)
	bst.Insert(4)
	bst.Insert(12)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(10)
	bst.Insert(14)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(7)
	bst.Insert(9)
	bst.Insert(11)
	bst.Insert(13)

	t.Run("Remove leaf node", func(t *testing.T) {
		bst.Remove(1)
		node := bst.Root
		want := 2

		for node.Left != nil {
			node = node.Left
		}

		if node.Value != want {
			t.Errorf("parent to leaf node still points to a Value but shoulnd't")
		}
	})

	t.Run("Remove a node with one child (Left)", func(t *testing.T) {
		bst.Remove(14)
		node := bst.Root
		want := 13

		for node.Right != nil {
			node = node.Right
		}

		if got := node.Value; got != want {
			t.Errorf("invalid node replacement, should replace with: %d but Value is: %d", want, got)
		}
	})

	t.Run("Remove a node with one child (Right)", func(t *testing.T) {
		bst.Remove(2)
		node := bst.Root
		want := 3

		for node.Left != nil {
			node = node.Left
		}

		if got := node.Value; got != want {
			t.Errorf("invalid node replacement, should replace with: %d but Value is: %d", want, got)
		}
	})

	t.Run("Remove parent with two children", func(t *testing.T) {
		bst.Remove(6)
		want := 7
		root := bst.Root
		got := root.Left.Right.Value

		if got != want {
			t.Errorf("invalid node replacement, should replace with: %d but Value is: %d", want, got)
		}
	})

	t.Run("Remove Root node", func(t *testing.T) {
		want := 9
		bst.Remove(8)

		if got := bst.Root.Value; got != want {
			t.Errorf("invalid")
		}

		if bst.Root.Right.Value != 12 || bst.Root.Left.Value != 4 {
			t.Errorf("Root removal does not preserve correct element order")
		}
	})

	t.Run("check on tree order after removals", func(t *testing.T) {
		elements := []int{9, 4, 12, 3, 7, 10, 13, 5, 11}
		want := bst.Root.Value

		if want != elements[0] {
			t.Errorf("unexpected Root element, want:%d but got:%d", elements[0], want)
		}

		want = bst.Root.Left.Value
		if want != elements[1] {
			t.Errorf("unexpected Root element, want:%d but got:%d", elements[1], want)
		}

		want = bst.Root.Right.Value
		if want != elements[2] {
			t.Errorf("unexpected Root element, want:%d but got:%d", elements[2], want)
		}

		want = bst.Root.Left.Left.Value
		if want != elements[3] {
			t.Errorf("unexpected Root element, want:%d but got:%d", elements[3], want)
		}
		want = bst.Root.Left.Right.Value
		if want != elements[4] {
			t.Errorf("unexpected Root element, want:%d but got:%d", elements[4], want)
		}

		want = bst.Root.Right.Left.Value
		if want != elements[5] {
			t.Errorf("unexpected Root element, want:%d but got:%d", elements[5], want)
		}

		want = bst.Root.Right.Right.Value
		if want != elements[6] {
			t.Errorf("unexpected Root element, want:%d but got:%d", elements[6], want)
		}

		want = bst.Root.Left.Right.Left.Value
		if want != elements[7] {
			t.Errorf("unexpected Root element, want:%d but got:%d", elements[7], want)
		}

		want = bst.Root.Right.Left.Right.Value
		if want != elements[8] {
			t.Errorf("unexpected Root element, want:%d but got:%d", elements[8], want)
		}
	})
}

func TestTraverse(t *testing.T) {

	bst := newBinarySearchTree()

	bst.Insert(9)
	bst.Insert(4)
	bst.Insert(20)
	bst.Insert(1)
	bst.Insert(6)
	bst.Insert(15)
	bst.Insert(170)

	got := bst.Root.traverse()
	want := []int{9, 4, 1, 6, 20, 15, 170}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("invalid element order in traversal expected: %v but got: %v", want, got)
		}
	}
}
