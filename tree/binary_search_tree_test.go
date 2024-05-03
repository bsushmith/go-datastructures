package tree

import (
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Insert(10)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(12)
	bt.Insert(17)

	if bt.root.Val != 10 {
		t.Errorf("Expected root to be 10, got %d", bt.root.Val)
	}

	if bt.Size() != 7 {
		t.Errorf("Expected size to be 7, got %d", bt.Size())
	}

	if bt.root.Left.Val != 5 {
		t.Errorf("Expected root.Left to be 5, got %d", bt.root.Left.Val)
	}

	if bt.root.Right.Val != 15 {
		t.Errorf("Expected root.Right to be 15, got %d", bt.root.Right.Val)
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Insert(10)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(12)
	bt.Insert(17)

	if !bt.Search(10) {
		t.Errorf("Expected 10 to be found")
	}

	if !bt.Search(5) {
		t.Errorf("Expected 5 to be found")
	}

	if !bt.Search(15) {
		t.Errorf("Expected 15 to be found")
	}

	if !bt.Search(3) {
		t.Errorf("Expected 3 to be found")
	}

	if !bt.Search(7) {
		t.Errorf("Expected 7 to be found")
	}

	if !bt.Search(12) {
		t.Errorf("Expected 12 to be found")
	}

	if !bt.Search(17) {
		t.Errorf("Expected 17 to be found")
	}

	if bt.Search(1) {
		t.Errorf("Expected 1 to not be found")
	}

	if bt.Search(20) {
		t.Errorf("Expected 20 to not be found")
	}
}
