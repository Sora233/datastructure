package treemap

import (
	"fmt"
	"testing"
)

func TestNewMap(t *testing.T) {
	m := NewMap[string, int]()
	m.Put("a", 1)
	m.Items()(func(s string, i int) bool {
		fmt.Println(s, i)
		return true
	})
}

type mystruct struct {
	a int
}

func myless(a, b *mystruct) bool {
	if a == nil && b == nil {
		return false
	} else if a == nil {
		return true
	} else if b == nil {
		return false
	} else {
		return a.a < b.a
	}
}

func TestNewMapWithLess(t *testing.T) {
	m := NewMapWithLess[*mystruct, int](myless)
	m.Put(&mystruct{a: 20}, 2)
	m.Put(&mystruct{a: 10}, 1)

	m.Items()(func(m *mystruct, i int) bool {
		fmt.Println(m.a, i)
		return true
	})
}
