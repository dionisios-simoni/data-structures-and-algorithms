package singly

import (
	"testing"
)

func TestPrepend(t *testing.T) {
	t.Run("prepending on empty linked list", func(t *testing.T) {
		ll := newLinkedList(t)
		populateList(t, 100, ll.prepend)

		assertHead(t, ll.head.data, 64)
		assertTail(t, ll.tail.data, 1)
		assertListLength(t, ll, 7)
	})

	t.Run("single node in linked list becomes head and tail", func(t *testing.T) {
		ll := newLinkedList(t)
		n := newNode(t, 10)

		ll.prepend(n.data)

		assertHead(t, ll.head.data, 10)
		assertTail(t, ll.tail.data, 10)
		assertListLength(t, ll, 1)
	})
}

func TestAppend(t *testing.T) {
	t.Run("appending on empty linked list", func(t *testing.T) {
		ll := newLinkedList(t)
		populateList(t, 100, ll.append)

		assertHead(t, ll.head.data, 1)
		assertTail(t, ll.tail.data, 64)
		assertListLength(t, ll, 7)
	})

	t.Run("single node in linked list becomes head and tail", func(t *testing.T) {
		ll := newLinkedList(t)
		n := newNode(t, 10)

		ll.append(n.data)

		assertHead(t, ll.head.data, 10)
		assertTail(t, ll.tail.data, 10)
		assertListLength(t, ll, 1)
	})

	t.Run("inserting into a linked list", func(t *testing.T) {

		ll := newLinkedList(t)

		t.Run("populates empty list as expected", func(t *testing.T) {
			index := 0
			for i := 1; i < 100; i *= 2 {
				n := &node{data: i}
				ll.insert(n.data, index)
				index++
			}

			assertHead(t, ll.head.data, 1)
			assertTail(t, ll.tail.data, 64)
			assertListLength(t, ll, 7)

		})

		t.Run("insert at the start of a pre-populated linked list", func(t *testing.T) {
			n := newNode(t, 1002)
			err := ll.insert(n.data, 0)

			assertHead(t, ll.head.data, 1002)
			assertTail(t, ll.tail.data, 64)
			assertListLength(t, ll, 8)
			assertNoError(t, err)
		})

		t.Run("insert at the end of a pre-populated linked list", func(t *testing.T) {
			n := newNode(t, 1005)
			err := ll.insert(n.data, ll.length)

			assertHead(t, ll.head.data, 1002)
			assertTail(t, ll.tail.data, 1005)
			assertListLength(t, ll, 9)
			assertNoError(t, err)
		})

		t.Run("insert at the middle of a pre-populated linked list", func(t *testing.T) {
			n := newNode(t, 1500)
			err := ll.insert(n.data, ll.length/2)

			assertHead(t, ll.head.data, 1002)
			assertTail(t, ll.tail.data, 1005)
			assertListLength(t, ll, 10)
			assertNoError(t, err)
		})
	})

	t.Run("search a node in the linked list", func(t *testing.T) {
		ll := newLinkedList(t)

		index := 0
		for i := 0; i < 10; i++ {
			n := newNode(t, i)
			ll.insert(n.data, index)
			index++
		}

		got, err := ll.search(5)
		if err != nil {
			t.Errorf("should have been able to find node in linkedlist but couldn't: %v", err)
		}

		if got == nil {
			t.Errorf("should have been able to find node in linkedlist but got: %v", got)
		}

		if got.data != 5 {
			t.Errorf("unexpected data retrieved from search, want %d but got %d", 5, got.data)
		}
	})

	t.Run("deleting a node", func(t *testing.T) {
		ll := newLinkedList(t)
		index := 0
		for i := 1; i < 100; i *= 2 {
			n := &node{data: i}
			ll.insert(n.data, index)
			index++
		}

		t.Run("when deleting head, the next node becomes the head", func(t *testing.T) {
			err := ll.delete(1)

			if err != nil {
				t.Errorf("should have not responded with error but did: %d", err)
			}

			assertHead(t, ll.head.data, 2)
			assertListLength(t, ll, 6)
		})

		t.Run("when deleting tail, the preceding node becomes the tail", func(t *testing.T) {
			err := ll.delete(64)

			if err != nil {
				t.Errorf("should have not responded with error but did: %d", err)
			}

			assertTail(t, ll.tail.data, 32)
			assertListLength(t, ll, 5)
		})

		t.Run("deleting a node from the middle of the linked list", func(t *testing.T) {
			err := ll.delete(8)

			if err != nil {
				t.Errorf("should have not responded with error but did: %d", err)
			}

			assertHead(t, ll.head.data, 2)
			assertTail(t, ll.tail.data, 32)
			assertListLength(t, ll, 4)
		})

		t.Run("length of 1 list", func(t *testing.T) {
			ll2 := newLinkedList(t)
			n := newNode(t, 10)
			ll2.append(n.data)

			err := ll2.delete(10)
			if err != nil {
				t.Errorf("returned an error: %v but shouldn't", err)
			}

			if ll2.head != nil || ll2.tail != nil {
				t.Errorf("list should reset but didn't")
			}
		})
	})

	t.Run("reversing a singly linked list", func(t *testing.T) {
		ll := newLinkedList(t)
		populateList(t, 100, ll.prepend)

		ll.reverse()
		assertHead(t, ll.head.data, 1)
		assertTail(t, ll.tail.data, 64)
	})
}

func newLinkedList(t testing.TB) *linkedList {
	t.Helper()
	return &linkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func populateList(t testing.TB, limit int, f func(v interface{})) {
	t.Helper()
	for i := 1; i < limit; i *= 2 {
		n := &node{
			data: i,
			next: nil,
		}
		f(n.data)
	}
}

func newNode(t testing.TB, data interface{}) *node {
	t.Helper()
	return &node{
		data: data,
		next: nil,
	}
}

// assert head data
func assertHead(t testing.TB, data interface{}, want int) {
	t.Helper()
	if got := data; got != want {
		t.Errorf("unexpected head data, want %d but got %d", want, got)
	}
}

// assert tail data
func assertTail(t testing.TB, data interface{}, want int) {
	t.Helper()
	if got := data; got != want {
		t.Errorf("unexpected tail data, want %d but got %d", want, got)
	}
}

// assert length of list
func assertListLength(t testing.TB, ll *linkedList, want int) {
	t.Helper()
	if got := ll.length; got != want {
		t.Errorf("unexpected linked list length, want: %d but got %d", want, got)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("an error occurred where it shouldn't: %v", err)
	}
}
