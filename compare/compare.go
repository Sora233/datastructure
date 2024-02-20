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
