package linkedlist

import "fmt"

type singlyLinkedListWithTail struct {
	head   Node
	tail   Node
	length int
}

// AddBack implements LinkedList
func (s *singlyLinkedListWithTail) AddBack(item int) {
	var node = NewNode(item)
	if s.IsEmpty() {
		s.SetHead(node)
		s.SetTail(node)
	} else {
		s.GetTail().SetNext(node)
		s.SetTail(node)
	}
	s.length++
}

// AddFront implements LinkedList
func (s *singlyLinkedListWithTail) AddFront(item int) {
	var node = NewNode(item)
	if s.IsEmpty() {
		s.SetHead(node)
		s.SetTail(node)
	} else {
		node.SetNext(s.GetHead())
		s.SetHead(node)
	}
	s.length++
}

// Back implements LinkedList
func (s *singlyLinkedListWithTail) Back() int {
	if s.GetTail() == nil {
		return SENTINEL_VALUE
	}
	return s.GetTail().GetValue()
}

// Delete implements LinkedList
func (s *singlyLinkedListWithTail) Delete(index int) {
	if s.IsEmpty() || index >= s.Length() {
		return
	}
	if index == 0 {
		s.PopFront()
		return
	}
	if index == s.Length()-1 {
		s.PopBack()
		return
	}

	count := 0
	current := s.GetHead()

	for count < index && current != nil {
		current = current.GetNext()
		count++
	}

	if current.GetNext() != nil {
		next := current.GetNext()
		current.SetValue(next.GetValue())
		current.SetNext(next.GetNext())
		if next == s.GetTail() {
			s.SetTail(current)
		}

	}

}

// Elements implements LinkedList
func (s *singlyLinkedListWithTail) Elements() []int {
	current := s.GetHead()
	elements := make([]int, 0)
	for current != nil {
		elements = append(elements, current.GetValue())
		current = current.GetNext()
	}
	return elements
}

// Front implements LinkedList
func (s *singlyLinkedListWithTail) Front() int {
	if s.GetHead() == nil {
		return SENTINEL_VALUE
	}
	return s.GetHead().GetValue()
}

// GetHead implements LinkedList
func (s *singlyLinkedListWithTail) GetHead() Node {
	return s.head
}

// GetTail implements LinkedList
func (s *singlyLinkedListWithTail) GetTail() Node {
	return s.tail
}

// IsEmpty implements LinkedList
func (s *singlyLinkedListWithTail) IsEmpty() bool {
	return s.length == 0
}

// Length implements LinkedList
func (s *singlyLinkedListWithTail) Length() int {
	return s.length
}

// PopBack implements LinkedList
func (s *singlyLinkedListWithTail) PopBack() int {
	if s.length == 0 {
		return SENTINEL_VALUE
	}
	if s.length == 1 {
		var value int = s.GetHead().GetValue()
		s.SetHead(nil)
		s.SetTail(nil)
		s.length--
		return value
	}

	dummy := NewNode(SENTINEL_VALUE)
	dummy.SetNext(s.GetHead())
	current := s.GetHead()
	for current != nil && current.GetNext() != nil {
		current = current.GetNext()
		dummy = dummy.GetNext()
	}
	dummy.SetNext(nil)
	s.SetTail(dummy)
	s.length--
	return current.GetValue()

}

// PopFront implements LinkedList
func (s *singlyLinkedListWithTail) PopFront() int {
	if s.length == 0 {
		return SENTINEL_VALUE
	}
	if s.length == 1 {
		var value int = s.GetHead().GetValue()
		s.SetHead(nil)
		s.SetTail(nil)
		s.length--
		return value
	}

	next := s.GetHead().GetNext()
	value := s.GetHead().GetValue()
	s.SetHead(next)
	s.length--
	return value
}

// Print implements LinkedList
func (s *singlyLinkedListWithTail) Print() {
	elements := s.Elements()
	n := len(elements)
	for i := 0; i < n; i++ {
		fmt.Print(elements[i])
		if i != n-1 {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

// Remove implements LinkedList
func (s *singlyLinkedListWithTail) Remove(item int) {
	index := -1
	current := s.GetHead()

	for current != nil && item != current.GetValue() {
		current = current.GetNext()
	}

	if index != -1 {
		s.Delete(index)
	}
}

// Reverse implements LinkedList
func (s *singlyLinkedListWithTail) Reverse() {
	var prev Node = nil
	current := s.GetHead()
	tail := current
	for current != nil {
		next := current.GetNext()
		current.SetNext(prev)
		prev = current
		current = next
	}
	s.SetTail(tail)
	s.SetHead(prev)

}

// SetHead implements LinkedList
func (s *singlyLinkedListWithTail) SetHead(node Node) {
	s.head = node
}

// SetTail implements LinkedList
func (s *singlyLinkedListWithTail) SetTail(node Node) {
	s.tail = node
}

func NewSinglyLinkedListWithTailPointer() LinkedList {
	return &singlyLinkedListWithTail{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}
