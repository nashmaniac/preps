package linkedlist

type Element interface {
	string | int | int32 | int64 | float32 | float64 | interface{}
}

type LinkedList interface {
	Size() int
	Empty() bool
	ValueAt(index int) Element
	PushFront(value Element)
	PopFront() Element
	PushBack(value Element)
	PopBack() Element
	Front() Element
	Back() Element
	Insert(index int, value Element)
	Erase(index int)
	ValueFromEnd(n int) Element
	Reverse()
	RemoveValue(value Element)
}
