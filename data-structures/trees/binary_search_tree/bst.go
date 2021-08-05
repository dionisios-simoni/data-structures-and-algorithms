package binary_search_tree

// BinarySearchTree is a data structure that holds nodes
type BinarySearchTree struct {
	Root *TreeNode
}

// TreeNode represents an entry on the binary Search tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func newBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Insert adds a new node to the tree, creating multiple levels of hierarchy
func (bst *BinarySearchTree) Insert(value int) {
	node := &TreeNode{Value: value}

	if bst.Root == nil {
		bst.Root = node
		return
	}

	currentNode := bst.Root
	for {
		if value > currentNode.Value {
			if currentNode.Right == nil {
				currentNode.Right = node
				return
			}
			currentNode = currentNode.Right
		} else {
			if currentNode.Left == nil {
				currentNode.Left = node
				return
			}
			currentNode = currentNode.Left
		}
	}
}

// Search returns a node if found in the tree, given a Value
func (bst *BinarySearchTree) Search(value int) (node, parent *TreeNode) {

	currentNode := bst.Root
	parentNode := &TreeNode{}

	for currentNode != nil {
		if value == currentNode.Value {
			return currentNode, parentNode
		}

		if value > currentNode.Value {
			parentNode = currentNode
			currentNode = currentNode.Right
		} else {
			parentNode = currentNode
			currentNode = currentNode.Left
		}
	}

	return nil, nil
}

func (bst *BinarySearchTree) Remove(value int) {
	node, parent := bst.Search(value)

	if isLeaf(node) {
		if node.Value > parent.Value {
			parent.Right = nil
		} else {
			parent.Left = nil
		}
	}

	if onlyLeft(node) {
		if parent.Value > node.Value {
			parent.Left = node.Left
		} else {
			parent.Right = node.Left
		}
	}

	if onlyRight(node) {
		if parent.Value > node.Value {
			parent.Left = node.Right
		} else {
			parent.Right = node.Right
		}
	}

	if isParent(node) {
		predecessor := node.Right
		leaf := &TreeNode{}

		for predecessor.Left != nil {
			leaf = predecessor
			predecessor = predecessor.Left
		}

		if isRoot(parent) {
			bst.Root.Value = predecessor.Value
		} else {
			parent.Right = predecessor
			predecessor.Left = node.Left
		}

		leaf.Left = nil
	}
}

func (tn *TreeNode) traverse() []int {

	var preOrder []int

	if tn != nil {
		preOrder = append(preOrder, tn.Value)
		tn.Left.traverse()
		tn.Right.traverse()
	}
	return preOrder
}

func isRoot(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func isParent(node *TreeNode) bool {
	return node.Left != nil && node.Right != nil
}

func onlyLeft(node *TreeNode) bool {
	return node.Left != nil && node.Right == nil
}

func onlyRight(node *TreeNode) bool {
	return node.Right != nil && node.Left == nil
}
