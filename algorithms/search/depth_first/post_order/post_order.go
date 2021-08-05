package post_order

import (
	"data-structures-and-algorithms/data-structures/trees/binary_search_tree"
)

type TreeNode *binary_search_tree.TreeNode

func postOrderTraversal(node TreeNode) []int {

	result := []int{}
	stack := []TreeNode{node}

	for len(stack) > 0 {

		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.Value)

		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}

	// reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func postOrderTraversalR(node TreeNode) []int {

	if node == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, postOrderTraversalR(node.Left)...)
	result = append(result, postOrderTraversalR(node.Right)...)
	result = append(result, node.Value)

	return result

}
