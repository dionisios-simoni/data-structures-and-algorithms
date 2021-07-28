package trie

const size = 26

// node represents each node in a trie
type node struct {
	isEnd    bool
	children [size]*node
}

// trie represents a trie data structure
type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{root: &node{}}
}

// insert creates a new node if one does not exist already
// to represent a lowercase word
func (t *trie) insert(w string) {
	currentNode := t.root

	for i := 0; i < len(w); i++ {
		idx := w[i] - 'a'
		if currentNode.children[idx] == nil {
			currentNode.children[idx] = &node{}
		}
		currentNode = currentNode.children[idx]
	}
	currentNode.isEnd = true
}

// search receives a lowercase word to search for and returns true if found
func (t *trie) search(w string) bool {
	currentNode := t.root

	for i := 0; i < len(w); i++ {
		idx := w[i] - 'a'
		if currentNode.children[idx] == nil {
			return false
		}
		currentNode = currentNode.children[idx]
	}

	return currentNode.isEnd
}
