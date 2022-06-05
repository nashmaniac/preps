package queue

import "fmt"

type priorityQueue struct {
	isMinQueue bool
	elements   []QueueElement
	size       int
}

// Offer implements Queue
func (p *priorityQueue) Offer(item Element) {
	panic("unimplemented")
}

func Left(index int) int {
	return 2*index + 1
}

func Right(index int) int {
	return 2*index + 2
}

func HasLeft(index int, size int) bool {
	return Left(index) < size
}

func HasRight(index int, size int) bool {
	return Right(index) < size
}

func Parent(index int) int {
	return (index - 1) / 2
}

func swap(elements []QueueElement, i, j int) {
	e := elements[i]
	elements[i] = elements[j]
	elements[j] = e
}

func (p *priorityQueue) upHeap(index int) {
	for index != 0 {
		parentIndex := Parent(index)
		parent := p.elements[parentIndex]
		element := p.elements[index]

		if p.isMinQueue && parent.GetKey() < element.GetKey() {
			break
		}

		if !p.isMinQueue && parent.GetKey() > element.GetKey() {
			break
		}
		swap(p.elements, parentIndex, index)
		index = parentIndex
	}
}

// IsEmpty implements Queue
func (p *priorityQueue) IsEmpty() bool {
	return p.size == 0
}

// Offer implements Queue
func (p *priorityQueue) OfferItem(key int, item Element) {
	p.elements = append(p.elements, NewQueueElement(key, item))
	p.size++
	p.upHeap(p.size - 1)

}

func (p *priorityQueue) downHeap(index int) {
	for HasLeft(index, p.size) {
		parent := p.elements[index]
		minIndex := Left(index)
		minElement := p.elements[minIndex]
		if HasRight(index, p.size) {
			rightIndex := Right(index)
			rightElement := p.elements[rightIndex]

			if (p.isMinQueue && (rightElement.GetKey() < minElement.GetKey())) || (!p.isMinQueue && (rightElement.GetKey() > minElement.GetKey())) {
				minIndex = rightIndex
				minElement = rightElement
			}
		}

		if (p.isMinQueue && (parent.GetKey() < minElement.GetKey())) || (!p.isMinQueue && (parent.GetKey() > minElement.GetKey())) {
			break
		}

		swap(p.elements, index, minIndex)
		index = minIndex
	}
}

// Poll implements Queue
func (p *priorityQueue) Poll() Element {
	if p.IsEmpty() {
		return nil
	}

	element := p.elements[0]
	swap(p.elements, 0, p.size-1)
	p.size--
	p.elements = p.elements[:p.size]
	p.downHeap(0)
	return element
}

// Print implements Queue
func (p *priorityQueue) Print() {
	for _, i := range p.elements {
		fmt.Printf("(%d)%s --> ", i.GetKey(), i.GetValue())
	}
	fmt.Printf("nil\n")
}

func NewPriorityQueue[e Element](minQueue bool) Queue {
	return &priorityQueue{
		isMinQueue: minQueue,
		elements:   make([]QueueElement, 0),
		size:       0,
	}
}

func NewMinQueue[e Element]() Queue {
	return NewPriorityQueue[e](true)
}

func NewMaxQueue[e Element]() Queue {
	return NewPriorityQueue[e](false)
}
