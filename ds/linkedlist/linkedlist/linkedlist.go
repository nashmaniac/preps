package linkedlist

import "math"

var (
	SENTINEL_VALUE = math.MinInt16
)

type LinkedList interface {
	IsEmpty() bool
	Length() int
	AddFront(item int)
	AddBack(item int)
	PopFront() int
	PopBack() int
	Front() int
	Back() int
	Delete(index int)
	Remove(item int)
	Reverse()
	Print()
	Elements() []int
	GetHead() Node
	GetTail() Node
	SetHead(node Node)
	SetTail(node Node)
}

// Time complexity

/***

	Length() int - O(1) as we will keep value in length variable
	AddFront()
		1. SinglyLinkedList
			With Tail Pointer - O(1)
			Without Tail Pointer - O(1)
		2. Doubly LinkedList
			With Tail Pointer - O(1)
			Without Tail Pointer - O(1)
	AddBack()
		1. SinglyLinkedList
			With Tail Pointer - O(1)
			Without Tail Pointer - O(n)
		2. Doubly LinkedList
			With Tail Pointer - O(1)
			Without Tail Pointer - O(n)
	PopFront() int
		1. SinglyLinkedList
			With Tail Pointer - O(1)
			Without Tail Pointer - O(1)
		2. Doubly LinkedList
			With Tail Pointer - O(1)
			Without Tail Pointer - O(1)
	PopBack() int
		1. SinglyLinkedList
			With Tail Pointer - O(n)
			Without Tail Pointer - O(n)
		2. Doubly LinkedList
			With Tail Pointer - O(1)
			Without Tail Pointer - O(n)
	Front() int
	Back() int
	Delete(index int)
	Remove(item int)
	Reverse()
	Print()
	Elements() []int

**/
