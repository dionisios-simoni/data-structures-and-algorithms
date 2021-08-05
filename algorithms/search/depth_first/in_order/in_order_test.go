package in_order

import (
	"data-structures-and-algorithms/data-structures/trees/binary_search_tree"
	"testing"
)

func TestInOrderTraversal(t *testing.T) {
	bst := binary_search_tree.BinarySearchTree{}
	bst.Insert(9)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(20)
	bst.Insert(170)
	bst.Insert(15)
	bst.Insert(1)

	want := []int{1, 4, 6, 9, 15, 20, 170}
	got := inOrderTraversal(bst.Root)

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("unexpected traversal result, want: %v got: %v", want, got)
		}
	}

	got = inOrderTraversalR(bst.Root)
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("unexpected traversal result, want: %v got: %v", want, got)
		}
	}
}
