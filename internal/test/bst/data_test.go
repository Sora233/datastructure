package bst

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/Sora233/datastructure/bst"
	"github.com/Sora233/datastructure/bst/treap_nore"
	"github.com/Sora233/datastructure/entry"
	"github.com/stretchr/testify/suite"
	"io"
	"strconv"
	"strings"
	"testing"
)

//go:embed testdata/loj104/input*
var loj104 embed.FS

type BSTDataSuite struct {
	suite.Suite
	treeSet []struct {
		name string
		tree bst.BinarySearchTree[entry.Duplicate[int]]
	}
}

func (s *BSTDataSuite) SetupTest() {
	s.treeSet = append(s.treeSet, struct {
		name string
		tree bst.BinarySearchTree[entry.Duplicate[int]]
	}{
		name: "treap",
		tree: treap.New[entry.Duplicate[int]](entry.OrderedDuplicateLessCompareF[int]()),
	})

}

func (s *BSTDataSuite) TearDownTest() {
	for _, t := range s.treeSet {
		t.tree.Clear()
	}
}

func (s *BSTDataSuite) TearDownSubTest() {
	for _, t := range s.treeSet {
		t.tree.Clear()
	}
}

func (s *BSTDataSuite) TestLOJ104() {
	dir, err := loj104.ReadDir("testdata/loj104")
	s.Nil(err)
	count := len(dir) / 2
	var testcase []struct {
		name   string
		input  string
		output string
	}
	for i := 0; i < count; i++ {
		input, err := loj104.ReadFile(fmt.Sprintf("testdata/loj104/input%v.in", i))
		s.Nil(err)
		output, err := loj104.ReadFile(fmt.Sprintf("testdata/loj104/input%v.out", i))
		s.Nil(err)
		testcase = append(testcase, struct {
			name   string
			input  string
			output string
		}{
			name:   fmt.Sprintf("case %v", i),
			input:  string(input),
			output: string(output),
		})
	}
	for _, tc := range testcase {
		for _, ts := range s.treeSet {
			s.Run(tc.name+"_"+ts.name, func() {
				var result strings.Builder
				s.runCase(ts.tree, strings.NewReader(tc.input), &result)
				s1 := strings.Split(tc.output, "\n")
				s2 := strings.Split(result.String(), "\n")
				s.EqualValues(len(s1), len(s2))
				for i := range s1 {
					s.Equalf(s1[i], s2[i], "line %v mismatched, %v vs %v", i, s1[i], s2[i])
				}
			})
		}
	}
}

func (s *BSTDataSuite) runCase(tree bst.BinarySearchTree[entry.Duplicate[int]], input io.Reader, w *strings.Builder) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	var n, op, x int
	fmt.Sscanf(scanner.Text(), "%d", &n)
	var cnt = 0
	for ; n != 0; n-- {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &op, &x)
		switch op {
		case 1:
			tree.InsertOrVisit(entry.NewDuplicate(x), entry.InsertDuplicate[int](1))
		case 2:
			tree.DeleteIf(entry.NewDuplicate(x), entry.DeleteDuplicate[int](1))
		case 3:
			w.WriteString(strconv.Itoa(tree.Rank(entry.NewDuplicate(x))))
			w.WriteRune('\n')
			cnt++
		case 4:
			ent, _ := tree.RankNth(x).Get()
			w.WriteString(strconv.Itoa(ent.Key))
			w.WriteRune('\n')
			cnt++
		case 5:
			ent, _ := tree.Prev(entry.NewDuplicate(x)).Get()
			w.WriteString(strconv.Itoa(ent.Key))
			w.WriteRune('\n')
			cnt++
		case 6:
			ent, _ := tree.Next(entry.NewDuplicate(x)).Get()
			w.WriteString(strconv.Itoa(ent.Key))
			w.WriteRune('\n')
			cnt++
		}
	}
}

func TestBSTDataSuite(t *testing.T) {
	suite.Run(t, new(BSTDataSuite))
}
