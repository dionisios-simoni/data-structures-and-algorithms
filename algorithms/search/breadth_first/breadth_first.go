package breadth_first

import (
	"data-structures-and-algorithms/data-structures/trees/binary_search_tree"
)

func breadthFirstSearch(bst *binary_search_tree.BinarySearchTree) []int {
	result := []int{}
	q := []*binary_search_tree.TreeNode{}
	currentNode := bst.Root

	q = append(q, currentNode)

	for len(q) > 0 {

		currentNode = q[0]
		q = q[1:]

		result = append(result, currentNode.Value)

		if currentNode.Left != nil {
			q = append(q, currentNode.Left)
		}

		if currentNode.Right != nil {
			q = append(q, currentNode.Right)
		}

	}

	return result
}
