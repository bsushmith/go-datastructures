package stack

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Push(element T) {
	*s = append(*s, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	var t T
	if s.Len() > 0 {
		i := s.Len() - 1
		res := (*s)[i]
		*s = (*s)[:i]
		return res, true
	}
	return t, false
}

func (s *Stack[T]) Peek() (T, bool) {
	var t T
	if s.Len() > 0 {
		return (*s)[s.Len()-1], true
	}
	return t, false
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}
