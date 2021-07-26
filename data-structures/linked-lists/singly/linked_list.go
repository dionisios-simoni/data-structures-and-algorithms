package singly

import (
	"errors"
	"fmt"
)

// linkedList represents a singly linked list
type linkedList struct {
	head, tail *node
	length     int
}

// node represents the data stored in the linked list
type node struct {
	data interface{}
	next *node
}

// isEmpty returns true if the linked list has no elements.
func (ll *linkedList) isEmpty() bool {
	return ll.length == 0
}

// prepend adds a new node at the start of a linked list which becomes the head.
func (ll *linkedList) prepend(v interface{}) {
	n := &node{data: v}

	if ll.isEmpty() {
		ll.head = n
		ll.tail = n
	} else {
		n.next = ll.head
		ll.head = n
	}
	ll.length++
}

// append adds a new node at the end of a linked list which becomes the tail.
func (ll *linkedList) append(v interface{}) {
	n := &node{data: v}

	if ll.isEmpty() {
		ll.head = n
		ll.tail = n
	} else {
		ll.tail.next = n
		ll.tail = n
	}
	ll.length++
}

// insert adds a new node in the linked list at a given position
func (ll *linkedList) insert(v interface{}, position int) error {
	n := &node{data: v}

	switch {
	// Head
	case position == 0:
		ll.prepend(n.data)
		return nil
	// Tail
	case position == ll.length:
		ll.append(n.data)
		return nil
	case position > 0 && position < ll.length:
		currentNode := ll.head.next
		for i := 1; i < ll.length; i++ {
			if i == position-1 {
				n.next = currentNode.next
				currentNode.next = n
				ll.length++
			} else {
				currentNode = currentNode.next
			}
		}
	default:
		return errors.New("position out of bounds")
	}
	return nil
}

// search returns the node if found in the linked list or an error
func (ll *linkedList) search(v interface{}) (*node, error) {

	if ll.isEmpty() {
		return nil, errors.New("linkedList is empty")
	}

	currentNode := ll.head
	for i := 0; i <= ll.length; i++ {
		if currentNode.data == v {
			return currentNode, nil
		}
		currentNode = currentNode.next
	}
	return nil, fmt.Errorf("element %v does not exist in linkedList", v)
}

// delete removes a node from the linked list
func (ll *linkedList) delete(v interface{}) error {

	if ll.isEmpty() {
		return errors.New("the linked list is empty")
	}

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		ll.length = 0
		return nil
	}

	if ll.head.data == v {
		ll.head = ll.head.next
		ll.length--
		return nil
	}

	currentNode := ll.head
	for i := 1; i <= ll.length; i++ {
		if currentNode.next.data != v {
			currentNode = currentNode.next
		} else {
			if currentNode.next == ll.tail {
				ll.tail = currentNode
			}
			currentNode.next = currentNode.next.next
			ll.length--
			return nil
		}
	}

	return fmt.Errorf("the value %v was not found in linked list", v)
}

// printAll prints all data present in the linked list
func (ll *linkedList) printAll() {
	currentNode := ll.head
	for i := 0; i < ll.length; i++ {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
}

func (ll *linkedList) reverse() {

	hd := ll.head
	ll.tail = hd
	var prev *node

	for hd != nil {
		hd, prev, hd.next = hd.next, hd, prev
	}

	ll.head = prev
}