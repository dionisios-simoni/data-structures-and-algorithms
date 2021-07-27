package binary_search_tree

type binarySearchTree struct {
	root *treeNode
}

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func newBinarySearchTree() *binarySearchTree {
	return &binarySearchTree{}
}

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

func (bst *binarySearchTree) search(value int) *treeNode {

	currentNode := bst.root

	for currentNode != nil {
		if value == currentNode.value {
			return currentNode
		}

		if value > currentNode.value {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}
	}

	return nil
}
