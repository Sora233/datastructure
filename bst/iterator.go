package bst

// Iterator is the interface that wraps the basic operations of a iterator.
// all returned Iterator should always be non-nil, even if the data does not exist.
// After the tree is modified, the Iterator is no longer valid.
type Iterator[T any] interface {
	// Get return the data.
	// If the data does not exist, return zero value and false.
	Get() (data T, exists bool)

	// Next return the next Iterator.
	Next() Iterator[T]

	// Prev return the previous Iterator.
	Prev() Iterator[T]
}
