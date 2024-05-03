package queue

type Queue[T any] []T

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) Enqueue(element T) {
	*q = append(*q, element)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var t T
	if q.Len() > 0 {
		res := (*q)[0]
		*q = (*q)[1:]
		return res, true
	}
	return t, false
}

func (q *Queue[T]) Peek() (T, bool) {
	var t T
	if q.Len() > 0 {
		return (*q)[0], true
	}
	return t, false
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}
