package treap

import (
	"github.com/Sora233/datastructure/allocator"
)

type option[T any] struct {
	alloc allocator.IAllocator[Node[T]]
	r     func() int
}

type OptionFunc[T any] func(*option[T])

// WithAllocator set the allocator of the tree
func WithAllocator[T any](alloc allocator.IAllocator[Node[T]]) OptionFunc[T] {
	return func(o *option[T]) {
		o.alloc = alloc
	}
}

// WithRand set the rand of the treap
// It may be useful to perform a specific result
func WithRand[T any](r func() int) OptionFunc[T] {
	return func(o *option[T]) {
		o.r = r
	}
}

func getOption[T any](opts []OptionFunc[T]) *option[T] {
	var opt = new(option[T])
	for _, o := range opts {
		o(opt)
	}
	return opt
}
