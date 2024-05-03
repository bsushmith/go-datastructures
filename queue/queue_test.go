package queue

import "testing"

func TestQueue_IsEmpty(t *testing.T) {
	t.Run("empty int queue", func(t *testing.T) {
		q := NewQueue[int]()
		if !q.IsEmpty() {
			t.Errorf("Expected queue to be empty")
		}
	})

	t.Run("empty string queue", func(t *testing.T) {
		q := NewQueue[string]()
		if !q.IsEmpty() {
			t.Errorf("Expected queue to be empty")
		}
	})
}

func TestQueue_Enqueue(t *testing.T) {
	t.Run("enqueue int", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1)
		if q.Len() != 1 {
			t.Errorf("Expected queue length to be 1, got %d", q.Len())
		}
	})

	t.Run("enqueue string", func(t *testing.T) {
		q := NewQueue[string]()
		q.Enqueue("abc")
		if q.Len() != 1 {
			t.Errorf("Expected queue length to be 1, got %d", q.Len())
		}
	})
}

func TestQueue_Dequeue(t *testing.T) {
	t.Run("non-empty queue", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1)
		val, ok := q.Dequeue()
		if !ok || val != 1 {
			t.Errorf("Expected to dequeue 1 from queue, got %d", val)
		}
		if q.Len() != 0 {
			t.Errorf("Expected queue length to be 0, got %d", q.Len())
		}
	})

	t.Run("empty queue", func(t *testing.T) {
		q := NewQueue[int]()
		val, ok := q.Dequeue()
		if ok || val != 0 {
			t.Errorf("Expected to dequeue nothing from queue, got %d", val)
		}
		if q.Len() != 0 {
			t.Errorf("Expected queue length to be 0, got %d", q.Len())
		}
	})
}

func TestQueue_Peek(t *testing.T) {
	t.Run("non-empty queue", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1)
		val, ok := q.Peek()
		if !ok || val != 1 {
			t.Errorf("Expected to peek 1 from queue, got %d", val)
		}
		if q.Len() != 1 {
			t.Errorf("Expected queue length to be 1, got %d", q.Len())
		}
	})

	t.Run("empty queue", func(t *testing.T) {
		q := NewQueue[int]()
		val, ok := q.Peek()
		if ok || val != 0 {
			t.Errorf("Expected to peek nothing from queue, got %d", val)
		}
		if q.Len() != 0 {
			t.Errorf("Expected queue length to be 0, got %d", q.Len())
		}
	})
}
