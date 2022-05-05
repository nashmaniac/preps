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

// Interate implements LinkedList
func (s *singlyLinkedList) Iterate() []Element {
	nodes := make([]Element, 0)
	current := s.Head
	for current != nil {
		nodes = append(nodes, current.value)
		current = current.next
	}
	return nodes
}

// RemoveValue implements LinkedList
func (s *singlyLinkedList) RemoveValue(value Element) {
	if s.Empty() {
		return
	}

	if s.Head.value == value {
		s.PopFront()
		return
	}

	current := s.Head
	next := s.Head.next
	for next != nil && next.value != value {
		current = next
		next = next.next
	}

	if next == nil {
		return
	}
	current.next = next.next
	s.Tail = current
	s.size--
	if s.Empty() {
		s.Head = nil
		s.Tail = nil
	}

}

// Reverse implements LinkedList
func (s *singlyLinkedList) Reverse() {
	current := s.Head
	tail := s.Head
	var prev *singlyLinkedNode = nil

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	s.Head = prev
	s.Tail = tail
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
	if index == 0 {
		s.PopFront()
		return
	}
	if index >= s.size {
		s.PopBack()
		return
	}
	current := s.Head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	} else {
		current.next = nil
	}
	s.size--
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
	if index == 0 {
		s.PushFront(value)
		return
	}
	if index >= s.size {
		s.PushBack(value)
		return
	}

	current := s.Head

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	node := &singlyLinkedNode{
		value: value,
		next:  nil,
	}

	next := current.next
	current.next = node
	node.next = next
	s.size++
}

// PopBack implements LinkedList
func (s *singlyLinkedList) PopBack() Element {
	if s.Empty() {
		return nil
	}
	dummy := &singlyLinkedNode{
		value: -1,
		next:  s.Head,
	}
	current := s.Head
	for current.next != nil {
		dummy = dummy.next
		current = current.next
	}
	s.size--
	if s.Empty() {
		s.Head = nil
		s.Tail = nil
	} else {
		s.Tail = dummy
		s.Tail.next = current.next
	}

	return current.value
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
	if index >= s.size {
		return nil
	}
	current := s.Head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value
}

// ValueFromEnd implements LinkedList
func (s *singlyLinkedList) ValueFromEnd(n int) Element {
	// wrapping up strategy
	n = n % s.size
	return s.ValueAt(s.size - n)
}

func NewSinglyLinkedList[e Element]() LinkedList {
	return &singlyLinkedList{
		Head: nil,
		Tail: nil,
		size: 0,
	}
}
