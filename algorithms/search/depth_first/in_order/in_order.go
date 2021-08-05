package in_order

import (
	"data-structures-and-algorithms/data-structures/trees/binary_search_tree"
)

type TreeNode *binary_search_tree.TreeNode

func inOrderTraversal(node TreeNode) []int {
	result := []int{}
	stack := []TreeNode{}

	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, last.Value)
			node = last.Right
		}
	}
	return result
}

func inOrderTraversalR(node TreeNode) []int {
	if node == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, inOrderTraversalR(node.Left)...)
	result = append(result, node.Value)
	result = append(result, inOrderTraversalR(node.Right)...)

	return result
}
