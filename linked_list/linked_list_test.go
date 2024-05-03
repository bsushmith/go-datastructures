package linked_list

import "testing"

func TestLinkedList_Insert(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)
	if ll.Len() != 5 {
		t.Errorf("Expected length 5, got %d", ll.Len())
	}
}

func TestLinkedList_Remove(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Remove(1)
	if ll.Len() != 0 {
		t.Errorf("Expected length 0, got %d", ll.Len())
	}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)
	ll.Remove(3)
	if ll.Len() != 4 {
		t.Errorf("Expected length 4, got %d", ll.Len())
	}
	ll.Remove(1)
	if ll.Len() != 3 {
		t.Errorf("Expected length 3, got %d", ll.Len())
	}
	ll.Remove(5)
	if ll.Len() != 2 {
		t.Errorf("Expected length 2, got %d", ll.Len())
	}
}
