package treemap

import (
	"github.com/Sora233/datastructure/bst"
	"github.com/Sora233/datastructure/bst/avl"
	"github.com/Sora233/datastructure/compare"
	"github.com/Sora233/datastructure/entry"
)

// TreeMap is the interface that wraps the basic operations of a map.
// It is not safe for concurrent use.
type TreeMap[K any, V any] interface {
	Put(key K, value V) (old V, replaced bool)
	PutIfAbsent(key K, value V) (success bool)
	Get(key K) (value V, exists bool)
	Delete(key K) (value V, exists bool)
	Len() int
	Clear()
	KeySet() func(yield func(K) bool)
	Items() func(yield func(K, V) bool)
	Next(key K) (value V, exists bool)
	Prev(key K) (value V, exists bool)
	Rank(key K) int
	RankNth(n int) (key K, value V, exists bool)
}

type treeMap[K any, V any] struct {
	tree bst.BinarySearchTree[entry.KV[K, V]]
}

func (t *treeMap[K, V]) Put(key K, value V) (old V, replaced bool) {
	e, replaced := t.tree.Insert(entry.NewKV(key, value))
	old = e.Value
	return
}

func (t *treeMap[K, V]) PutIfAbsent(key K, value V) (success bool) {
	success = t.tree.InsertOrIgnore(entry.NewKV(key, value))
	return
}
func (t *treeMap[K, V]) Get(key K) (value V, exists bool) {
	e, exists := t.tree.Find(entry.Key[K, V](key)).Get()
	value = e.Value
	return
}

func (t *treeMap[K, V]) Delete(key K) (value V, exists bool) {
	e, exists := t.tree.Delete(entry.Key[K, V](key))
	value = e.Value
	return
}

func (t *treeMap[K, V]) Len() int {
	return t.tree.Size()
}

func (t *treeMap[K, V]) Clear() {
	t.tree.Clear()
}

func (t *treeMap[K, V]) KeySet() func(yield func(K) bool) {
	return func(yield func(K) bool) {
		t.tree.Range(func(e entry.KV[K, V]) bool {
			return yield(e.Key)
		})
	}
}

func (t *treeMap[K, V]) Items() func(yield func(K, V) bool) {
	return func(yield func(K, V) bool) {
		t.tree.Range(func(e entry.KV[K, V]) bool {
			return yield(e.Key, e.Value)
		})
	}
}

func (t *treeMap[K, V]) Next(key K) (value V, exists bool) {
	e, exists := t.tree.Next(entry.Key[K, V](key)).Get()
	value = e.Value
	return
}
func (t *treeMap[K, V]) Prev(key K) (value V, exists bool) {
	e, exists := t.tree.Prev(entry.Key[K, V](key)).Get()
	value = e.Value
	return
}
func (t *treeMap[K, V]) Rank(key K) int {
	return t.tree.Rank(entry.Key[K, V](key))
}
func (t *treeMap[K, V]) RankNth(n int) (key K, value V, exists bool) {
	e, exists := t.tree.RankNth(n).Get()
	key = e.Key
	value = e.Value
	return
}

func New[K compare.Ordered, V any]() TreeMap[K, V] {
	return NewWithCompare[K, V](compare.OrderedLessCompareF[K]())
}

func NewWithLesser[K interface{ Less(K) bool }, V any]() TreeMap[K, V] {
	return NewWithCompare[K, V](compare.LesserF[K]())
}

func NewWithLessKeyBy[PK interface{ *K }, V any, K any, O compare.Ordered](keyBy func(PK) O) TreeMap[PK, V] {
	return NewWithCompare[PK, V](compare.WithLessOrderedKey[PK](keyBy))
}

func NewWithCompare[K any, V any](keyCompare compare.ICompare[K]) TreeMap[K, V] {
	return As[K, V](avl.New(entry.KeyCompareWrapper[K, V](keyCompare)))
}

// As Create a TreeMap base on the BinarySearchTree
func As[K any, V any](tree bst.BinarySearchTree[entry.KV[K, V]]) TreeMap[K, V] {
	if tree == nil {
		panic("As: tree is nil")
	}
	m := &treeMap[K, V]{
		tree: tree,
	}
	return m
}
