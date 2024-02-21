package bst

import (
	"github.com/Sora233/datastructure/bst"
	"github.com/Sora233/datastructure/bst/avl"
	"github.com/Sora233/datastructure/bst/treap"
	"github.com/Sora233/datastructure/compare"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"sort"
	"testing"
)

type BSTIteratorSuite struct {
	suite.Suite
	treeSet []struct {
		name string
		tree bst.BinarySearchTree[int]
	}
}

func (s *BSTIteratorSuite) SetupTest() {
	s.treeSet = append(s.treeSet, struct {
		name string
		tree bst.BinarySearchTree[int]
	}{
		name: "treap",
		tree: treap.New[int](compare.OrderedLessCompareF[int]()),
	})
	s.treeSet = append(s.treeSet, struct {
		name string
		tree bst.BinarySearchTree[int]
	}{
		name: "AVL",
		tree: avl.New[int](compare.OrderedLessCompareF[int]()),
	})
}

func (s *BSTIteratorSuite) TearDownTest() {
	for _, st := range s.treeSet {
		st.tree.Clear()
	}
}

func (s *BSTIteratorSuite) TearDownSubTest() {
	for _, st := range s.treeSet {
		st.tree.Clear()
	}
}

func (s *BSTIteratorSuite) TestIterator() {
	var testcases = []struct {
		name     string
		data     []int
		expected []int
	}{
		{
			name:     "test iterator",
			data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "test iterator duplicate",
			data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, st := range s.treeSet {
		for _, tc := range testcases {
			for _, v := range tc.data {
				st.tree.Insert(v)
			}
			idx := 0
			it := st.tree.Min()
			for {
				v, ok := it.Get()
				if !ok {
					break
				}
				s.EqualValuesf(tc.expected[idx], v, "iterator error %v vs %v", tc.expected[idx], v)
				it = it.Next()
				idx++
			}
			s.EqualValuesf(len(tc.expected), idx, "iterator error %v vs %v", len(tc.expected), idx)

			idx = len(tc.expected)
			it = st.tree.Max()
			for {
				v, ok := it.Get()
				if !ok {
					break
				}
				idx--
				s.EqualValuesf(tc.expected[idx], v, "iterator error %v vs %v", tc.expected[idx], v)
				it = it.Prev()
			}
			s.EqualValuesf(0, idx, "iterator error %v vs %v", 0, idx)
		}
	}
}

func (s *BSTIteratorSuite) TestRandomIter() {
	var data []int
	for i := -10000; i <= 10000; i++ {
		data = append(data, i)
	}
	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	for _, st := range s.treeSet {
		tree := st.tree
		for _, v := range data {
			tree.Insert(v)
		}
		for i := 0; i < 10000; i++ {
			r := rand.Intn(19990) - 9995
			it := tree.Find(r)
			v, ok := it.Get()
			s.Truef(ok, "%v should be found in tree", r)
			s.EqualValuesf(r, v, "Find error %v vs %v", r, v)

			next := it
			for j := 1; j <= 5; j++ {
				next = next.Next()
				v, ok := next.Get()
				s.Truef(ok, "%v should be found in tree", r+j)
				s.EqualValuesf(r+j, v, "Next error %v vs %v", r+j, v)
			}

			prev := it
			for j := 1; j <= 5; j++ {
				prev = prev.Prev()
				v, ok := prev.Get()
				s.Truef(ok, "%v should be found in tree", r-j)
				s.EqualValuesf(r-j, v, "Prev error %v vs %v", r-j, v)
			}
		}
	}
}

type arr struct {
	data []int
}

func (a *arr) insert(num int) {
	pos := sort.Search(len(a.data), func(i int) bool {
		return a.data[i] >= num
	})
	if pos < len(a.data) && a.data[pos] == num {
		return
	}
	if pos == 0 {
		a.data = append([]int{num}, a.data...)
	} else if pos == len(a.data) {
		a.data = append(a.data, num)
	} else {
		a.data = append(a.data[:pos], append([]int{num}, a.data[pos:]...)...)
	}
}

func (a *arr) delete(num int) {
	pos := sort.Search(len(a.data), func(i int) bool {
		return a.data[i] >= num
	})
	if pos < len(a.data) && a.data[pos] == num {
		a.data = append(a.data[:pos], a.data[pos+1:]...)
	}
}

func (a *arr) findOrNext(num int) int {
	return sort.Search(len(a.data), func(i int) bool {
		return a.data[i] >= num
	})
}

func (s *BSTIteratorSuite) TestIteratorWithModify() {
	var a arr
	rand.Seed(1)

	for i := 0; i < 100000; i++ {
		op := rand.Intn(3)
		num := rand.Intn(10)
		switch op {
		case 0: // insert
			for _, ts := range s.treeSet {
				ts.tree.Insert(num)
			}
			a.insert(num)
		case 1: // delete
			for _, ts := range s.treeSet {
				ts.tree.Delete(num)
			}
			a.delete(num)
		case 2: // iter
			pos := a.findOrNext(num)
			var its []bst.Iterator[int]
			for _, ts := range s.treeSet {
				its = append(its, ts.tree.FindOrNext(num))
			}
			var next []bst.Iterator[int]
			var prev []bst.Iterator[int]
			for _, it := range its {
				next = append(next, it)
				prev = append(prev, it)
			}

			if pos == len(a.data) {
				// not found in arr
				for j := range its {
					_, ok := its[j].Get()
					s.Falsef(ok, "%v: %v should not be found in tree", s.treeSet[j].name, num)
				}
				break
			}

			for i := 0; i < 5 && pos+i < len(a.data); i++ {
				for j := range next {
					v, ok := next[j].Get()
					s.Truef(ok, "%v: %v should be found in tree", s.treeSet[j].name, a.data[pos+i])
					s.EqualValuesf(a.data[pos+i], v, "%v: next error %v vs %v", s.treeSet[j].name, a.data[pos+i], v)
					if i == 0 && a.data[pos+i] != v {
						s.treeSet[j].tree.FindOrNext(num)
					}
					next[j] = next[j].Next()
				}
			}
			for i := 0; i < 5 && pos-i >= 0; i++ {
				for j := range prev {
					v, ok := prev[j].Get()
					s.Truef(ok, "%v: %v should be found in tree", s.treeSet[j].name, a.data[pos-i])
					s.EqualValuesf(a.data[pos-i], v, "%v: prev error %v vs %v", s.treeSet[j].name, a.data[pos-i], v)
					if i == 0 && a.data[pos-i] != v {
						s.treeSet[j].tree.FindOrNext(num)
					}
					prev[j] = prev[j].Prev()
				}
			}
		}
	}
}

func TestBSTIteratorSuite(t *testing.T) {
	suite.Run(t, new(BSTIteratorSuite))
}
