package avl

import (
	"github.com/Sora233/datastructure"
	"github.com/Sora233/datastructure/allocator"
	"github.com/Sora233/datastructure/bst"
	"github.com/Sora233/datastructure/compare"
)

type AVL[T any] struct {
	root           *Node[T]
	alloc          allocator.IAllocator[Node[T]]
	cmp            compare.ICompare[T]
	countableCheck bool
}

func New[T any](cmp compare.ICompare[T], opts ...OptionFunc[T]) *AVL[T] {
	var opt = getOption(opts)
	tree := &AVL[T]{
		alloc: opt.alloc,
		cmp:   cmp,
	}

	if tree.alloc == nil {
		tree.alloc = allocator.NewBlockAllocator[Node[T]](64)
	}
	var init T
	if _, ok := any(init).(bst.Countable); ok {
		tree.countableCheck = true
	}
	return tree
}

// Clear clears the AVL.
func (t *AVL[T]) Clear() {
	t.root = nil
	t.alloc.Release()
}

// Empty return true if the AVL is empty.
func (t *AVL[T]) Empty() bool {
	return t.root.getSize() == 0
}

// Size return the size of the AVL.
func (t *AVL[T]) Size() int {
	return t.root.getSize()
}

// Insert inserts data into the AVL.
// If data already exists, the data will be overwritten.
// return the old data if data is overwritten, or the zero value.
func (t *AVL[T]) Insert(data T) (old T, replaced bool) {
	t.root = t.insert(t.root, data, func(n *Node[T]) {
		old = n.val
		replaced = true
		n.setVal(data, t.countableCheck)
	})
	return
}

// InsertOrVisit insert data into the AVL.
// If data already exists, the visit function f will be called instead.
// It is guaranteed that f is called at most once.
func (t *AVL[T]) InsertOrVisit(data T, f datastructure.VisitFunc[T]) {
	t.root = t.insert(t.root, data, nodeVisitWrap(f))
}

// InsertOrIgnore inserts data into the AVL.
// If data already exists, the operator is no effect.
// return true if the data is inserted successfully.
func (t *AVL[T]) InsertOrIgnore(data T) (success bool) {
	success = true
	t.root = t.insert(t.root, data, func(n *Node[T]) {
		success = false
	})
	return
}

// Delete deletes data from the treap.
// If data does not exist, the operator is no effect.
// return true if the data is deleted successfully.
func (t *AVL[T]) Delete(data T) (old T, success bool) {
	t.root = t.delete(t.root, data, func(n *Node[T]) bool {
		old = n.getValue()
		success = true
		return true
	})
	return
}

// DeleteIf deletes data from the AVL if the condition function f returns true.
// If data does not exist or f return false, the operator is no effect.
// return true if the data exists and is deleted successfully.
// It is guaranteed that f is called at most once.
func (t *AVL[T]) DeleteIf(data T, f datastructure.ConditionFunc[T]) (success bool) {
	t.root = t.delete(t.root, data, func(n *Node[T]) bool {
		res := f(n.getValue())
		success = res
		return res
	})
	return
}

func (t *AVL[T]) Find(data T) (res T, exists bool) {
	enterLeft := func(root *Node[T]) bool {
		return t.cmp.Compare(root.val, data).GT()
	}
	enterCur := func(root *Node[T]) bool {
		return t.cmp.Compare(root.val, data).EQ()
	}
	enterRight := func(root *Node[T]) bool {
		return t.cmp.Compare(root.val, data).LT()
	}
	t.root.postorder(enterLeft, enterRight, enterCur, func(n *Node[T]) bool {
		res = n.val
		exists = true
		return false
	})
	return
}

func (t *AVL[T]) Exists(data T) (exists bool) {
	_, exists = t.Find(data)
	return
}

func (t *AVL[T]) Min() (res T, exists bool) {
	if t.Empty() {
		return
	}
	node := t.root
	exists = true
	for node.l != nil {
		node = node.l
	}
	res = node.getValue()
	return
}

