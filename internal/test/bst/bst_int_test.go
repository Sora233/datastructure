package bst

import (
	"github.com/Sora233/datastructure/bst"
	"github.com/Sora233/datastructure/bst/avl"
	"github.com/Sora233/datastructure/bst/treap"
	"github.com/Sora233/datastructure/compare"
	"github.com/stretchr/testify/suite"
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
					r1, r2 := ts.tree.Prev(data)
					actual[0] = [2]any{r1, r2}
					r1, r2 = ts.tree.Next(data)
					actual[1] = [2]any{r1, r2}
					r1, r2 = ts.tree.FindOrNext(data)
					actual[2] = [2]any{r1, r2}
					r1, r2 = ts.tree.FindOrPrev(data)
					actual[3] = [2]any{r1, r2}
				}
			}
		})
	}

}
func TestBST(t *testing.T) {
	suite.Run(t, new(BSTIntSuite))
}
