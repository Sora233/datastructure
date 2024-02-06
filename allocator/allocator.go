package allocator

type IAllocator[T any] interface {
	Allocate() *T
	Release()
}

type SimpleAllocators[T any] struct{}

func (s *SimpleAllocators[T]) Release() {}

func (s *SimpleAllocators[T]) Allocate() *T {
	return new(T)
}

type BlockAllocator[T any] struct {
	blockSize int
	block     []T
	pos       int
}

func (b *BlockAllocator[T]) Release() {
	b.block = make([]T, b.blockSize)
	b.pos = 0
}

func (b *BlockAllocator[T]) Allocate() *T {
	if b.pos == b.blockSize {
		b.block = make([]T, b.blockSize)
		b.pos = 0
	}
	res := &b.block[b.pos]
	b.pos++
	return res
}

func NewBlockAllocator[T any](blockSize int) *BlockAllocator[T] {
	if blockSize <= 0 {
		panic("block size must be greater than 0")
	}
	return &BlockAllocator[T]{
		blockSize: blockSize,
		block:     make([]T, blockSize),
		pos:       0,
	}
}

func NewSimpleAllocator[T any]() *SimpleAllocators[T] {
	return &SimpleAllocators[T]{}
}
