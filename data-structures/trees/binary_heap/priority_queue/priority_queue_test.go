package priority_queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := &PriorityQueue{}

	t.Run("enqueue values in priority queue", func(t *testing.T) {
		for i := 1; i <= 20; i++ {
			pq.Enqueue(i * 5)
		}

		got := pq.heap.Values
		want := []int{100, 95, 70, 85, 90, 55, 65, 50, 80, 45, 40, 10, 30, 25, 60, 5, 35, 20, 75, 15}

		for i := 0; i < len(want); i++ {
			if want[i] != got[i] {
				t.Errorf("order of binary heap elements is invalid got %v but want %v", got, want)
			}
		}
	})

	t.Run("dequeue returns removes the highest value from the queue", func(t *testing.T) {
		values := pq.heap.Values

		want := 100
		for i := 0; i < len(values); i++ {
			got := pq.Dequeue()
			if want != got {
				t.Errorf("unexpected value was dequeued, expected: %d but got: %d", want, got)
			}
			want -= 5
		}
	})

}
