package queue_array

import (
	"testing"
)

func TestStack(t *testing.T)  {
	s := Queue{}

	t.Run("Enqueue new Items", func(t *testing.T) {

		for i := 10; i <= 100; i += 10 {
			s.enqueue(i)
		}

		want := 10
		if got := len(s.Items); got != want {
			t.Errorf("invalid Queue length, want %d but got %d", want, got)
		}
	})

	t.Run("Dequeue Items from the Queue", func(t *testing.T) {
		s.Dequeue()
		s.Dequeue()

		want := 8
		if got := len(s.Items); got != want {
			t.Errorf("invalid Queue length, want %d but got %d", want, got)
		}
	})

	t.Run("Peek the current first element of the Queue", func(t *testing.T) {
		want := 30
		if got := s.Peek(); got != want {
			t.Errorf("invalid first element of Queue, want: %d but got: %d", want, got)
		}
	})
}
