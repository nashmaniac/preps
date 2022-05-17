package heap

type Element interface {
	string | int | int64 | interface{}
}

type Heap interface {
	Insert(k int, v Element)
	Remove() HeapElement
	Peek() HeapElement
	UpHeap(index int)
	DownHeap(index int)
	IsEmpty() bool
	ElementList() []HeapElement
}
type HeapElement interface {
	GetKey() int
	GetValue() Element
}

type heapElement struct {
	Key   int
	Value Element
}

func (h *heapElement) GetKey() int {
	return h.Key
}

func (h *heapElement) GetValue() Element {
	return h.Value
}

func NewHeapElement(k int, v Element) HeapElement {
	return &heapElement{
		Key:   k,
		Value: v,
	}
}

type heapImplementation struct {
	IsMaxHeap bool
	Size      int
	Elements  []HeapElement
}

// ElementList implements Heap
func (h *heapImplementation) ElementList() []HeapElement {
	return h.Elements
}

func Parent(index int) int {
	return (index - 1) / 2
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

// DownHeap implements Heap
func (h *heapImplementation) DownHeap(index int) {
	for HasLeft(index, h.Size) {
		left := Left(index)
		minIndex := left

		if HasRight(index, h.Size) {
			right := Right(index)
			if h.IsMaxHeap {
				if h.Elements[left].GetKey() < h.Elements[right].GetKey() {
					minIndex = right
				} else {
					minIndex = left
				}
			} else {
				if h.Elements[left].GetKey() < h.Elements[right].GetKey() {
					minIndex = left
				} else {
					minIndex = right
				}
			}
		}

		swap(h.Elements, minIndex, index)
		index = minIndex
	}
}

func swap(elements []HeapElement, i, j int) {
	m := elements[i]
	elements[i] = elements[j]
	elements[j] = m
}

// UpHeap implements Heap
func (h *heapImplementation) UpHeap(index int) {
	for index > 0 {
		parent := Parent(index)
		if h.IsMaxHeap {
			if h.Elements[parent].GetKey() > h.Elements[index].GetKey() {
				break
			}
		} else {
			if h.Elements[parent].GetKey() <= h.Elements[index].GetKey() {
				break
			}
		}
		swap(h.Elements, parent, index)
		index = parent
	}
}

func (h *heapImplementation) IsEmpty() bool {
	return h.Size == 0
}

// Insert implements Heap
func (h *heapImplementation) Insert(k int, v Element) {
	e := NewHeapElement(k, v)
	h.Elements = append(h.Elements, e)
	h.Size++
	h.UpHeap(h.Size - 1)
}

// Peek implements Heap
func (h *heapImplementation) Peek() HeapElement {
	if h.IsEmpty() {
		return nil
	}
	return h.Elements[0]
}

// Remove implements Heap
func (h *heapImplementation) Remove() HeapElement {
	firstElement := h.Elements[0]
	swap(h.Elements, 0, h.Size-1)
	h.Elements = h.Elements[:h.Size-1]
	h.Size--
	h.DownHeap(0)
	return firstElement
}

func MaxHeap[e Element]() Heap {
	return &heapImplementation{
		Size:      0,
		Elements:  []HeapElement{},
		IsMaxHeap: true,
	}
}

func MinHeap[e Element]() Heap {
	return &heapImplementation{
		Size:      0,
		Elements:  []HeapElement{},
		IsMaxHeap: false,
	}
}
