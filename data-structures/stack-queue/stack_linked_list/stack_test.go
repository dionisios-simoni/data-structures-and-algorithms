package stack_linked_list

import "testing"

func TestStack(t *testing.T)  {
	s := stack{}

	t.Run("push new items", func(t *testing.T) {

		for i := 10; i <= 100; i += 10 {
			s.push(i)
		}

		want := 10
		if got := s.length; got != want {
			t.Errorf("invalid stack length, want %d but got %d", want, got)
		}
	})

	t.Run("pop items from the stack", func(t *testing.T) {
		s.pop()
		s.pop()

		want := 8
		if got := s.length; got != want {
			t.Errorf("invalid stack length, want %d but got %d", want, got)
		}
	})

	t.Run("peek the top of the stack", func(t *testing.T) {
		want := 80
		if got := s.peek(); got != want {
			t.Errorf("invalid top of stack item, want: %d but got: %d", want, got)
		}
	})
}
