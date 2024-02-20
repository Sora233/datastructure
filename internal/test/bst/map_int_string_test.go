package bst

import (
	"github.com/Sora233/datastructure/bst/avl"
	"github.com/Sora233/datastructure/bst/treap"
	"github.com/Sora233/datastructure/compare"
	"github.com/Sora233/datastructure/entry"
	"github.com/Sora233/datastructure/treemap"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
)

func randString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type stdMap[K comparable, V any] struct {
	m map[K]V
}

func (s *stdMap[K, V]) Put(key K, value V) (old V, replaced bool) {
	old, replaced = s.m[key]
	s.m[key] = value
	return
}

func (s *stdMap[K, V]) PutIfAbsent(key K, value V) (success bool) {
	if _, found := s.m[key]; !found {
		s.m[key] = value
		success = true
	}
	return
}

func (s *stdMap[K, V]) Get(key K) (value V, exists bool) {
	value, exists = s.m[key]
	return
}

func (s *stdMap[K, V]) Delete(key K) (value V, exists bool) {
	value, exists = s.m[key]
	delete(s.m, key)
	return
}

func (s *stdMap[K, V]) Len() int {
	return len(s.m)
}

func (s *stdMap[K, V]) Clear() {
	s.m = make(map[K]V)
}

func newStdMap[K compare.Ordered, V any]() treemap.TreeMap[K, V] {
	return &stdMap[K, V]{
		m: make(map[K]V),
	}
}

type MapIntStringSuite struct {
	suite.Suite
	maps   []treemap.TreeMap[int, string]
	maxKey []int
	N      int
}

func (s *MapIntStringSuite) SetupTest() {
	s.maxKey = []int{1, 10, 100, 10000, 1000000, 100000000}
	s.N = 300000
	s.maps = append(s.maps, newStdMap[int, string]())
	s.maps = append(s.maps, treemap.AsMap[int, string](treap.New[entry.KV[int, string]](entry.OrderedKeyLessCompareF[int, string]())))
	s.maps = append(s.maps, treemap.AsMap[int, string](avl.New[entry.KV[int, string]](entry.OrderedKeyLessCompareF[int, string]())))
}

func (s *MapIntStringSuite) TearDownSubTest() {
	for _, m := range s.maps {
		m.Clear()
	}
}

type op[K compare.Ordered, V any] struct {
	op    string
	key   K
	value V
}

func (o *op[K, V]) do(m treemap.TreeMap[K, V]) (any, any) {
	switch o.op {
	case "get":
		a1, a2 := m.Get(o.key)
		return a1, a2
	case "put":
		a1, a2 := m.Put(o.key, o.value)
		return a1, a2
	case "len":
		a1 := m.Len()
		return a1, nil
	case "delete":
		a1, a2 := m.Delete(o.key)
		return a1, a2
	case "put_if_absent":
		a1 := m.PutIfAbsent(o.key, o.value)
		return a1, nil
	default:
		panic("impossible")
	}
}

var ops = []string{
	"get",
	"put",
	"len",
	"delete",
	"put_if_absent",
}

func genOp(maxKey int) *op[int, string] {
	return &op[int, string]{
		op:    ops[rand.Intn(len(ops))],
		key:   rand.Intn(maxKey),
		value: randString(10),
	}
}

func (s *MapIntStringSuite) TestFuzzy() {
	for _, maxKey := range s.maxKey {
		for i := 0; i < s.N; i++ {
			op := genOp(maxKey)
			var results [][2]any
			for _, m := range s.maps {
				r1, r2 := op.do(m)
				results = append(results, [2]any{r1, r2})
			}
			for i := 1; i < len(results); i++ {
				s.EqualValues(results[0], results[i])
			}
		}
	}
}

func TestBSTMapSuite(t *testing.T) {
	suite.Run(t, new(MapIntStringSuite))
}
