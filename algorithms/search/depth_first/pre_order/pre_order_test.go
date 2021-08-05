package pre_order

import (
	"data-structures-and-algorithms/data-structures/trees/binary_search_tree"
	"testing"
)

func TestPreOrderTraversal(t *testing.T) {
	bst := binary_search_tree.BinarySearchTree{}
	bst.Insert(9)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(20)
	bst.Insert(170)
	bst.Insert(15)
	bst.Insert(1)

	want := []int{9,4,1,6,20,15,170}
	got := preOrderTraversal(bst.Root)

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("unexpected traversal result, want: %v got: %v", want, got)
		}
	}

	got = preOrderTraversalR(bst.Root, []int{})

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("unexpected traversal result, want: %v got: %v", want, got)
		}
	}
}
