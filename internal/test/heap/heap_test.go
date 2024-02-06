package heap

import (
	"fmt"
	"github.com/Sora233/datastructure/compare"
	"github.com/Sora233/datastructure/heap/binaryheap"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HeapSuite struct {
	suite.Suite
}

func (s *HeapSuite) TestHeap() {

}

func TestHeap(t *testing.T) {
	suite.Run(t, new(HeapSuite))
}

func TestXX(t *testing.T) {
	h := binaryheap.New[int](compare.OrderedLessCompareF[int]())
	for i := 10; i >= 1; i-- {
		h.Insert(i)
	}
	fmt.Println(h.Top())
	for i := 10; i >= 1; i-- {
		h.DecreaseKey(i, func(oldKey int) int {
			return oldKey - 100
		})
		fmt.Println(h.Top())
	}
	for i := 10; i >= 1; i-- {
		h.Insert(i)
	}
	fmt.Println("///////////")
	for !h.Empty() {
		fmt.Println(h.Pop())
	}
	for i := 1; i <= 10; i++ {
		h.Insert(i)
	}
	fmt.Println(h.Size())
	fmt.Println(h.Top())
	h.Clear()
	fmt.Println(h.Size())
	fmt.Println(h.Top())
}
