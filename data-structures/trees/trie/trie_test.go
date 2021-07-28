package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := newTrie()

	t.Run("insert a word in trie", func(t *testing.T) {
		word := "alice"
		trie.insert(word)

		currentNode := trie.root
		for i := 0; i < len(word); i++ {
			idx := word[i] - 'a'
			if currentNode.children[idx] == nil {
				t.Errorf("trie node should not be nil but it is")
			}
			currentNode = currentNode.children[idx]
		}

		if currentNode.isEnd != true {
			t.Errorf("last node should indicate the end of the word")
		}
	})

	t.Run("insert another word in trie", func(t *testing.T) {
		word := "alison"
		trie.insert(word)

		currentNode := trie.root
		for i := 0; i < len(word); i++ {
			idx := word[i] - 'a'
			if currentNode.children[idx] == nil {
				t.Errorf("trie node should not be nil but it is")
			}
			currentNode = currentNode.children[idx]
		}

		if currentNode.isEnd != true {
			t.Errorf("last node should indicate the end of the word")
		}
	})

	t.Run("the two inserts should share the same path for three levels", func(t *testing.T) {
		currentNode := trie.root
		word := "alice"
		for n := 0; n < 3; n++ {
			for i := range currentNode.children {
				allocations := 0
				if currentNode.children[i] != nil {
					allocations++
				}

				if allocations > 1 {
					t.Errorf("unexpected memory allocation, the two inserts should share the same address")
				}
			}
			currentNode = currentNode.children[(word[n] - 'a')]
		}

		t.Run("the two inserts stop sharing the same path on the fourth level", func(t *testing.T) {
			fmt.Println(currentNode)
			allocations := 0

			for i := range currentNode.children {
				if currentNode.children[i] != nil {
					allocations++
				}
			}

			if allocations != 2 {
				t.Errorf("unexpected memory allocation, the two inserts should not share the same address")
			}
		})
	})

	t.Run("search for a valid word", func(t *testing.T) {
		word := "alison"
		if got := trie.search(word); got != true {
			t.Errorf("unexpected result on search, want: %t, but got %t", true, got)
		}
	})

	t.Run("search for invalid word", func(t *testing.T) {
		word := "alistair"
		if got := trie.search(word); got != false {
			t.Errorf("unexpected result on search, want: %t, but got %t", false, got)
		}
	})
}
