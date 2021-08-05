package pre_order

import "data-structures-and-algorithms/data-structures/trees/binary_search_tree"

type TreeNode *binary_search_tree.TreeNode

func preOrderTraversal(node TreeNode) []int {
	result := []int{}
	stack := []TreeNode{}

	for len(stack) > 0 || node != nil {
		if node != nil {
			result = append(result, node.Value)
			stack = append(stack, node)
			node = node.Left
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = last.Right
		}
	}

	return result
}

func preOrderTraversalR(node TreeNode, result []int) []int {
	if node == nil {
		return []int{}
	}

	result = append(result, node.Value)

	preOrderTraversal(node.Left)
	preOrderTraversal(node.Right)

	return result

}
