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
	if a > b {
		return LT
	} else if a == b {
		return EQ
	} else {
		return GT
	}
}

func OrderedLessCompareF[T Ordered]() ICompare[T] {
	return WithFunc[T](OrderedLessCompare[T])
}

func OrderedGreaterCompareF[T Ordered]() ICompare[T] {
	return WithFunc[T](OrderedGreaterCompare[T])
}
