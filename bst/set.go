package bst

type TreeSet[T any] interface {
	Put(elem T) (old T, replaced bool)
	PutIfAbsent(elem T) (success bool)
	Get(elem T) (res T, exists bool)
	Delete(elem T) (res T, exists bool)
	Len() int
	Clear()
	Items() func(yield func(T) bool)
}

type treeSet[T any] struct {
	tree BinarySearchTree[T]
}

func (t *treeSet[T]) Put(elem T) (old T, replaced bool) {
	old, replaced = t.tree.Insert(elem)
	return
}

func (t *treeSet[T]) PutIfAbsent(elem T) (success bool) {
	success = t.tree.InsertOrIgnore(elem)
	return
}

func (t *treeSet[T]) Get(elem T) (res T, exists bool) {
	res, exists = t.tree.Find(elem)
	return
}

func (t *treeSet[T]) Delete(elem T) (res T, exists bool) {
	res, exists = t.tree.Delete(elem)
	return
}

func (t *treeSet[T]) Len() int {
	return t.tree.Size()
}

func (t *treeSet[T]) Clear() {
	t.tree.Clear()
}

func (t *treeSet[T]) Items() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		t.tree.Range(func(e T) bool {
			return yield(e)
		})
	}
}

// AsSet Create a TreeSet base on the BinarySearchTree
func AsSet[T any](tree BinarySearchTree[T]) TreeSet[T] {
	s := &treeSet[T]{
		tree: tree,
	}
	return s
}
