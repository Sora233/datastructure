package entry

import (
	"github.com/Sora233/datastructure"
	"github.com/Sora233/datastructure/compare"
)

type Duplicate[T any] struct {
	Key   T
	count *int
}

func (du Duplicate[T]) Count() int {
	if du.count == nil {
		return 0
	}
	return *du.count
}

func (du Duplicate[T]) Add(delta uint) {
	*du.count += int(delta)
}

func (du Duplicate[T]) Sub(delta uint) {
	*du.count -= int(delta)
}

func NewDuplicate[T any](key T) Duplicate[T] {
	count := 1
	return Duplicate[T]{Key: key, count: &count}
}

func InsertDuplicate[T any](count uint) datastructure.VisitFunc[Duplicate[T]] {
	return func(d Duplicate[T]) {
		d.Add(count)
	}
}

func DeleteDuplicate[T any](count uint) datastructure.ConditionFunc[Duplicate[T]] {
	return func(d Duplicate[T]) bool {
		d.Sub(count)
		return d.Count() <= 0
	}
}

func NewDuplicateCount[T any](key T, count uint) Duplicate[T] {
	c := int(count)
	return Duplicate[T]{Key: key, count: &c}
}

func OrderedDuplicateLessCompare[K compare.Ordered](a, b Duplicate[K]) compare.Result {
	return compare.OrderedLessCompare(a.Key, b.Key)
}

func OrderedDuplicateGreaterCompare[K compare.Ordered](a, b Duplicate[K]) compare.Result {
	return compare.OrderedGreaterCompare(a.Key, b.Key)
}

func OrderedDuplicateLessCompareF[K compare.Ordered]() compare.ICompare[Duplicate[K]] {
	return compare.WithFunc[Duplicate[K]](OrderedDuplicateLessCompare[K])
}

func OrderedDuplicateGreaterCompareF[K compare.Ordered]() compare.ICompare[Duplicate[K]] {
	return compare.WithFunc[Duplicate[K]](OrderedDuplicateGreaterCompare[K])
}
