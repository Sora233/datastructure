# datastructure

Zero Dependency Data Structure in golang with generic type.

Supported:
- BinarySearchTree
  - AVL Tree
  - Treap
- Heap
  - BinaryHeap

### Usage

#### TreeMap

go builtin `map` is a hash-based map, so it is not ordered. 
If you want to use a ordered map, you can use `bst.TreeMap`.

```go
package main

import (
	"fmt"
	"github.com/Sora233/datastructure/bst"
	"github.com/Sora233/datastructure/bst/avl"
	"github.com/Sora233/datastructure/entry"
)

func main() {
	m := bst.AsMap[int, int](avl.New[entry.KV[int, int]](entry.OrderedKeyLessCompareF[int, int]()))
	m.Put(1, 1)
	m.Put(2, 2)

	m.KeySet()(func(key int) bool {
		m.Put(key, key*2)
		return true
	})

	m.Put(100, 100)

	m.Items()(func(key int, value int) bool {
		fmt.Println(key, value)
		if key >= 2 {
			return false
		}
		return true
	})
}

// As Go1.22 introduced the new feature "range-over-function iterators"
// You can use the following code to iterate the treemap
// see https://tip.golang.org/wiki/RangefuncExperiment for detail.
func mainGo1_22() {
	m := bst.AsMap[int, int](avl.New[entry.KV[int, int]](entry.OrderedKeyLessCompareF[int, int]()))
	m.Put(1, 1)
	m.Put(2, 2)

	for key := range m.KeySet() {
		m.Put(key, key*2)
	}
	m.Put(100, 100)

	for key, value := range m.Items() {
		fmt.Println(key, value)
		if key >= 2 {
			break
		}
	}
}
```

### benchmark

see `internal/test/bst_benchmark_test.go` for benchmark code.

```text
$ go test -cpu=1 -bench='Bench' -benchtime=5s -benchmem ./internal/test/bst/bst_benchmark_test.go
goos: linux
goarch: amd64
cpu: 13th Gen Intel(R) Core(TM) i5-13400
BenchmarkBSTInt/random-insert||treap-int                 6657888              1184 ns/op               3 B/op          0 allocs/op
BenchmarkBSTInt/random-insert||treap-simple-int          7230386              1007 ns/op               3 B/op          0 allocs/op
BenchmarkBSTInt/random-insert||avl-int                   9168198               966.2 ns/op             2 B/op          0 allocs/op
--------------------------------------------------
BenchmarkBSTInt/order-insert||treap-int                 25129320               221.2 ns/op            61 B/op          0 allocs/op
BenchmarkBSTInt/order-insert||treap-simple-int          24854574               219.9 ns/op            61 B/op          0 allocs/op
BenchmarkBSTInt/order-insert||avl-int                   21139848               315.5 ns/op            60 B/op          0 allocs/op
--------------------------------------------------
BenchmarkBSTInt/random-find||treap-int                   6492459               949.0 ns/op             0 B/op          0 allocs/op
BenchmarkBSTInt/random-find||treap-simple-int            6945920               887.4 ns/op             0 B/op          0 allocs/op
BenchmarkBSTInt/random-find||avl-int                     9804452               623.1 ns/op             0 B/op          0 allocs/op
--------------------------------------------------
BenchmarkBSTInt/order-find||treap-int                   23705587               252.5 ns/op             0 B/op          0 allocs/op
BenchmarkBSTInt/order-find||treap-simple-int            20264535               264.6 ns/op             0 B/op          0 allocs/op
BenchmarkBSTInt/order-find||avl-int                     24643952               208.8 ns/op             0 B/op          0 allocs/op
--------------------------------------------------
BenchmarkBSTInt/random-insert-delete-90-10||treap-int            5211402              1185 ns/op               5 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-delete-90-10||treap-simple-int     6068265              1014 ns/op               5 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-delete-90-10||avl-int              6999675               976.6 ns/op             5 B/op          0 allocs/op
--------------------------------------------------
BenchmarkBSTInt/random-insert-delete-50-50||treap-int            4939803              1138 ns/op              15 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-delete-50-50||treap-simple-int     5753443               986.4 ns/op            15 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-delete-50-50||avl-int              5655814              1050 ns/op              15 B/op          0 allocs/op
--------------------------------------------------
BenchmarkBSTInt/random-insert-delete-10-90||treap-int            8976766               617.1 ns/op             5 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-delete-10-90||treap-simple-int    10567660               520.6 ns/op             5 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-delete-10-90||avl-int              8613273               604.5 ns/op             5 B/op          0 allocs/op
--------------------------------------------------
BenchmarkBSTInt/random-insert-find-90-10||treap-int              6214443              1133 ns/op               3 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-find-90-10||treap-simple-int       6843032               993.6 ns/op             3 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-find-90-10||avl-int                6946468               980.8 ns/op             3 B/op          0 allocs/op
--------------------------------------------------
BenchmarkBSTInt/random-insert-find-50-50||treap-int              5668333              1076 ns/op               0 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-find-50-50||treap-simple-int       6201783               961.1 ns/op             0 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-find-50-50||avl-int                7636495               872.7 ns/op             0 B/op          0 allocs/op
--------------------------------------------------
BenchmarkBSTInt/random-insert-find-10-90||treap-int              6330392               985.5 ns/op             0 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-find-10-90||treap-simple-int       6643911               925.6 ns/op             0 B/op          0 allocs/op
BenchmarkBSTInt/random-insert-find-10-90||avl-int                8983776               746.5 ns/op             0 B/op          0 allocs/op
--------------------------------------------------
PASS
ok      command-line-arguments  455.886s
```