package bst

import "github.com/Sora233/datastructure"

// BinarySearchTree is the interface that wraps the basic operations of a binary search tree.
type BinarySearchTree[T any] interface {
	// Clear removes all elements from the tree.
	Clear()

	// Empty returns true if the tree is empty.
	Empty() bool

	// Size returns the size of the tree.
	Size() int

	// Insert inserts data into the tree.
	// If data already exists, the data will be overwritten.
	// return the old data if data is overwritten, or the zero value.
	Insert(data T) (old T, replaced bool)

	// InsertOrVisit insert data into the tree.
	// If data already exists, it remains unchanged and the datastructure.VisitFunc f will be called.
	// It is guaranteed that f is called at most once.
	InsertOrVisit(data T, f datastructure.VisitFunc[T])

	// InsertOrIgnore inserts data into the tree.
	// If data already exists, the operator is no effect.
	// return true if the data is inserted successfully.
	InsertOrIgnore(data T) (success bool)

	// Delete deletes data from the tree.
	// If data does not exist, the operator is no effect.
	// return true if the data is deleted successfully.
	Delete(data T) (old T, success bool)

	// DeleteIf deletes data from the tree if the datastructure.ConditionFunc f returns true.
	// If data does not exist or f return false, the operator is no effect.
	// return true if the data exists and is deleted successfully.
	// It is guaranteed that f is called at most once.
	DeleteIf(data T, f datastructure.ConditionFunc[T]) (success bool)

	// Find return the data and true if the data exists in the tree.
	// if the data doesn't exist, return the zero value and false.
	Find(data T) (res T, exists bool)

	// Exists return true if the data exists in the tree.
	Exists(data T) (exists bool)

	// Min return the minimum element in the tree.
	Min() (res T, exists bool)
	// Max return the maximum element in the tree.
	Max() (res T, exists bool)

	// Prev return the maximum element E that satisfies E < data.
	// If no such element, return zero value and false.
	Prev(data T) (res T, exists bool)

	// Next return the minimum element E that satisfies E > data,
	// If no such element, return zero value and false.
	Next(data T) (res T, exists bool)

	// FindOrNext return the minimum element E that satisfies E >= data,
	// If no such element, return zero value and false.
	FindOrNext(data T) (res T, exists bool)

	// FindOrPrev return the maximum element E that satisfies E <= data,
	// If no such element, return zero value and false.
	FindOrPrev(data T) (res T, exists bool)

	// Rank return the rank of data in the tree.
	// if the rank of data is N, it means there are (N-1) elements is smaller data
	// if data is the minimum element, the rank is 1
	Rank(data T) int

	// RankNth return the element that has the rank-th value.
	RankNth(rank int) (res T, exists bool)

	// Range iterate over all elements in the tree in ascending order.
	// The iteration will be interrupted if f returns false.
	// The compare-key should not be modified during the iteration.
	Range(f datastructure.ConditionFunc[T])

	// RangeS iterate over all elements E in the tree that satisfy E >= start in ascending order.
	// The iteration will be interrupted if f returns false.
	// The compare-key should not be modified during the iteration.
	RangeS(start T, f datastructure.ConditionFunc[T])

	// RangeSE iterate over all elements E in the tree that satisfy start <= E < end in ascending order.
	// The iteration will be interrupted if f returns false.
	// The compare-key should not be modified during the iteration.
	RangeSE(start, end T, f datastructure.ConditionFunc[T])

	// RangeE iterate over all elements E in the tree that satisfy E < end in ascending order.
	// The iteration will be interrupted if f returns false.
	// The compare-key should not be modified during the iteration.
	RangeE(end T, f datastructure.ConditionFunc[T])
}

type Countable interface {
	Count() int
}
