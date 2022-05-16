package heap

type Heap interface {
	Offer(value int)
	Poll() *int
	Peek() *int
	IsEmpty() bool
	HeapSize() int
	ElementList() []int
}

type HeapSort interface {
	Sort(m Heap)
}
