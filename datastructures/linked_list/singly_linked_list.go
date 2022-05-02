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
func (*singlyLinkedList) Back() Element {
	panic("unimplemented")
}

// Empty implements LinkedList
func (*singlyLinkedList) Empty() bool {
	panic("unimplemented")
}

// Erase implements LinkedList
func (*singlyLinkedList) Erase(index int) {
	panic("unimplemented")
}

// Front implements LinkedList
func (*singlyLinkedList) Front() Element {
	panic("unimplemented")
}

// Insert implements LinkedList
func (*singlyLinkedList) Insert(index int, value Element) {
	panic("unimplemented")
}

// PopBack implements LinkedList
func (*singlyLinkedList) PopBack() Element {
	panic("unimplemented")
}

// PopFront implements LinkedList
func (*singlyLinkedList) PopFront() Element {
	panic("unimplemented")
}

// PushBack implements LinkedList
func (*singlyLinkedList) PushBack(value Element) {
	panic("unimplemented")
}

// PushFront implements LinkedList
func (*singlyLinkedList) PushFront(value Element) {
	panic("unimplemented")
}

// Size implements LinkedList
func (*singlyLinkedList) Size() int {
	panic("unimplemented")
}

// ValueAt implements LinkedList
func (*singlyLinkedList) ValueAt(index int) Element {
	panic("unimplemented")
}

// ValueFromEnd implements LinkedList
func (*singlyLinkedList) ValueFromEnd(n int) Element {
	panic("unimplemented")
}

func NewSinglyLinkedList[e Element]() LinkedList {
	return &singlyLinkedList{
		Head: nil,
		Tail: nil,
		size: 0,
	}
}
