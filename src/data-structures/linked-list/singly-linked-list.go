package linked_list

// Node represents a single node in the singly linked list.
type Node[T comparable] struct {
	value T
	next  *Node[T]
}

// SinglyLinkedList represents a singly linked list data structure.
type SinglyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size uint64
}

// NewSinglyLinkedList creates a new instance of a singly linked list.
func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Size returns the number of elements in the singly linked list.
func (l *SinglyLinkedList[T]) Size() uint64 {
	return l.size
}

// IsEmpty checks if the singly linked list is empty.
func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear removes all elements from the singly linked list.
func (l *SinglyLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// PushBack adds a new element to the end of the singly linked list.
func (l *SinglyLinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		value: value,
		next:  nil,
	}

	if l.IsEmpty() {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}

	l.size++
}

// PushFront adds a new element to the front of the singly linked list.
func (l *SinglyLinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		value: value,
		next:  nil,
	}

	if l.IsEmpty() {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}

	l.size++
}

// Pop removes a specific node from the singly linked list.
func (l *SinglyLinkedList[T]) Pop(node *Node[T]) *Node[T] {
	if l.IsEmpty() || node == nil {
		return nil
	}

	prev := l.head
	for prev.next != node || prev.next != nil {
		prev = prev.next
	}

	found := prev.next

	if found != nil {
		prev.next = found.next
	}

	l.size--
	return found
}

// PopFront removes the first node from the singly linked list.
func (l *SinglyLinkedList[T]) PopFront() *Node[T] {
	if l.IsEmpty() {
		return nil
	}

	found := l.head
	l.head = l.head.next
	l.size--
	return found
}

// PopBack removes the last node from the singly linked list.
func (l *SinglyLinkedList[T]) PopBack() *Node[T] {
	if l.IsEmpty() {
		return nil
	}

	if l.size == 1 {
		found := l.head
		l.head = nil
		l.tail = nil
		l.size--
		return found
	}

	prev := l.head
	for prev.next.next != nil {
		prev = prev.next
	}

	found := prev.next
	prev.next = nil
	l.tail = prev
	l.size--
	return found
}

// Head returns the first node of the singly linked list.
func (l *SinglyLinkedList[T]) Head() *Node[T] {
	return l.head
}

// Tail returns the last node of the singly linked list.
func (l *SinglyLinkedList[T]) Tail() *Node[T] {
	return l.tail
}

// Find searches for a node with the specified value in the singly linked list.
func (l *SinglyLinkedList[T]) Find(value T) *Node[T] {
	found := l.head
	for found != nil && found.value != value {
		found = found.next
	}

	return found
}
