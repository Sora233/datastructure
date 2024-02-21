package bst

import (
	"fmt"
	"github.com/Sora233/datastructure/allocator"
	"github.com/Sora233/datastructure/bst"
	"github.com/Sora233/datastructure/bst/avl"
	"github.com/Sora233/datastructure/bst/treap"
	"github.com/Sora233/datastructure/compare"
	"math/rand"
	"testing"
)

const (
	Max = 1000000
)

func benchmarkInsertWithTreeInt(tree bst.BinarySearchTree[int]) func(*testing.B) {
	r := rand.New(rand.NewSource(999888777))
	return func(b *testing.B) {
		var data = make([]int, b.N)
		for i := 0; i < b.N; i++ {
			data[i] = r.Intn(Max)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tree.Insert(data[i])
		}
	}
}

func benchmarkOrderInsertWithTreeInt(tree bst.BinarySearchTree[int]) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree.Insert(i)
		}
	}
}

func benchmarkFindWithTreeInt(tree bst.BinarySearchTree[int]) func(*testing.B) {
	r := rand.New(rand.NewSource(999888777))
	for i := 0; i < Max; i++ {
		tree.Insert(r.Intn(Max))
	}
	return func(b *testing.B) {
		var data = make([]int, b.N)
		for i := 0; i < b.N; i++ {
			data[i] = r.Intn(Max)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tree.Find(data[i])
		}
	}
}

func benchmarkOrderFindWithTreeInt(tree bst.BinarySearchTree[int]) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree.Insert(i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tree.Find(i)
		}
	}
}

func benchmarkInsertAndDelete(insertRate float64) func(tree bst.BinarySearchTree[int]) func(*testing.B) {
	return func(tree bst.BinarySearchTree[int]) func(*testing.B) {
		r := rand.New(rand.NewSource(999888777))
		return func(b *testing.B) {
			var f = make([]float64, b.N)
			var data = make([]int, b.N)
			for i := 0; i < b.N; i++ {
				f[i] = r.Float64()
				data[i] = r.Intn(Max)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if f[i] < insertRate {
					tree.Insert(data[i])
				} else {
					tree.Delete(data[i])
				}
			}
		}
	}
}

func benchmarkInsertAndFind(insertRate float64) func(tree bst.BinarySearchTree[int]) func(*testing.B) {
	return func(tree bst.BinarySearchTree[int]) func(*testing.B) {
		r := rand.New(rand.NewSource(999888777))
		return func(b *testing.B) {
			var f = make([]float64, b.N)
			var data = make([]int, b.N)
			for i := 0; i < b.N; i++ {
				f[i] = r.Float64()
				data[i] = r.Intn(Max)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if f[i] < insertRate {
					tree.Insert(data[i])
				} else {
					tree.Find(data[i])
				}
			}
		}
	}
}

func BenchmarkBSTInt(b *testing.B) {
	var treeSet = []struct {
		name string
		tree bst.BinarySearchTree[int]
	}{
		{
			name: "treap-int",
			tree: treap.New[int](compare.OrderedLessCompareF[int]()),
		},
		{
			name: "treap-simple-int",
			tree: treap.New[int](compare.OrderedLessCompareF[int](), treap.WithAllocator[int](allocator.NewSimpleAllocator[treap.Node[int]]())),
		},
		{
			name: "avl-int",
			tree: avl.New[int](compare.OrderedLessCompareF[int]()),
		},
	}
	var testcase = []struct {
		name string
		f    func(tree bst.BinarySearchTree[int]) func(*testing.B)
	}{
		{
			name: "random-insert",
			f:    benchmarkInsertWithTreeInt,
		},
		{
			name: "order-insert",
			f:    benchmarkOrderInsertWithTreeInt,
		},
		{
			name: "random-find",
			f:    benchmarkFindWithTreeInt,
		},
		{
			name: "order-find",
			f:    benchmarkOrderFindWithTreeInt,
		},
		{
			name: "random-insert-delete-90-10",
			f:    benchmarkInsertAndDelete(0.9),
		},
		{
			name: "random-insert-delete-50-50",
			f:    benchmarkInsertAndDelete(0.5),
		},
		{
			name: "random-insert-delete-10-90",
			f:    benchmarkInsertAndDelete(0.1),
		},
		{
			name: "random-insert-find-90-10",
			f:    benchmarkInsertAndFind(0.9),
		},
		{
			name: "random-insert-find-50-50",
			f:    benchmarkInsertAndFind(0.5),
		},
		{
			name: "random-insert-find-10-90",
			f:    benchmarkInsertAndFind(0.1),
		},
	}
	for _, tc := range testcase {
		for _, t := range treeSet {
			b.Run(fmt.Sprintf("%v||%v", tc.name, t.name), tc.f(t.tree))
		}
		fmt.Println("--------------------------------------------------")
	}
}