func (t *AVL[T]) Max() (res T, exists bool) {
	if t.Empty() {
		return
	}
	node := t.root
	exists = true
	for node.r != nil {
		node = node.r
	}
	res = node.getValue()
	return
}

func (t *AVL[T]) Prev(data T) (res T, exists bool) {
	enterRight := func(node *Node[T]) bool {
		return t.cmp.Compare(node.getValue(), data).LT()
	}
	enterLeft := func(node *Node[T]) bool {
		return t.cmp.Compare(node.getValue(), data).GTE()
	}
	enterCur := func(node *Node[T]) bool {
		return t.cmp.Compare(node.getValue(), data).LT()
	}
	t.root.reversePostorder(
		enterRight,
		enterLeft,
		enterCur,
		func(n *Node[T]) bool {
			res = n.val
			exists = true
			return false
		},
	)
	return
}

func (t *AVL[T]) Next(data T) (res T, exists bool) {
	enterLeft := func(node *Node[T]) bool {
		return t.cmp.Compare(node.getValue(), data).GT()
	}
	enterRight := func(node *Node[T]) bool {
		return t.cmp.Compare(node.getValue(), data).LTE()
	}
	enterCur := func(node *Node[T]) bool {
		return t.cmp.Compare(node.getValue(), data).GT()
	}
	t.root.postorder(
		enterLeft,
		enterRight,
		enterCur,
		func(n *Node[T]) bool {
			res = n.val
			exists = true
			return false
		},
	)
	return
}

func (t *AVL[T]) FindOrNext(data T) (res T, exists bool) {
	enterLeft := func(node *Node[T]) bool {
		return t.cmp.Compare(node.getValue(), data).GT()
	}
	enterRight := func(node *Node[T]) bool {
		return t.cmp.Compare(node.getValue(), data).LT()
	}
	t.root.postorder(
		enterLeft,
		enterRight,
		trueNodeConditionFunc[T],
		func(n *Node[T]) bool {
			res = n.val
			exists = true
			return false
		},
	)
	return
}

// FindOrPrev return the maximum element E that satisfies E <= data,
// If no such element, return zero value and false.
func (t *AVL[T]) FindOrPrev(data T) (res T, exists bool) {
	enterRight := func(node *Node[T]) bool {
		return t.cmp.Compare(node.getValue(), data).LT()
	}
	enterLeft := func(node *Node[T]) bool {
		return t.cmp.Compare(node.getValue(), data).GT()
	}
	t.root.reversePostorder(
		enterRight,
		enterLeft,
		trueNodeConditionFunc[T],
		func(n *Node[T]) bool {
			res = n.val
			exists = true
			return false
		},
	)
	return
}

func (t *AVL[T]) Rank(data T) int {
	return t.rank(t.root, data)
}

func (t *AVL[T]) RankNth(rank int) (res T, exists bool) {
	return t.rankNth(t.root, rank)
}

func (t *AVL[T]) Range(f datastructure.ConditionFunc[T]) {
	t.root.inorder(trueNodeConditionFunc[T], trueNodeConditionFunc[T], trueNodeConditionFunc[T], nodeConditionWrap(f))
}

func (t *AVL[T]) RangeS(start T, f datastructure.ConditionFunc[T]) {
	enterLeft := func(root *Node[T]) bool {
		return t.cmp.Compare(root.getValue(), start).GT()
	}
	enterCur := func(root *Node[T]) bool {
		return t.cmp.Compare(root.getValue(), start).GTE()
	}
	t.root.inorder(enterLeft, enterCur, trueNodeConditionFunc[T], nodeConditionWrap[T](f))
}

func (t *AVL[T]) RangeSE(start, end T, f datastructure.ConditionFunc[T]) {
	enterLeft := func(root *Node[T]) bool {
		return t.cmp.Compare(root.getValue(), start).GT()
	}
	enterCur := func(root *Node[T]) bool {
		r1 := t.cmp.Compare(root.getValue(), start)
		r2 := t.cmp.Compare(root.getValue(), end)
		return r1.GTE() && r2.LT()
	}
	enterRight := func(root *Node[T]) bool {
		return t.cmp.Compare(root.getValue(), end).LT()
	}
	t.root.inorder(enterLeft, enterCur, enterRight, nodeConditionWrap[T](f))
}

