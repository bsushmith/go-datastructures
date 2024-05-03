package tree

import (
	"cmp"
	"fmt"
)

type Node[T cmp.Ordered] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

type BinarySearchTree[T cmp.Ordered] struct {
	root *Node[T]
	size int
}

func NewBinaryTree[T cmp.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{root: nil, size: 0}
}

func (t *BinarySearchTree[T]) Insert(elem T) {
	t.insertN(t.root, elem)
	t.size++
}

func (t *BinarySearchTree[T]) Search(elem T) bool {
	isPresent := t.searchN(t.root, elem)
	return isPresent
}

func (t *BinarySearchTree[T]) Size() int {
	return t.size
}

func (t *BinarySearchTree[T]) PreOrderTraversal() {
	t.preOrderTraversalN(t.root)
}

func (t *BinarySearchTree[T]) PostOrderTraversal() {
	t.postOrderTraversalN(t.root)
}

func (t *BinarySearchTree[T]) InOrderTraversal() {
	t.inOrderTraversalN(t.root)
}

func (t *BinarySearchTree[T]) insertN(node *Node[T], elem T) *Node[T] {
	if t.root == nil {
		t.root = &Node[T]{Val: elem}
		return t.root
	}
	if node == nil {
		return &Node[T]{Val: elem}
	}
	if node.Val >= elem {
		node.Left = t.insertN(node.Left, elem)
	}
	if node.Val < elem {
		node.Right = t.insertN(node.Right, elem)
	}
	return node
}

func (t *BinarySearchTree[T]) searchN(node *Node[T], elem T) bool {
	if node == nil {
		return false
	}
	if node.Val == elem {
		return true
	}
	if node.Val > elem {
		return t.searchN(node.Left, elem)
	}
	if node.Val < elem {
		return t.searchN(node.Right, elem)
	}
	return false
}

func (t *BinarySearchTree[T]) preOrderTraversalN(node *Node[T]) {
	if node == nil {
		return
	}
	fmt.Printf("%v ", node.Val)
	t.preOrderTraversalN(node.Left)
	t.preOrderTraversalN(node.Right)
}

func (t *BinarySearchTree[T]) postOrderTraversalN(node *Node[T]) {
	if node == nil {
		return
	}
	t.postOrderTraversalN(node.Left)
	t.postOrderTraversalN(node.Right)
	fmt.Printf("%v ", node.Val)
}

func (t *BinarySearchTree[T]) inOrderTraversalN(node *Node[T]) {
	if node == nil {
		return
	}
	t.inOrderTraversalN(node.Left)
	fmt.Printf("%v ", node.Val)
	t.inOrderTraversalN(node.Right)
}
