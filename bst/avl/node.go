package avl

import (
	"github.com/Sora233/datastructure"
	"github.com/Sora233/datastructure/bst"
)

// Node is the node of AVL tree
type Node[T any] struct {
	fa, l, r *Node[T]
	val      T
	countval bst.Countable
	size     int
	height   int
}

func (node *Node[T]) getFa() *Node[T] {
	if node == nil {
		return nil
	}
	return node.fa
}

func (node *Node[T]) setFa(fa *Node[T]) {
	if node == nil || node.fa == fa {
		return
	}
	node.fa = fa
}

func (node *Node[T]) setVal(data T, cc bool) {
	node.val = data
	if !cc {
		return
	}
	if c, ok := any(data).(bst.Countable); ok {
		node.countval = c
	}
}

// pushUp recalculate the size of subtree
func (node *Node[T]) pushUp() {
	if node == nil {
		return
	}
	node.size = node.getCount() + node.l.getSize() + node.r.getSize()
	node.l.setFa(node)
	node.r.setFa(node)
	if node.l.getHeight() > node.r.getHeight() {
		node.height = 1 + node.l.getHeight()
	} else {
		node.height = 1 + node.r.getHeight()
	}
}

func (node *Node[T]) getFactor() int {
	if node == nil {
		return 0
	}
	return node.r.getHeight() - node.l.getHeight()
}

func (node *Node[T]) getHeight() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node[T]) getSize() int {
	if node == nil {
		return 0
	}
	return node.size
}

func (node *Node[T]) getValue() (res T) {
	if node != nil {
		res = node.val
	}
	return
}

func (node *Node[T]) getCount() int {
	if node == nil {
		return 0
	}
	if node.countval != nil {
		return node.countval.Count()
	}
	return 1
}

// leftRotate operator a left-rotate
// The right-child becomes the new root
// return the new root
func (node *Node[T]) leftRotate() *Node[T] {
	if node == nil {
		return nil
	}
	fa := node.fa
	rNode := node.r
	node.r = rNode.l
	rNode.l.setFa(node)
	rNode.l = node
	node.setFa(rNode)
	rNode.setFa(fa)

	node.pushUp()
	rNode.pushUp()

	return rNode
}

// rightRotate operator a right-rotate
// The left-child becomes the new root
// return the new root
func (node *Node[T]) rightRotate() *Node[T] {
	if node == nil {
		return nil
	}
	fa := node.fa
	lNode := node.l
	node.l = lNode.r
	lNode.r.setFa(node)
	lNode.r = node
	node.setFa(lNode)
	lNode.setFa(fa)

	node.pushUp()
	lNode.pushUp()

	return lNode
}

type nodeVisitFunc[T any] func(*Node[T])

func nodeVisitWrap[T any](f datastructure.VisitFunc[T]) nodeVisitFunc[T] {
	return func(node *Node[T]) {
		f(node.getValue())
	}
}

type nodeConditionFunc[T any] func(*Node[T]) bool

func nodeConditionWrap[T any](f datastructure.ConditionFunc[T]) nodeConditionFunc[T] {
	return func(node *Node[T]) bool {
		return f(node.getValue())
	}
}

func trueNodeConditionFunc[T any](*Node[T]) bool {
	return true
}

// inorder Inorder traversal the tree
// left first, then current, last right
func (node *Node[T]) inorder(enterLeft, enterCur, enterRight, f nodeConditionFunc[T]) bool {
	if node == nil {
		return true
	}
	if enterLeft != nil && enterLeft(node) {
		if !node.l.inorder(enterLeft, enterCur, enterRight, f) {
			return false
		}
	}
	if enterCur != nil && enterCur(node) {
		if !f(node) {
			return false
		}
	}
	if enterRight != nil && enterRight(node) {
		if !node.r.inorder(enterLeft, enterCur, enterRight, f) {
			return false
		}
	}
	return true
}

// reverseInorder Inorder traversal the tree in reverse order
// right first, then current, last left
func (node *Node[T]) reverseInorder(enterRight, enterCur, enterLeft, f nodeConditionFunc[T]) bool {
	if node == nil {
		return true
	}
	if enterRight != nil && enterRight(node) {
		if !node.r.reverseInorder(enterRight, enterCur, enterLeft, f) {
			return false
		}
	}
	if enterCur != nil && enterCur(node) {
		if !f(node) {
			return false
		}
	}
	if enterLeft != nil && enterLeft(node) {
		if !node.l.reverseInorder(enterRight, enterCur, enterLeft, f) {
			return false
		}
	}
	return true
}

// postorder Postorder traversal the tree
// left first, then right, last current
func (node *Node[T]) postorder(enterLeft, enterRight, enterCur, f nodeConditionFunc[T]) bool {
	if node == nil {
		return true
	}
	if enterLeft != nil && enterLeft(node) {
		if !node.l.postorder(enterLeft, enterRight, enterCur, f) {
			return false
		}
	}
	if enterRight != nil && enterRight(node) {
		if !node.r.postorder(enterLeft, enterRight, enterCur, f) {
			return false
		}
	}
	if enterCur != nil && enterCur(node) {
		if !f(node) {
			return false
		}
	}
	return true
}

// reversePostorder Postorder traversal the tree in reverse order
// right first, then left, last current
func (node *Node[T]) reversePostorder(enterRight, enterLeft, enterCur, f nodeConditionFunc[T]) bool {
	if node == nil {
		return true
	}
	if enterRight != nil && enterRight(node) {
		if !node.r.reversePostorder(enterRight, enterLeft, enterCur, f) {
			return false
		}
	}
	if enterLeft != nil && enterLeft(node) {
		if !node.l.reversePostorder(enterRight, enterLeft, enterCur, f) {
			return false
		}
	}
	if enterCur != nil && enterCur(node) {
		if !f(node) {
			return false
		}
	}
	return true
}

// preorder Preorder traversal the tree
// current first, then left, last right
func (node *Node[T]) preorder(enterCur, enterLeft, enterRight, f nodeConditionFunc[T]) bool {
	if node == nil {
		return true
	}
	if enterCur != nil && enterCur(node) {
		if !f(node) {
			return false
		}
	}
	if enterLeft != nil && enterLeft(node) {
		if !node.l.preorder(enterCur, enterLeft, enterRight, f) {
			return false
		}
	}
	if enterRight != nil && enterRight(node) {
		if !node.r.preorder(enterCur, enterLeft, enterRight, f) {
			return false
		}
	}
	return true
}

// reversePreorder Preorder traversal the tree in reverse order
// current first, then right, last left
func (node *Node[T]) reversePreorder(enterCur, enterRight, enterLeft, f nodeConditionFunc[T]) bool {
	if node == nil {
		return true
	}
	if enterCur != nil && enterCur(node) {
		if !f(node) {
			return false
		}
	}
	if enterRight != nil && enterRight(node) {
		if !node.r.reversePreorder(enterCur, enterRight, enterLeft, f) {
			return false
		}
	}
	if enterLeft != nil && enterLeft(node) {
		if !node.l.reversePreorder(enterCur, enterRight, enterLeft, f) {
			return false
		}
	}
	return true
}
