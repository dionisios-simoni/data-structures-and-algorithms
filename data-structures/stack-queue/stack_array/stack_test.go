package stack_array

import (
	"testing"
)

func TestStack(t *testing.T)  {
	s := Stack{}

	t.Run("Push new Items", func(t *testing.T) {

		for i := 10; i <= 100; i += 10 {
			s.Push(i)
		}

		want := 10
		if got := len(s.Items); got != want {
			t.Errorf("invalid Stack length, want %d but got %d", want, got)
		}
	})

	t.Run("Pop Items from the Stack", func(t *testing.T) {
		s.Pop()
		s.Pop()

		want := 8
		if got := len(s.Items); got != want {
			t.Errorf("invalid Stack length, want %d but got %d", want, got)
		}
	})

	t.Run("peek the top of the Stack", func(t *testing.T) {
		want := 80
		if got := s.peek(); got != want {
			t.Errorf("invalid top of Stack item, want: %d but got: %d", want, got)
		}
	})
}
