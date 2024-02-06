package binaryheap

import (
	"github.com/Sora233/datastructure"
	"github.com/Sora233/datastructure/compare"
)

type BinaryHeap[T any] struct {
	data []T
	cmp  compare.ICompare[T]
	size int
	// for DecreaseKey, we need locate the node for the given key.
	m map[any]int
}

func New[T any](cmp compare.ICompare[T]) *BinaryHeap[T] {
	tree := &BinaryHeap[T]{
		cmp:  cmp,
		data: make([]T, 1),
		m:    make(map[any]int),
	}

	return tree
}

// Clear clears the heap
// Time Complex: O(1)
func (b *BinaryHeap[T]) Clear() {
	b.data = make([]T, 1)
	b.size = 0
	b.m = make(map[any]int)
}

// Top returns the minimum data in the heap
// If the heap is empty, return zero value and false
// Time Complex: O(1)
func (b *BinaryHeap[T]) Top() (res T, exists bool) {
	if b.size == 0 {
		return
	}
	res = b.data[1]
	exists = true
	return
}

// Size return the size of the heap
// Time Complex: O(1)
func (b *BinaryHeap[T]) Size() (size int) {
	return b.size
}

// Empty return true if the heap is empty
// Time Complex: O(1)
func (b *BinaryHeap[T]) Empty() bool {
	return b.size == 0
}

// Insert inserts data into the heap
// Time Complex: O(logN)
func (b *BinaryHeap[T]) Insert(data T) (old T, replaced bool) {
	b.insert(data, func(idx int) {
		old = b.data[idx]
		replaced = true
		b.data[idx] = data
		return
	})
	return
}

// InsertOrVisit
// Time Complex: O(logN)
func (b *BinaryHeap[T]) InsertOrVisit(data T, f datastructure.VisitFunc[T]) {
	b.insert(data, func(idx int) {
		f(b.data[idx])
	})
}

// InsertOrIgnore
// Time Complex: O(logN)
func (b *BinaryHeap[T]) InsertOrIgnore(data T) (success bool) {
	return b.insert(data, nil)
}

// Pop returns the minimum data and delete it from heap
// Time Complex: O(logN)
func (b *BinaryHeap[T]) Pop() (res T, exists bool) {
	b.PopIf(func(t T) bool {
		res = t
		exists = true
		return true
	})
	return
}

// PopIf deletes the minimum data from the heap if the datastructure.ConditionFunc f returns true
// Time Complex: O(logN)
func (b *BinaryHeap[T]) PopIf(f datastructure.ConditionFunc[T]) (success bool) {
	if b.size == 0 {
		return
	}
	if f != nil && !f(b.data[1]) {
		return
	}
	res := b.data[1]
	success = true
	b.swap(1, b.size)
	b.size--
	b.data = b.data[:b.size+1]
	delete(b.m, res)
	b.down(1)
	return
}

// DecreaseKey
// Time Complex: O(logN)
func (b *BinaryHeap[T]) DecreaseKey(oldKey T, change datastructure.ModifyFunc[T]) (success bool) {
	if idx, found := b.m[oldKey]; found {
		newKey := change(b.data[idx])
		if b.cmp.Compare(oldKey, newKey).LT() {
			panic("binaryheap: DecreaseKey change with bigger key")
		}
		delete(b.m, oldKey)
		b.data[idx] = newKey
		b.m[b.data[idx]] = idx
		b.up(idx)
		success = true
	}
	return
}

// Private method

func (b *BinaryHeap[T]) compare(idx1, idx2 int) compare.Result {
	return b.cmp.Compare(b.data[idx1], b.data[idx2])
}

func (b *BinaryHeap[T]) swap(idx1, idx2 int) {
	b.data[idx1], b.data[idx2] = b.data[idx2], b.data[idx1]
	b.m[b.data[idx1]] = idx1
	b.m[b.data[idx2]] = idx2
}

func (b *BinaryHeap[T]) up(idx int) {
	for idx > 1 && b.compare(idx, idx/2).LT() {
		b.swap(idx, idx/2)
		idx /= 2
	}
}

func (b *BinaryHeap[T]) down(idx int) {
	for (idx << 1) <= b.size {
		t := idx << 1
		if t+1 <= b.size && b.compare(t+1, t).LT() {
			t++
		}
		if b.compare(t, idx).GTE() {
			break
		}
		b.swap(idx, t)
		idx = t
	}
}

func (b *BinaryHeap[T]) insert(data T, f func(idx int)) bool {
	if idx, found := b.m[data]; found {
		if f != nil {
			f(idx)
		}
		return false
	}
	b.size++
	b.data = append(b.data, data)
	b.m[data] = b.size
	b.up(b.size)
	return true
}
