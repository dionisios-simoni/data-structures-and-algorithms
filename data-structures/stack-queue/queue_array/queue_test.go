package stack_array

import (
	"testing"
)

func TestStack(t *testing.T)  {
	s := queue{}

	t.Run("enqueue new items", func(t *testing.T) {

		for i := 10; i <= 100; i += 10 {
			s.enqueue(i)
		}

		want := 10
		if got := len(s.items); got != want {
			t.Errorf("invalid queue length, want %d but got %d", want, got)
		}
	})

	t.Run("dequeue items from the queue", func(t *testing.T) {
		s.dequeue()
		s.dequeue()

		want := 8
		if got := len(s.items); got != want {
			t.Errorf("invalid queue length, want %d but got %d", want, got)
		}
	})

	t.Run("peek the current first element of the queue", func(t *testing.T) {
		want := 30
		if got := s.peek(); got != want {
			t.Errorf("invalid first element of queue, want: %d but got: %d", want, got)
		}
	})
}
