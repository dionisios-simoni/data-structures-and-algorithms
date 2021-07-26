package seperate_chaining

// HashTable is a data structure which contains buckets of Entry
type HashTable struct {
	bucket [][]Entry
}

// Entry represents the data to be stored in the hash table.
type Entry struct {
	key   string
	value interface{}
}

// NewHashTable is a constructor function
func NewHashTable(size int) *HashTable {
	return &HashTable{make([][]Entry, size)}
}

// hashFunc is the hashing function used which produces an index position to allocate for new entries.
func (ht *HashTable) hashFunc(key string) int {
	hash := 0
	for i, k := range key {
		hash = (hash + int(k)*i) % len(ht.bucket)
	}
	return hash
}

// Get returns the value of a stored Entry given the provided key exists
func (ht *HashTable) Get(key string) (interface{}, bool) {
	bucketAddress := ht.hashFunc(key)
	entries := ht.bucket[bucketAddress]

	for _, e := range entries {
		if e.key == key {
			return e.value, true
		}
	}
	return nil, false
}

// Set updates the value of an Entry if the key already exists or creates a new Entry.
func (ht *HashTable) Set(key string, value interface{}) {
	bucketAddress := ht.hashFunc(key)

	if ht.bucket[bucketAddress] == nil {
		ht.bucket[bucketAddress] = []Entry{{key, value}}
	} else {
		for i, e := range ht.bucket[bucketAddress] {
			if e.key == key {
				ht.bucket[bucketAddress][i].value = value
				return
			}
		}
		ht.bucket[bucketAddress] = append(ht.bucket[bucketAddress], Entry{key: key, value: value})
	}
}

