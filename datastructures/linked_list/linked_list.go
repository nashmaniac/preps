package linkedlist

type Element interface {
	string | int | int32 | int64 | float32 | float64 | interface{}
}

type LinkedList interface {
	Size() int
	Empty() bool
	PushFront(value Element)
	PushBack(value Element)
	Front() Element
	Back() Element
	PopFront() Element
	PopBack() Element
	ValueAt(index int) Element
	Insert(index int, value Element)
	Erase(index int)
	ValueFromEnd(n int) Element
	// Reverse()
	// RemoveValue(value Element)
}
