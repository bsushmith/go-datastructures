package stack

import "testing"

func TestStack_Push(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	if stack.Len() != 1 {
		t.Errorf("Expected stack length to be 1, got %d", stack.Len())
	}
}

func TestStack_Pop(t *testing.T) {

	t.Run("non-empty stack", func(t *testing.T) {
		stack := NewStack[string]()
		stack.Push("abc")
		val, ok := stack.Pop()
		if !ok || val != "abc" {
			t.Errorf("Expected to pop 'abc' from stack, got %s", val)
		}
		if stack.Len() != 0 {
			t.Errorf("Expected stack length to be 0, got %d", stack.Len())
		}
	})

	t.Run("empty stack", func(t *testing.T) {
		stack := NewStack[string]()
		val, ok := stack.Pop()
		if ok || val != "" {
			t.Errorf("Expected to pop nothing from stack, got %s", val)
		}
		if stack.Len() != 0 {
			t.Errorf("Expected stack length to be non-zero, got %d", stack.Len())
		}
	})

}

func TestStack_Peek(t *testing.T) {

	t.Run("non-empty stack", func(t *testing.T) {
		stack := NewStack[string]()
		stack.Push("xyz")
		val, ok := stack.Peek()
		if !ok || val != "xyz" {
			t.Errorf("Expected to peek 'xyz' from stack, got %s", val)
		}
		if stack.Len() != 1 {
			t.Errorf("Expected stack length to be 1, got %d", stack.Len())
		}
	})
	t.Run("empty stack", func(t *testing.T) {
		stack := NewStack[string]()
		val, ok := stack.Peek()
		if ok || val != "" {
			t.Errorf("Expected to peek nothing from stack, got %s", val)
		}
		if stack.Len() != 0 {
			t.Errorf("Expected stack length to be 0, got %d", stack.Len())
		}
	})

}

func TestStack_IsEmpty(t *testing.T) {

	t.Run("Empty int stack", func(t *testing.T) {
		stack := NewStack[int]()
		if !stack.IsEmpty() {
			t.Errorf("Expected stack to be empty")
		}
	})

	t.Run("Non-empty stack", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)
		if stack.IsEmpty() {
			t.Errorf("Expected stack to be non-empty")
		}
	})

	t.Run("Empty string stack", func(t *testing.T) {
		stack := NewStack[string]()
		if !stack.IsEmpty() {
			t.Errorf("Expected stack to be empty")
		}
	})

}