func (t *AVL[T]) RangeE(end T, f datastructure.ConditionFunc[T]) {
	enter := func(root *Node[T]) bool {
		return t.cmp.Compare(root.getValue(), end).LT()
	}
	t.root.inorder(trueNodeConditionFunc[T], enter, enter, nodeConditionWrap[T](f))
}

// Private method

func (t *AVL[T]) newNode(data T) *Node[T] {
	node := t.alloc.Allocate()
	node.setVal(data, t.countableCheck)
	node.height = 1
	node.l = nil
	node.r = nil
	node.pushUp()
	return node
}

func (t *AVL[T]) fixBalance(root *Node[T]) *Node[T] {
	if root.getFactor() < -1 {
		if root.l.getFactor() < 0 {
			// LL -> balance
			root = root.rightRotate()
		} else {
			// LR -> LL -> balance
			root.l = root.l.leftRotate()
			root = root.rightRotate()
		}
	} else if root.getFactor() > 1 {
		if root.r.getFactor() > 0 {
			// RR -> balance
			root = root.leftRotate()
		} else {
			// RL -> RR -> balance
			root.r = root.r.rightRotate()
			root = root.leftRotate()
		}
	}
	return root
}

func (t *AVL[T]) insert(root *Node[T], data T, f nodeVisitFunc[T]) *Node[T] {
	if root == nil {
		return t.newNode(data)
	}
	r := t.cmp.Compare(root.val, data)
	switch r {
	case compare.EQ:
		if f != nil {
			f(root)
		}
	case compare.GT:
		root.l = t.insert(root.l, data, f)
	case compare.LT:
		root.r = t.insert(root.r, data, f)
	default:
		panic("impossible")
	}
	root.pushUp()
	root = t.fixBalance(root)
	return root
}

func (t *AVL[T]) delete(root *Node[T], data T, f nodeConditionFunc[T]) *Node[T] {
	if root == nil {
		return nil
	}
	r := t.cmp.Compare(root.val, data)
	switch r {
	case compare.EQ:
		if f != nil && !f(root) {
			break
		}
		// make sure f is called only once
		f = trueNodeConditionFunc[T]
		if root.l == nil && root.r == nil {
			return nil
		} else if root.l != nil && root.r == nil {
			root = root.l
			break
		} else if root.l == nil && root.r != nil {
			root = root.r
			break
		} else {
			if root.getFactor() < 0 {
				root = root.rightRotate()
				root.r = t.delete(root.r, data, f)
			} else {
				root = root.leftRotate()
				root.l = t.delete(root.l, data, f)
			}
			break
		}
	case compare.GT:
		root.l = t.delete(root.l, data, f)
	case compare.LT:
		root.r = t.delete(root.r, data, f)
	default:
		panic("impossible")
	}
	root.pushUp()
	root = t.fixBalance(root)
	return root
}

func (t *AVL[T]) rankNth(root *Node[T], rank int) (res T, exists bool) {
	if root == nil {
		return
	}
	if rank <= root.l.getSize() {
		return t.rankNth(root.l, rank)
	} else if rank <= root.l.getSize()+root.getCount() {
		res = root.val
		exists = true
		return
	} else {
		return t.rankNth(root.r, rank-root.l.getSize()-root.getCount())
	}
}

func (t *AVL[T]) rank(root *Node[T], data T) int {
	if root == nil {
		return 1
	}
	result := t.cmp.Compare(root.val, data)
	switch result {
	case compare.EQ:
		return root.l.getSize() + 1
	case compare.LT:
		return root.l.getSize() + root.getCount() + t.rank(root.r, data)
	case compare.GT:
		return t.rank(root.l, data)
	default:
		panic("impossible")
	}
}
