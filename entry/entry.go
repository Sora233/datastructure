package entry

import "github.com/Sora233/datastructure/compare"

type KV[K any, V any] struct {
	Key   K
	Value V
}

func NewKV[K any, V any](key K, value V) KV[K, V] {
	return KV[K, V]{Key: key, Value: value}
}

func Key[K any, V any](key K) KV[K, V] {
	return KV[K, V]{Key: key}
}

func OrderedKeyLessCompare[K compare.Ordered, V any](a, b KV[K, V]) compare.Result {
	return compare.OrderedLessCompare(a.Key, b.Key)
}

func OrderedKeyGreaterCompare[K compare.Ordered, V any](a, b KV[K, V]) compare.Result {
	return compare.OrderedGreaterCompare(a.Key, b.Key)
}

func OrderedKeyLessCompareF[K compare.Ordered, V any]() compare.ICompare[KV[K, V]] {
	return compare.WithFunc[KV[K, V]](OrderedKeyLessCompare[K, V])
}

func OrderedKeyGreaterCompareF[K compare.Ordered, V any]() compare.ICompare[KV[K, V]] {
	return compare.WithFunc[KV[K, V]](OrderedKeyGreaterCompare[K, V])
}
