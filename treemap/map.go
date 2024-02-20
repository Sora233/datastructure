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
	e, exists := t.tree.Find(entry.Key[K, V](key))
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

func NewMap[K compare.Ordered, V any]() TreeMap[K, V] {
	return AsMap[K, V](avl.New[entry.KV[K, V]](entry.OrderedKeyLessCompareF[K, V]()))
}

func NewMapWithLess[K any, V any](less compare.Less[K]) TreeMap[K, V] {
	return AsMap[K, V](
		avl.New(
			entry.KeyCompareWrapper[K, V](compare.LessF[K](less)),
		),
	)
}

func NewMapWithCompare[K any, V any](keyCompare compare.ICompare[K]) TreeMap[K, V] {
	return AsMap[K, V](
		avl.New(
			entry.KeyCompareWrapper[K, V](keyCompare),
		),
	)
}

// AsMap Create a TreeMap base on the BinarySearchTree
func AsMap[K any, V any](tree bst.BinarySearchTree[entry.KV[K, V]]) TreeMap[K, V] {
	if tree == nil {
		panic("AsMap: tree is nil")
	}
	m := &treeMap[K, V]{
		tree: tree,
	}
	return m
}
