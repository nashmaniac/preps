package queue

import "fmt"

type Queue interface {
	Offer(item Element)
	OfferItem(key int, item Element)
	Poll() Element
	IsEmpty() bool
	Print()
}
type queue struct {
	head Node
	tail Node
	size int
}

// OfferItem implements Queue
func (q *queue) OfferItem(key int, item Element) {
	panic("unimplemented")
}

// IsEmpty implements Queue
func (q *queue) IsEmpty() bool {
	return q.size == 0
}

// Offer implements Queue
func (q *queue) Offer(item Element) {
	node := NewNode(item)
	if q.IsEmpty() {
		q.head = node
		q.tail = node
	} else {
		q.tail.SetNext(node)
		node.SetPrev(q.tail)
		q.tail = node
	}
	q.size++
}

// Poll implements Queue
func (q *queue) Poll() Element {
	if q.IsEmpty() {
		return nil
	}

	head := q.head
	next := q.head.GetNext()
	if next != nil {
		next.SetPrev(nil)
	}
	q.head = next
	q.size--
	return head.GetValue()
}

// Print implements Queue
func (q *queue) Print() {
	current := q.head
	for current != nil {
		fmt.Printf("%d --> ", current.GetValue())
		current = current.GetNext()
	}
	fmt.Printf("nil\n")
}

func NewQueue[E Element]() Queue {
	return &queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}
