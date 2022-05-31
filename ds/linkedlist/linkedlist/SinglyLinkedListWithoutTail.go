package linkedlist

import "fmt"

type singlyLinkedlistwithoutTailPointer struct {
	head   Node
	length int
}

// IsEmpty implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) IsEmpty() bool {
	return s.length == 0
}

// AddBack implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) AddBack(item int) {
	node := NewNode(item)
	if s.IsEmpty() {
		s.SetHead(node)
	} else {
		current := s.GetHead()
		for current.GetNext() != nil {
			current = current.GetNext()
		}
		current.SetNext(node)
	}
	s.length++
}

// AddFront implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) AddFront(item int) {
	node := NewNode(item)
	node.SetNext(s.GetHead())
	s.SetHead(node)
	s.length++
}

// Back implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) Back() int {
	panic("unimplemented")
}

// Delete implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) Delete(index int) {
	if s.IsEmpty() {
		return
	}
	if index == 0 {
		s.PopFront()
		return
	}
	if index == s.length-1 {
		s.PopBack()
		return
	}
	current := s.GetHead()
	count := 0
	for count < index && current != nil {
		count++
		current = current.GetNext()
	}

	if current.GetNext() != nil {
		current.SetValue(current.GetNext().GetValue())
		current.SetNext(current.GetNext().GetNext())
		s.length--
	}

}

// Elements implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) Elements() []int {
	elements := make([]int, 0)
	current := s.head

	for current != nil {
		elements = append(elements, current.GetValue())
		current = current.GetNext()
	}
	return elements
}

// Front implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) Front() int {
	head := s.GetHead()
	if head != nil {
		return head.GetValue()
	}
	return SENTINEL_VALUE
}

// GetHead implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) GetHead() Node {
	return s.head
}

// GetTail implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) GetTail() Node {
	panic("unimplemented")
}

// Length implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) Length() int {
	return s.length
}

// PopBack implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) PopBack() int {
	if s.IsEmpty() {
		return SENTINEL_VALUE
	}
	dummy := NewNode(SENTINEL_VALUE)
	dummy.SetNext(s.head)
	current := s.head
	for current.GetNext() != nil {
		current = current.GetNext()
		dummy = dummy.GetNext()
	}
	value := current.GetValue()
	dummy.SetNext(current.GetNext())
	s.length--
	return value

}

// PopFront implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) PopFront() int {
	if s.IsEmpty() {
		return SENTINEL_VALUE
	}

	head := s.GetHead()
	next := head.GetNext()
	s.SetHead(next)
	s.length--
	return head.GetValue()
}

// Print implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) Print() {
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
func (s *singlyLinkedlistwithoutTailPointer) Remove(item int) {
	index := 0
	current := s.head
	for current != nil {
		if current.GetValue() == item {
			break
		}
		index++
		current = current.GetNext()
	}

	if index != -1 {
		s.Delete(index)
	}
}

// Reverse implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) Reverse() {
	var prev Node = nil

	current := s.head
	for current != nil {
		next := current.GetNext()
		current.SetNext(prev)
		prev = current
		current = next
	}

	s.head = prev
}

// SetHead implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) SetHead(node Node) {
	s.head = node
}

// SetTail implements LinkedList
func (s *singlyLinkedlistwithoutTailPointer) SetTail(node Node) {
	panic("unimplemented")
}

func NewSinglyLinkedListWithoutTailPointer() LinkedList {
	return &singlyLinkedlistwithoutTailPointer{
		head:   nil,
		length: 0,
	}
}
