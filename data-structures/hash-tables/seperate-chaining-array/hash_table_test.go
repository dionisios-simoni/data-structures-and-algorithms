package seperate_chaining_array

import (
	"testing"
)

var noCollisionCases = []struct {
	key   string
	value int
	want  int
}{
	{"Apple", 100, 100},
	{"Banana", 2, 2},
	{"Clementine", 50, 50},
	{"Orange", 10, 10},
}

var collisionCases = []struct {
	key   string
	value int
	want  int
}{
	{"Cherry", 2000, 2000},
	{"Cantaloupe", 50, 50},
	{"Clementine", 250, 250},
	{"Cucumber", 3150, 3150},
}

func TestHashTable(t *testing.T) {
	t.Run("no collision cases", func(t *testing.T) {

		ht := NewHashTable(5)

		for _, tt := range noCollisionCases {
			ht.Set(tt.key, tt.value)
			if got, _ := ht.Get(tt.key); got != tt.want {
				t.Errorf("unexpected result, want: %d but got %d", tt.want, got)
			}

			for i := range ht.bucket {
				if got := len(ht.bucket[i]); got > 1 {
					t.Errorf("unexpected size of entries, want: 1 but got %d", got)
				}
			}
		}
	})

	t.Run("collision cases", func(t *testing.T) {
		ht := NewHashTable(5)

		ht.Set("Apple", 100)
		ht.Set("Banana", 500)

		for _, tt := range collisionCases {
			ht.Set(tt.key, tt.value)
			if got, _ := ht.Get(tt.key); got != tt.want {
				t.Errorf("unexpected result, want: %d but got %d", tt.want, got)
			}
		}

		for i := range ht.bucket {
			if ht.bucket[i] != nil {
				if got := len(ht.bucket[i]); got < 2 {
					t.Errorf("unexpected size of entries, want length > 1 got %d", got)
				}
			}
		}
	})

	t.Run("duplicate key only updates the value of existing entry", func(t *testing.T) {
		ht := NewHashTable(5)

		ht.Set("Apple", 100)
		ht.Set("Banana", 500)

		for _, tt := range collisionCases {
			ht.Set(tt.key, tt.value)
			if got, _ := ht.Get(tt.key); got != tt.want {
				t.Errorf("unexpected result, want: %d but got %d", tt.want, got)
			}
		}

		want := 150
		ht.Set("Cantaloupe", want)

		if _, ok := ht.Get("Cantaloupe"); !ok {
			t.Errorf("Get should return the value stored for the provided key")
		}

		if got, _ := ht.Get("Cantaloupe"); got != want {
			t.Errorf("Get should return the updated value stored for the provided key")
		}
	})
}
