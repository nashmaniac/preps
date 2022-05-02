package linkedlist

type singlyLinkedNode struct {
	value Element
	next  *singlyLinkedNode
}

type singlyLinkedList struct {
	Head *singlyLinkedNode
	Tail *singlyLinkedNode
	size int
}

// Back implements LinkedList
func (s *singlyLinkedList) Back() Element {
	if s.Empty() {
		return nil
	}
	return s.Tail.value
}

// Empty implements LinkedList
func (s *singlyLinkedList) Empty() bool {
	return s.size == 0
}

// Erase implements LinkedList
func (s *singlyLinkedList) Erase(index int) {
	panic("unimplemented")
}

// Front implements LinkedList
func (s *singlyLinkedList) Front() Element {
	if s.Empty() {
		return nil
	}
	return s.Head.value
}

// Insert implements LinkedList
func (s *singlyLinkedList) Insert(index int, value Element) {
	panic("unimplemented")
}

// PopBack implements LinkedList
func (s *singlyLinkedList) PopBack() Element {
	if s.Empty() {
		return nil
	}
	return nil
}

// PopFront implements LinkedList
func (s *singlyLinkedList) PopFront() Element {
	if s.Empty() {
		return nil
	}

	node := s.Head
	s.Head = node.next
	s.size--
	if s.Empty() {
		s.Tail = node.next
	}
	return node.value
}

// PushBack implements LinkedList
func (s *singlyLinkedList) PushBack(value Element) {
	node := &singlyLinkedNode{
		value: value,
		next:  nil,
	}
	if s.Empty() {
		s.Head = node
		s.Tail = node
	} else {
		s.Tail.next = node
		s.Tail = node
	}
	s.size++
}

// PushFront implements LinkedList
func (s *singlyLinkedList) PushFront(value Element) {
	node := &singlyLinkedNode{
		value: value,
		next:  nil,
	}
	if s.Empty() {
		s.Head = node
		s.Tail = node
	} else {
		node.next = s.Head
		s.Head = node
	}
	s.size++
}

// Size implements LinkedList
func (s *singlyLinkedList) Size() int {
	return s.size
}

// ValueAt implements LinkedList
func (s *singlyLinkedList) ValueAt(index int) Element {
	panic("unimplemented")
}

// ValueFromEnd implements LinkedList
func (s *singlyLinkedList) ValueFromEnd(n int) Element {
	panic("unimplemented")
}

func NewSinglyLinkedList[e Element]() LinkedList {
	return &singlyLinkedList{
		Head: nil,
		Tail: nil,
		size: 0,
	}
}
