package doubly

import (
	"errors"
	"fmt"
)

// linkedList represents a doubly linked list
type linkedList struct {
	head, tail *node
	length     int
}

// node represents the data stored in the linked list
type node struct {
	data interface{}
	prev *node
	next *node
}

// isEmpty returns true if the linked list has no elements.
func (l *linkedList) isEmpty() bool {
	return l.length == 0
}

// prepend adds a new node at the start of a linked list which becomes the head.
func (l *linkedList) prepend(v interface{}) {

	n := &node{data: v}

	if l.isEmpty() {
		l.head = n
		l.tail = n
	} else {
		l.head.prev = n
		n.next = l.head
		l.head = n
	}
	l.length++
}

// append adds a new node at the end of a linked list which becomes the tail.
func (l *linkedList) append(v interface{}) {

	n := &node{data: v}

	if l.isEmpty() {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
	}
	l.length++
}

// insert adds a new node in the linked list at a given position
// handling the possibility of the node becoming the new Head or Tail
func (l *linkedList) insert(v interface{}, position int) error {

	n := &node{data: v}

	switch {
	// Head
	case position == 0:
		l.prepend(n.data)
		return nil
	// Tail
	case position == l.length:
		l.append(n.data)
		return nil
	case position > 0 && position < l.length:
		data := l.head.next
		for i := 1; i < l.length; i++ {
			if i == position-1 {
				n.next = data.next
				n.prev = data
				data.next = n
				l.length++
			} else {
				data = data.next
			}
		}
	default:
		return errors.New("position out of bounds")
	}
	return nil
}

// search returns the node if found in the linked list or an error
func (l *linkedList) search(v interface{}) (*node, error) {

	if l.isEmpty() {
		return nil, errors.New("search failed as the linked list is empty")
	}

	if l.head.data == v {
		return l.head, nil
	}

	if l.tail.data == v {
		return l.tail, nil
	}

	n1 := l.head.next
	n2 := l.tail.prev
	for i, j := 0, l.length; i < j ; i, j = i + 1, j - 1 {
		if n1.data == v {
			return n1, nil
		} else if n2.data == v {
			return n2, nil
		}else {
			n1 = n1.next
			n2 = n2.prev
		}
	}

	return nil, fmt.Errorf("the element %v does not exist in linked list", v)
}

// delete removes a node from the linked list
// handling the scenarios of removing the head or tail
func (l *linkedList) delete(v interface{}) error {

	if l.isEmpty() {
		return errors.New("the linked list is empty")
	}

	if l.length == 1 && l.head.data == v {
		l.head = nil
		l.tail = nil
		l.length = 0
		return nil
	}

	if l.head.data == v {
		l.head.next.prev = nil
		l.head = l.head.next
		l.length--
		return nil
	}

	currentNode := l.head
	for i := 1; i <= l.length; i++ {
		if currentNode.next.data != v {
			currentNode = currentNode.next
		} else {
			if currentNode.next == l.tail {
				l.tail = currentNode
			}
			currentNode.next = currentNode.next.next
			l.length--
			return nil
		}
	}

	return fmt.Errorf("the value %v was not found in linked list", v)
}

// printAll prints all data present in the linked list
func (l *linkedList) printAll() {
	nd := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", nd.data)
		nd = nd.next
	}
}