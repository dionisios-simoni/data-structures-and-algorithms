package stack_linked_list

import "testing"

func TestStack(t *testing.T)  {
	s := queue{}

	t.Run("enqueue new items", func(t *testing.T) {

		for i := 10; i <= 100; i += 10 {
			s.enqueue(i)
		}

		want := 10
		if got := s.length; got != want {
			t.Errorf("invalid queue length, want %d but got %d", want, got)
		}
	})

	t.Run("dequeue items from the queue", func(t *testing.T) {
		s.dequeue()
		s.dequeue()

		want := 8
		if got := s.length; got != want {
			t.Errorf("invalid queue length, want %d but got %d", want, got)
		}
	})

	t.Run("peek the start of the queue", func(t *testing.T) {
		want := 30
		if got := s.peek(); got != want {
			t.Errorf("invalid start of queue item, want: %d but got: %d", want, got)
		}
	})
}
