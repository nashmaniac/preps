package stack

import "fmt"

type Stack interface {
	Top() Element
	Push(e Element)
	Pop() Element
	IsEmpty() bool
	Print()
}

type stack struct {
	head Node
	size int
}

// Print implements Stack
func (s *stack) Print() {
	current := s.head
	for current != nil {
		fmt.Println(current.GetValue())
		fmt.Println("----")
		current = current.GetNext()
	}
}

// IsEmpty implements Stack
func (s *stack) IsEmpty() bool {
	return s.size == 0
}

// Pop implements Stack
func (s *stack) Pop() Element {
	if s.IsEmpty() {
		return nil
	}

	head := s.head
	next := head.GetNext()
	if next != nil {
		next.SetPrev(nil)
	}
	s.head = next
	s.size--
	return head.GetValue()
}

// Push implements Stack
func (s *stack) Push(e Element) {
	node := NewNode(e)
	if s.IsEmpty() {
		s.head = node
		s.size++
		return
	}
	node.SetNext(s.head)
	s.head.SetPrev(node)
	s.head = node
	s.size++
}

// Top implements Stack
func (s *stack) Top() Element {
	if s.IsEmpty() {
		return nil
	}
	return s.head.GetValue()
}

func NewStack[e Element]() Stack {
	return &stack{
		head: nil,
		size: 0,
	}
}
