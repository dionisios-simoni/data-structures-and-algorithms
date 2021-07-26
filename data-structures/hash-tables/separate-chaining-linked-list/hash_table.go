package separate_chaining_linked_list

import (
	"errors"
)

// size is a constant variable which defines the size of the
// hashTable bucket array.
const size = 7

// hashTable is a data structure which holds an array of linked lists.
type hashTable struct {
	bucket [size]*linkedList
}

// linkedList is a data structure which holds nodes.
type linkedList struct {
	head   *node
	length int
}

// node defines the elements of a linked list.
type node struct {
	key  string
	next *node
}

// init func which constructs a hashTable
func newHashTable() *hashTable {
	ht := &hashTable{}
	for i := 0; i < size; i++ {
		ht.bucket[i] = &linkedList{}
	}
	return ht
}

// Insert adds the key in the bucket if it doesn't already exist.
func (ht *hashTable) Insert(key string) error {
	index := hashFunc(key)
	err := ht.bucket[index].insert(key)
	if err != nil {
		return err
	}
	return nil
}

func (ll *linkedList) insert(key string) error {
	if !ll.search(key) {
		newNode := &node{key: key}
		newNode.next = ll.head
		ll.head = newNode
		ll.length++
	} else {
		return errors.New("the key already exists")
	}
	return nil
}

// Search returns true or false depending on whether the key provided
// was found in the linked list.
func (ht *hashTable) Search(key string) bool {
	index := hashFunc(key)
	// index is at which bucket this is stored.
	bucket := ht.bucket[index]
	return bucket.search(key)
}

func (ll *linkedList) search(key string) bool {
	currentNode := ll.head

	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

// Delete removes a node from the link list if exists.
func (ht *hashTable) Delete(key string) {
	index := hashFunc(key)
	ht.bucket[index].delete(key)
}

func (ll *linkedList) delete(key string) {
	previous := ll.head

	if previous.key == key {
		ll.head = ll.head.next
		ll.length--
		return
	}

	for previous.next != nil {
		if previous.next.key == key {
			previous.next = previous.next.next
			ll.length--
			return
		}
		previous = previous.next
	}
}

// hashFunc determines based on the provided key
// which bucket index will be allocated - causes hashing collisions.
func hashFunc(key string) int {
	sum := 0
	for _, r := range key {
		sum += int(r)
	}
	return sum % size
}
