package avl

import (
	"github.com/Sora233/datastructure/bst"
)

type Iterator[T any] struct {
	node *Node[T]
}

func (i *Iterator[T]) Get() (data T, exists bool) {
	if i != nil && i.node != nil {
		data = i.node.getValue()
		exists = true
	}
	return
}

func (i *Iterator[T]) Next() bst.Iterator[T] {
	if i == nil || i.node == nil {
		return &Iterator[T]{node: nil}
	}
	node := i.node

	if node.r != nil {
		node = node.r
		for node.l != nil {
			node = node.l
		}
		return &Iterator[T]{node: node}
	}
	var fa *Node[T]
	for {
		fa = node.getFa()
		if fa == nil || node != fa.r {
			break
		}
		node = fa
	}
	return &Iterator[T]{node: fa}
}

func (i *Iterator[T]) Prev() bst.Iterator[T] {
	if i == nil || i.node == nil {
		return &Iterator[T]{node: nil}
	}
	node := i.node
	if node.l != nil {
		node = node.l
		for node.r != nil {
			node = node.r
		}
		return &Iterator[T]{node: node}
	}

	var fa *Node[T]
	for {
		fa = node.getFa()
		if fa == nil || node != fa.l {
			break
		}
		node = fa
	}
	return &Iterator[T]{node: fa}
}
