package bst

import (
	"github.com/Sora233/datastructure/bst"
	"github.com/Sora233/datastructure/bst/avl"
	"github.com/Sora233/datastructure/bst/treap"
	"github.com/Sora233/datastructure/compare"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
)

type BSTIntSuite struct {
	suite.Suite
	treeSet []struct {
		name string
		tree bst.BinarySearchTree[int]
	}
}

func (s *BSTIntSuite) SetupTest() {
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

func (s *BSTIntSuite) TearDownTest() {
	for _, st := range s.treeSet {
		st.tree.Clear()
	}
}

func (s *BSTIntSuite) TearDownSubTest() {
	for _, st := range s.treeSet {
		st.tree.Clear()
	}
}

func (s *BSTIntSuite) TestInsert() {
	var testcases = []struct {
		name     string
		data     []int
		expected []int
	}{
		{
			name:     "test insert",
			data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "test insert duplicate",
			data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			for _, ts := range s.treeSet {
				for _, v := range tc.data {
					ts.tree.Insert(v)
				}
				var actual []int
				ts.tree.Range(func(i int) bool {
					actual = append(actual, i)
					return true
				})
				s.EqualValues(tc.expected, actual)
			}
		})
	}
}

func (s *BSTIntSuite) TestDelete() {
	var testcases = []struct {
		name     string
		prepared []int
		data     []int
		expected []int
	}{
		{
			name:     "test delete",
			prepared: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			data:     []int{1, 3, 5, 7, 9},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "test redundant delete",
			prepared: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			data:     []int{1, 3, 5, 7, 9, 1, 3, 5, 7, 9, 11, 13, 15, 1, 1, 1, 1, 1, 3, 3, 3},
			expected: []int{2, 4, 6, 8, 10},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			for _, ts := range s.treeSet {
				for _, i := range tc.prepared {
					ts.tree.Insert(i)
				}
				for _, i := range tc.data {
					ts.tree.Delete(i)
				}
				var actual []int
				ts.tree.Range(func(n int) bool {
					actual = append(actual, n)
					return true
				})
				s.EqualValues(tc.expected, actual)
			}
		})
	}
}

func (s *BSTIntSuite) TestPrevAndNextAndFindOrNextAndFindOrPrev() {
	var testcases = []struct {
		name     string
		prepared []int
		data     []int
		expected [][4][2]any
	}{
		{
			name:     "test prev and next and findOrNext and findOrPrev",
			prepared: []int{1},
			data:     []int{1, 0, 2},
			expected: [][4][2]any{
				{
					{0, false},
					{0, false},
					{1, true},
					{1, true},
				},
				{
					{0, false},
					{1, true},
					{1, true},
					{0, false},
				},
				{
					{1, true},
					{0, false},
					{0, false},
					{1, true},
				},
			},
		},
		{
			name:     "test prev and next and findOrNext and findOrPrev 2",
			prepared: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			data:     []int{0, 1, 5, 10, 11},
			expected: [][4][2]any{
				{
					{0, false},
					{1, true},
					{1, true},
					{0, false},
				},
				{
					{0, false},
					{2, true},
					{1, true},
					{1, true},
				},
				{
					{4, true},
					{6, true},
					{5, true},
					{5, true},
				},
				{
					{9, true},
					{0, false},
					{10, true},
					{10, true},
				},
				{
					{10, true},
					{0, false},
					{0, false},
					{10, true},
				},
			},
		},
	}
	for _, tc := range testcases {
		s.Run(tc.name, func() {
			for _, ts := range s.treeSet {
				for _, i := range tc.prepared {
					ts.tree.Insert(i)
				}
				var actual [4][2]any
				for _, data := range tc.data {
					r1, r2 := ts.tree.Prev(data).Get()
					actual[0] = [2]any{r1, r2}
					r1, r2 = ts.tree.Next(data).Get()
					actual[1] = [2]any{r1, r2}
					r1, r2 = ts.tree.FindOrNext(data).Get()
					actual[2] = [2]any{r1, r2}
					r1, r2 = ts.tree.FindOrPrev(data).Get()
					actual[3] = [2]any{r1, r2}
				}
			}
		})
	}
}

var bstOps = []string{
	"Size",
	"Insert",
	"InsertOrIgnore",
	"Delete",
	"DeleteIf",
	"Find",
	"Exists",
	"Min",
	"Max",
	"Prev",
	"Next",
	"FindOrNext",
	"FindOrPrev",
	"Rank",
	"RankNth",
	"Range",
	"RangeS",
	"RangeSE",
	"RangeE",
}

type bstOp struct {
	op string
	p1 int
	p2 int
}

func (op *bstOp) do(tree bst.BinarySearchTree[int]) (any, any) {
	switch op.op {
	case "Size":
		return tree.Size(), nil
	case "Insert":
		return tree.Insert(op.p1)
	case "InsertOrIgnore":
		return tree.InsertOrIgnore(op.p1), nil
	case "Delete":
		return tree.Delete(op.p1)
	case "DeleteIf":
		return tree.DeleteIf(op.p1, func(i int) bool {
			return op.p2%2 == 0
		}), nil
	case "Find":
		return tree.Find(op.p1).Get()
	case "Exists":
		return tree.Exists(op.p1), nil
	case "Min":
		return tree.Min().Get()
	case "Max":
		return tree.Max().Get()
	case "Prev":
		return tree.Prev(op.p1).Get()
	case "Next":
		return tree.Next(op.p1).Get()
	case "FindOrNext":
		return tree.FindOrNext(op.p1).Get()
	case "FindOrPrev":
		return tree.FindOrPrev(op.p1).Get()
	case "Rank":
		return tree.Rank(op.p1), nil
	case "RankNth":
		return tree.RankNth(op.p1).Get()
	case "Range":
		var data []int
		tree.Range(func(i int) bool {
			data = append(data, i)
			return true
		})
		return data, nil
	case "RangeS":
		var data []int
		tree.RangeS(op.p1, func(i int) bool {
			data = append(data, i)
			return true
		})
		return data, nil
	case "RangeSE":
		var data []int
		tree.RangeSE(op.p1, op.p2, func(i int) bool {
			data = append(data, i)
			return true
		})
		return data, nil
	case "RangeE":
		var data []int
		tree.RangeE(op.p1, func(i int) bool {
			data = append(data, i)
			return true
		})
		return data, nil
	default:
		panic("impossible")
	}
}

func (s *BSTIntSuite) TestFuzzy() {
	for _, maxKey := range []int{10, 10000, 1000000} {
		for _, ts := range s.treeSet {
			ts.tree.Clear()
		}
		for i := 0; i < 300000; i++ {
			op := genBstOps(maxKey)
			var results [][2]any
			for _, ts := range s.treeSet {
				r1, r2 := op.do(ts.tree)
				results = append(results, [2]any{r1, r2})
			}
			for i := 1; i < len(results); i++ {
				s.EqualValuesf(results[0], results[i], "%v not match with %v, op %v, p %v %v", s.treeSet[0].name, s.treeSet[i].name, op.op, op.p1, op.p2)
			}
		}
	}

}

func genBstOps(maxKey int) *bstOp {
	op := &bstOp{
		op: bstOps[rand.Intn(len(bstOps))],
		p1: rand.Intn(maxKey),
		p2: rand.Intn(maxKey),
	}
	if op.p1 > op.p2 {
		op.p1, op.p2 = op.p2, op.p1
	}

	return op
}

func TestBST(t *testing.T) {
	suite.Run(t, new(BSTIntSuite))
}
