package linkedlist

import "fmt"

type doublyLinkedList struct {
	head   Node
	tail   Node
	length int
}

// AddBack implements LinkedList
func (d *doublyLinkedList) AddBack(item int) {
	var node Node = NewNode(item)
	if d.IsEmpty() {
		d.SetHead(node)
		d.SetTail(node)
		d.length++
		return
	}

	node.SetPrev(d.GetTail())
	d.GetTail().SetNext(node)
	d.SetTail(node)
	d.length++
}

// AddFront implements LinkedList
func (d *doublyLinkedList) AddFront(item int) {
	var node Node = NewNode(item)
	if d.IsEmpty() {
		d.SetHead(node)
		d.SetTail(node)
		d.length++
		return
	}

	node.SetNext(d.GetHead())
	d.GetHead().SetPrev(node)
	d.SetHead(node)
	d.length++
}

// Back implements LinkedList
func (d *doublyLinkedList) Back() int {
	if d.GetTail() == nil {
		return SENTINEL_VALUE
	}
	return d.GetTail().GetValue()
}

// Delete implements LinkedList
func (d *doublyLinkedList) Delete(index int) {
	if d.IsEmpty() {
		return
	}
	if index == 0 {
		d.PopFront()
		return
	}
	if index == d.length-1 {
		d.PopBack()
		return
	}

	count := 0
	current := d.GetHead()
	for current != nil && count < index {
		current = current.GetNext()
		count++
	}

	if current != nil {
		if current.GetNext() != nil {
			next := current.GetNext()
			current.SetValue(next.GetValue())
			current.SetNext(next.GetNext())
			next.SetPrev(current.GetPrev())
		} else {
			current.GetPrev().SetNext(nil)
			d.SetTail(current.GetPrev())
		}
		d.length--
	}
}

// Elements implements LinkedList
func (d *doublyLinkedList) Elements() []int {
	var elements []int = make([]int, 0)
	current := d.GetHead()
	for current != nil {
		elements = append(elements, current.GetValue())
		current = current.GetNext()
	}
	return elements

}

// Front implements LinkedList
func (d *doublyLinkedList) Front() int {
	if d.GetHead() == nil {
		return SENTINEL_VALUE
	}
	return d.GetHead().GetValue()
}

// GetHead implements LinkedList
func (d *doublyLinkedList) GetHead() Node {
	return d.head
}

// GetTail implements LinkedList
func (d *doublyLinkedList) GetTail() Node {
	return d.tail
}

// IsEmpty implements LinkedList
func (d *doublyLinkedList) IsEmpty() bool {
	return d.length == 0
}

// Length implements LinkedList
func (d *doublyLinkedList) Length() int {
	return d.length
}

// PopBack implements LinkedList
func (d *doublyLinkedList) PopBack() int {
	if d.IsEmpty() {
		return SENTINEL_VALUE
	}

	tail := d.GetTail().GetPrev()
	var oldTail Node = d.GetTail()

	if tail != nil {
		tail.SetNext(nil)
		d.SetTail(tail)
	} else {
		d.SetHead(nil)
		d.SetTail(nil)
	}
	d.length--
	return oldTail.GetValue()
}

// PopFront implements LinkedList
func (d *doublyLinkedList) PopFront() int {
	if d.IsEmpty() {
		return SENTINEL_VALUE
	}

	var oldHead Node = d.GetHead()
	var head = d.GetHead().GetNext()

	if head != nil {
		head.SetPrev(nil)
		d.SetHead(head)
	} else {
		d.SetHead(nil)
		d.SetTail(nil)
	}

	d.length--
	return oldHead.GetValue()
}

// Print implements LinkedList
func (d *doublyLinkedList) Print() {
	elements := d.Elements()
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
func (d *doublyLinkedList) Remove(item int) {
	current := d.GetHead()

	for current != nil && item != current.GetValue() {
		current = current.GetNext()
	}

	if current != nil {
		if current.GetNext() != nil {
			next := current.GetNext()
			current.SetValue(next.GetValue())
			current.SetNext(next.GetNext())
			next.SetPrev(current.GetPrev())
		} else {
			current.GetPrev().SetNext(nil)
			d.SetTail(current.GetPrev())
		}
		d.length--
	}

}

// Reverse implements LinkedList
func (d *doublyLinkedList) Reverse() {
	var prev Node = nil
	tail := d.GetHead()
	current := tail
	for current != nil {
		prev = current.GetPrev()
		current.SetPrev(current.GetNext())
		current.SetNext(prev)
		current = current.GetPrev()
	}

	d.SetTail(tail)
	if prev != nil {
		d.SetHead(prev.GetPrev())
	}

}

// SetHead implements LinkedList
func (d *doublyLinkedList) SetHead(node Node) {
	d.head = node
}

// SetTail implements LinkedList
func (d *doublyLinkedList) SetTail(node Node) {
	d.tail = node
}

func NewDoublyLinkedList() LinkedList {
	return &doublyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}
