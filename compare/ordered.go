package compare

type Ordered interface {
	~float32 | ~float64 |
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~string
}

func OrderedLessCompare[T Ordered](a, b T) Result {
	if a < b {
		return LT
	} else if a == b {
		return EQ
	} else {
		return GT
	}
}

func OrderedGreaterCompare[T Ordered](a, b T) Result {
	return OrderedLessCompare(b, a)
}

func OrderedLessCompareF[T Ordered]() ICompare[T] {
	return WithFunc[T](OrderedLessCompare[T])
}

func OrderedGreaterCompareF[T Ordered]() ICompare[T] {
	return WithFunc[T](OrderedGreaterCompare[T])
}

func WithLessOrderedKey[T interface{ *K }, K any, O Ordered](keyBy func(T) O) ICompare[T] {
	return WithFunc[T](func(t, t2 T) Result {
		if t == nil && t2 == nil {
			return LT
		} else if t != nil && t2 == nil {
			return GT
		} else if t == nil && t2 != nil {
			return LT
		} else {
			return OrderedLessCompare(keyBy(t), keyBy(t2))
		}
	})
}

func WithGreaterOrderedKey[T interface{ *K }, K any, O Ordered](keyBy func(T) O) ICompare[T] {
	return WithFunc[T](func(t2, t T) Result {
		if t == nil && t2 == nil {
			return LT
		} else if t != nil && t2 == nil {
			return GT
		} else if t == nil && t2 != nil {
			return LT
		} else {
			return OrderedLessCompare(keyBy(t), keyBy(t2))
		}
	})
}
