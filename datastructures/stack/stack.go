package stack

import linkedlist "github.com/nashmaniac/golang-coding/datastructures/linked_list"

type Stack interface {
	Push(value linkedlist.Element)
	Pop() linkedlist.Element
	Peek() linkedlist.Element
	IsEmpty() bool
}

type stack struct {
	List linkedlist.LinkedList
}

// IsEmpty implements Stack
func (s *stack) IsEmpty() bool {
	return s.List.Empty()
}

// Peek implements Stack
func (s *stack) Peek() linkedlist.Element {
	return s.List.Front()
}

// Pop implements Stack
func (s *stack) Pop() linkedlist.Element {
	return s.List.PopFront()
}

// Push implements Stack
func (s *stack) Push(value linkedlist.Element) {
	s.List.PushFront(value)
}

func NewStack() Stack {
	return &stack{
		List: linkedlist.NewSinglyLinkedList[int](),
	}
}
