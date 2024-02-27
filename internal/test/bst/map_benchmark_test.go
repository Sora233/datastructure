package bst

import (
	"fmt"
	"github.com/Sora233/datastructure/bst/avl"
	"github.com/Sora233/datastructure/bst/treap"
	"github.com/Sora233/datastructure/entry"
	"github.com/Sora233/datastructure/treemap"
	"math/rand"
	"testing"
)

func benchmarkInsertWithMapIntString(m TreeMap[int, string]) func(*testing.B) {
	r := rand.New(rand.NewSource(999888777))
	return func(b *testing.B) {
		var data = make([]int, b.N)
		var v = make([]string, b.N)
		for i := 0; i < b.N; i++ {
			data[i] = r.Intn(Max)
			v[i] = randString(10)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			m.Put(data[i], v[i])
		}
	}
}

func BenchmarkMap(b *testing.B) {
	var mapSet = []struct {
		name string
		tree TreeMap[int, string]
	}{
		{
			name: "stdMap",
			tree: newStdMap[int, string](),
		},
		{
			name: "treap",
			tree: treemap.As[int, string](treap.New[entry.KV[int, string]](entry.OrderedKeyLessCompareF[int, string]())),
		},
		{
			name: "avl",
			tree: treemap.As[int, string](avl.New[entry.KV[int, string]](entry.OrderedKeyLessCompareF[int, string]())),
		},
	}

	var testcases = []struct {
		name string
		f    func(tree TreeMap[int, string]) func(*testing.B)
	}{
		{
			name: "random-insert",
			f:    benchmarkInsertWithMapIntString,
		},
	}
	for _, tc := range testcases {
		for _, t := range mapSet {
			b.Run(fmt.Sprintf("%v||%v", tc.name, t.name), tc.f(t.tree))
		}
		fmt.Println("--------------------------------------------------")
	}
}
