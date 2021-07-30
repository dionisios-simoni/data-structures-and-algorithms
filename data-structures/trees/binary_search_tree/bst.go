package binary_search_tree

// binarySearchTree is a data structure that holds nodes
type binarySearchTree struct {
	root *treeNode
}

// treeNode represents an entry on the binary search tree
type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func newBinarySearchTree() *binarySearchTree {
	return &binarySearchTree{}
}

// insert adds a new node to the tree, creating multiple levels of hierarchy
func (bst *binarySearchTree) insert(value int) {
	node := &treeNode{value: value}

	if bst.root == nil {
		bst.root = node
		return
	}

	currentNode := bst.root
	for {
		if value > currentNode.value {
			if currentNode.right == nil {
				currentNode.right = node
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = node
				return
			}
			currentNode = currentNode.left
		}
	}
}

// search returns a node if found in the tree, given a value
func (bst *binarySearchTree) search(value int) (node, parent *treeNode) {

	currentNode := bst.root
	parentNode := &treeNode{}

	for currentNode != nil {
		if value == currentNode.value {
			return currentNode, parentNode
		}

		if value > currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.right
		} else {
			parentNode = currentNode
			currentNode = currentNode.left
		}
	}

	return nil, nil
}

func (bst *binarySearchTree) remove(value int) {
	node, parent := bst.search(value)

	if isLeaf(node) {
		if node.value > parent.value {
			parent.right = nil
		} else {
			parent.left = nil
		}
	}

	if onlyLeft(node) {
		if parent.value > node.value {
			parent.left = node.left
		} else {
			parent.right = node.left
		}
	}

	if onlyRight(node) {
		if parent.value > node.value {
			parent.left = node.right
		} else {
			parent.right = node.right
		}
	}

	if isParent(node) {
		predecessor := node.right
		leaf := &treeNode{}

		for predecessor.left != nil {
			leaf = predecessor
			predecessor = predecessor.left
		}

		if isRoot(parent) {
			bst.root.value = predecessor.value
		} else {
			parent.right = predecessor
			predecessor.left = node.left
		}

		leaf.left = nil
	}
}

func isRoot(node *treeNode) bool {
	return node.left == nil && node.right == nil
}

func isLeaf(node *treeNode) bool {
	return node.left == nil && node.right == nil
}

func isParent(node *treeNode) bool {
	return node.left != nil && node.right != nil
}

func onlyLeft(node *treeNode) bool {
	return node.left != nil && node.right == nil
}

func onlyRight(node *treeNode) bool {
	return node.right != nil && node.left == nil
}
