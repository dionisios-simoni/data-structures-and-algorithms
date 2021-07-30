package binary_search_tree

import (
	"testing"
)

func TestInsert(t *testing.T) {

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
}

func TestSearch(t *testing.T) {

	bst := newBinarySearchTree()
	bst.insert(100)
	bst.insert(150)
	bst.insert(50)
	bst.insert(170)
	bst.insert(130)

	t.Run("search for a node", func(t *testing.T) {
		got, _ := bst.search(130)
		want := 130

		if got == nil {
			t.Errorf("unexpected nil return, should return node")
		} else {
			if got.value != want {
				t.Errorf("invalid node value was found, got: %d but want: %d: ", got.value, want)
			}
		}
	})

	t.Run("search with invalid input", func(t *testing.T) {
		got, _ := bst.search(007)
		if got != nil {
			t.Errorf("search should return nil, but got: %v", got)
		}
	})
}

func TestRemove(t *testing.T) {

	bst := newBinarySearchTree()

	bst.insert(8)
	bst.insert(4)
	bst.insert(12)
	bst.insert(2)
	bst.insert(6)
	bst.insert(10)
	bst.insert(14)
	bst.insert(1)
	bst.insert(3)
	bst.insert(5)
	bst.insert(7)
	bst.insert(9)
	bst.insert(11)
	bst.insert(13)

	t.Run("remove leaf node", func(t *testing.T) {
		bst.remove(1)
		node := bst.root
		want := 2

		for node.left != nil {
			node = node.left
		}

		if node.value != want {
			t.Errorf("parent to leaf node still points to a value but shoulnd't")
		}
	})

	t.Run("remove a node with one child (left)", func(t *testing.T) {
		bst.remove(14)
		node := bst.root
		want := 13

		for node.right != nil {
			node = node.right
		}

		if got := node.value; got != want {
			t.Errorf("invalid node replacement, should replace with: %d but value is: %d", want, got)
		}
	})

	t.Run("remove a node with one child (right)", func(t *testing.T) {
		bst.remove(2)
		node := bst.root
		want := 3

		for node.left != nil {
			node = node.left
		}

		if got := node.value; got != want {
			t.Errorf("invalid node replacement, should replace with: %d but value is: %d", want, got)
		}
	})

	t.Run("remove parent with two children", func(t *testing.T) {
		bst.remove(6)
		want := 7
		root := bst.root
		got := root.left.right.value

		if got != want {
			t.Errorf("invalid node replacement, should replace with: %d but value is: %d", want, got)
		}
	})

	t.Run("remove root node", func(t *testing.T) {
		want := 9
		bst.remove(8)

		if got := bst.root.value; got != want {
			t.Errorf("invalid")
		}

		if bst.root.right.value != 12 || bst.root.left.value != 4 {
			t.Errorf("root removal does not preserve correct element order")
		}
	})

	t.Run("check on tree order after removals", func(t *testing.T) {
		elements := []int{9, 4, 12, 3, 7, 10, 13, 5, 11}
		want := bst.root.value

		if want != elements[0] {
			t.Errorf("unexpected root element, want:%d but got:%d", elements[0], want)
		}

		want = bst.root.left.value
		if want != elements[1] {
			t.Errorf("unexpected root element, want:%d but got:%d", elements[1], want)
		}

		want = bst.root.right.value
		if want != elements[2] {
			t.Errorf("unexpected root element, want:%d but got:%d", elements[2], want)
		}

		want = bst.root.left.left.value
		if want != elements[3] {
			t.Errorf("unexpected root element, want:%d but got:%d", elements[3], want)
		}
		want = bst.root.left.right.value
		if want != elements[4] {
			t.Errorf("unexpected root element, want:%d but got:%d", elements[4], want)
		}

		want = bst.root.right.left.value
		if want != elements[5] {
			t.Errorf("unexpected root element, want:%d but got:%d", elements[5], want)
		}

		want = bst.root.right.right.value
		if want != elements[6] {
			t.Errorf("unexpected root element, want:%d but got:%d", elements[6], want)
		}

		want = bst.root.left.right.left.value
		if want != elements[7] {
			t.Errorf("unexpected root element, want:%d but got:%d", elements[7], want)
		}

		want = bst.root.right.left.right.value
		if want != elements[8] {
			t.Errorf("unexpected root element, want:%d but got:%d", elements[8], want)
		}
	})
}
