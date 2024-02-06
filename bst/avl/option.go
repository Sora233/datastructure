package avl

import (
	"github.com/Sora233/datastructure/allocator"
)

type option[T any] struct {
	alloc allocator.IAllocator[Node[T]]
}

type OptionFunc[T any] func(*option[T])

// WithAllocator set the allocator of the tree
func WithAllocator[T any](alloc allocator.IAllocator[Node[T]]) OptionFunc[T] {
	return func(o *option[T]) {
		o.alloc = alloc
	}
}

func getOption[T any](opts []OptionFunc[T]) *option[T] {
	var opt = new(option[T])
	for _, o := range opts {
		o(opt)
	}
	return opt
}
