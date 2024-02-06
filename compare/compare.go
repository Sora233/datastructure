package compare

type Result int

const (
	EQ Result = iota + 1
	LT
	GT
)

func (r Result) GTE() bool {
	return r == EQ || r == GT
}

func (r Result) LTE() bool {
	return r == EQ || r == LT
}

func (r Result) EQ() bool {
	return r == EQ
}

func (r Result) GT() bool {
	return r == GT
}

func (r Result) LT() bool {
	return r == LT
}

type ICompare[T any] interface {
	Compare(a, b T) Result
}

type Func[T any] func(a, b T) Result

func (f Func[T]) Compare(a, b T) Result {
	return f(a, b)
}

func WithFunc[T any](f func(T, T) Result) ICompare[T] {
	return Func[T](f)
}

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
