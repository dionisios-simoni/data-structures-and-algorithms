package separate_chaining_linked_list

import (
	"strings"
	"testing"
)

func TestHashTable(t *testing.T)  {

	ht := newHashTable()

	mustInsert("Alice", t, ht)
	mustInsert("Foo", t, ht)
	mustInsert("Bar", t, ht)

	t.Run("inserting unique keys no collision", func(t *testing.T) {

		index := hashFunc("Alice")
		want := 1
		if got := ht.bucket[index].length; got != want {
			t.Errorf("invalid linked list length, got %d but want %d", got, want)
		}

		index = hashFunc("Foo")
		if got := ht.bucket[index].length; got != want {
			t.Errorf("invalid linked list length, got %d but want %d", got, want)
		}

		index = hashFunc("Bar")
		if got := ht.bucket[index].length; got != want {
			t.Errorf("invalid linked list length, got %d but want %d", got, want)
		}
	})

	t.Run("inserting unique keys which result in collision", func(t *testing.T) {
		mustInsert("Bob", t, ht)
		mustInsert("FooBar", t, ht)

		index := hashFunc("Alice")
		want := 3
		if got := ht.bucket[index].length; got != want {
			t.Errorf("invalid linked list length, got %d but want %d", got, want)
		}
	})

	t.Run("inserting duplicate keys", func(t *testing.T) {

		err := ht.Insert("Bob")

		if err == nil {
			t.Errorf("inserting duplicate key should return error")
		}

		want := "the key already exists"
		if err != nil {
			strings.Contains(err.Error(), want)
		}

	})

	t.Run("search existing keys", func(t *testing.T) {
		if got := ht.Search("Alice"); got != true {
			t.Errorf("got search result %t but want %t", got, true)
		}
	})

	t.Run("search invalid keys", func(t *testing.T) {
		if got := ht.Search("Invalid"); got != false {
			t.Errorf("got search result %t but want %t", got, false)
		}
	})

	t.Run("delete existing key", func(t *testing.T) {
		index := hashFunc("Alice")
		want := 2

		ht.Delete("Alice")

		if got := ht.bucket[index].length; got != want {
			t.Errorf("invalid linked list length, got %d but want %d", got, want)
		}
	})

	t.Run("delete head", func(t *testing.T) {
		index := hashFunc("FooBar")
		want := 1

		ht.Delete("FooBar")

		if got := ht.bucket[index].length; got != want {
			t.Errorf("invalid linked list length, got %d but want %d", got, want)
		}
	})

	t.Run("delete invalid key", func(t *testing.T) {
		index := hashFunc("Alice")
		want := 1

		ht.Delete("Invalid")

		if got := ht.bucket[index].length; got != want {
			t.Errorf("invalid linked list length, got %d but want %d", got, want)
		}
	})
}

func mustInsert(key string, tb testing.TB, ht *hashTable) {
	tb.Helper()
	if err := ht.Insert(key); err != nil {
		panic(err)
	}
}