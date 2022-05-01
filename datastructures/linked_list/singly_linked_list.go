package linkedlist

type singleLinkedListNode struct {
	value Element
	next  *singleLinkedListNode
}

type singlyLinkedList struct {
	Head *singleLinkedListNode
	Tail *singleLinkedListNode
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

	if s.Size() == 1 {
		node := s.Head
		s.Head = nil
		s.Tail = nil
		s.size--
		return node.value
	}

	dummy := &singleLinkedListNode{
		value: -1,
	}
	dummy.next = s.Head
	current := s.Head

	for current.next != nil {
		current = current.next
		dummy = dummy.next
	}
	s.Tail = dummy
	s.size--
	return current.value
}

// PopFront implements LinkedList
func (s *singlyLinkedList) PopFront() Element {
	if s.Empty() {
		return nil
	}

	node := s.Head
	s.Head = s.Head.next
	s.size--
	return node.value
}

// PushBack implements LinkedList
func (s *singlyLinkedList) PushBack(value Element) {
	node := &singleLinkedListNode{
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
	node := &singleLinkedListNode{
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

// RemoveValue implements LinkedList
func (s *singlyLinkedList) RemoveValue(value Element) {
	panic("unimplemented")
}

// Reverse implements LinkedList
func (s *singlyLinkedList) Reverse() {
	panic("unimplemented")
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
