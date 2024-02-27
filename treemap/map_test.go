package treemap

import (
	"fmt"
	"testing"
)

func TestNewMap(t *testing.T) {
	m := New[string, int]()
	m.Put("a", 1)
	m.Items()(func(s string, i int) bool {
		fmt.Println(s, i)
		return true
	})
}

type mystruct struct {
	a int
}

func (a *mystruct) Less(b *mystruct) bool {
	return a.a < b.a
}

func TestNewMapWithLessKeyBy(t *testing.T) {
	m2 := NewWithLessKeyBy[*mystruct, int](func(a *mystruct) int {
		return a.a
	})

	m2.Put(&mystruct{1}, 1)
	m2.Put(&mystruct{2}, 2)

	m2.Items()(func(m *mystruct, i int) bool {
		fmt.Println(m.a, i)
		return true
	})
}

func TestNewWithLesser(t *testing.T) {
	m2 := NewWithLesser[*mystruct, int]()

	m2.Put(&mystruct{1}, 1)
	m2.Put(&mystruct{2}, 2)

	m2.Items()(func(m *mystruct, i int) bool {
		fmt.Println(m.a, i)
		return true
	})
}
