package compare

type Less[T any] func(a, b T) bool

func LessF[T any](less Less[T]) ICompare[T] {
	return WithFunc[T](func(l1 T, l2 T) Result {
		r1 := less(l1, l2)
		r2 := less(l2, l1)
		if !r1 && !r2 {
			return EQ
		} else if r1 {
			return LT
		} else {
			return GT
		}
	})
}
