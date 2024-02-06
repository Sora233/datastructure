package heap

import "github.com/Sora233/datastructure"

type Heap[T any] interface {
	// Clear clears the heap.
	Clear()

	// Size returns the size of the heap.
	Size() (size int)

	// Empty returns true if the heap is empty.
	Empty() bool

	// Top return the minimum data in the heap.
	// If the heap is empty, return zero value and false.
	Top() (res T, exists bool)

	// Insert inserts data into the heap.
	// If data already exists, the data will be overwritten.
	// return the old data if data is overwritten, or the zero value.
	Insert(data T) (old T, replaced bool)

	// InsertOrVisit insert data into the heap.
	// If data already exists, it remains unchanged and the datastructure.VisitFunc f will be called.
	// It is guaranteed that f is called at most once.
	InsertOrVisit(data T, f datastructure.VisitFunc[T])

	// InsertOrIgnore inserts data into the heap.
	// If data already exists, the operator is no effect.
	// return true if the data is inserted successfully.
	InsertOrIgnore(data T) (success bool)

	// Pop returns the minimum data and delete it from heap.
	// if the heap is empty, return zero value and false.
	Pop() (res T, exists bool)

	// PopIf deletes data from the heap if the datastructure.ConditionFunc f returns true.
	// If data does not exist or f return false, the operator is no effect.
	// return true if the data exists and is deleted successfully.
	// It is guaranteed that f is called at most once.
	PopIf(f datastructure.ConditionFunc[T]) (success bool)

	// DecreaseKey modify the oldKey by function change.
	// This operator panics when new key is greater than oldKey
	// return true if the key exists and modify success.
	DecreaseKey(oldKey, change datastructure.ModifyFunc[T]) (success bool)
}
