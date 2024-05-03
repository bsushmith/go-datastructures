package linked_list

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, size: 0}
}

func (l *LinkedList[T]) Insert(elem T) {
	node := &Node[T]{Val: elem}
	if l.head == nil {
		l.head = node
	} else {
		ptr := l.head
		for ptr.Next != nil {
			ptr = ptr.Next
		}
		ptr.Next = node
	}
	l.size++
}

func (l *LinkedList[T]) Remove(elem T) {
	if l.head == nil {
		return
	}
	if l.head.Val == elem {
		l.head = l.head.Next
		l.size--
		return
	}
	ptr := l.head
	for ptr.Next != nil {
		if ptr.Next.Val == elem {
			ptr.Next = ptr.Next.Next
			break
		}
		ptr = ptr.Next
	}
	l.size--
}

func (l *LinkedList[T]) Len() int {
	return l.size
}
