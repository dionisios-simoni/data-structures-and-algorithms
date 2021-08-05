package breadth_first

import (
	"data-structures-and-algorithms/data-structures/trees/binary_search_tree"
	"testing"
)

func TestBFS(t *testing.T) {
	var bst = &binary_search_tree.BinarySearchTree{}

	bst.Insert(9)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(20)
	bst.Insert(170)
	bst.Insert(15)
	bst.Insert(1)

	//	9
	//	4 - 20
	//  1-6  15-170

	got := breadthFirstSearch(bst)
	want := []int{9, 4, 20, 1, 6, 15, 170}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("unexpected result: want %v got %v", want, got)
		}
	}

	got = breadthFirstR(bst, []*binary_search_tree.TreeNode{}, []int{})
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("unexpected result: want %v got %v", want, got)
		}
	}
}
