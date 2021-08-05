package breadth_first

import "data-structures-and-algorithms/data-structures/trees/binary_search_tree"

func breadthFirstR(bst *binary_search_tree.BinarySearchTree, q []*binary_search_tree.TreeNode, result []int) []int {
	if len(q) == 0 {
		return result
	}

	currentNode := q[0]
	q = q[1:]

	if currentNode.Left != nil {
		q = append(q, currentNode.Left)
	}

	if currentNode.Right != nil {
		q = append(q, currentNode.Right)
	}

	return breadthFirstR(bst, q, result)
}
