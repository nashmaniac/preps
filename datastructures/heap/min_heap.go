package heap

type minHeap struct {
	Elements []int
	Size     int
}

func (m *minHeap) ElementList() []int {
	return m.Elements
}

func (m *minHeap) HeapSize() int {
	return m.Size
}

// IsEmpty implements Heap
func (m *minHeap) IsEmpty() bool {
	return m.Size == 0
}

func swap(elements []int, i, j int) {
	temp := elements[i]
	elements[i] = elements[j]
	elements[j] = temp
}

func BubbleDown(elements []int, index int, size int) {

	for index < size {
		l := 2*index + 1
		r := 2*index + 2

		minIndex := -1

		if r < size {
			if elements[l] <= elements[r] {
				minIndex = l
			} else {
				minIndex = r
			}
			swap(elements, minIndex, index)
		} else {
			if l < size {
				swap(elements, l, index)
			}
		}
		index++
	}
}

func BubbleUp(elements []int, index int, size int) {
	for index != 0 {
		parent := index / 2
		parentValue := elements[parent]
		if parentValue > elements[index] {
			t := elements[index]
			elements[index] = parentValue
			elements[parent] = t
		}
		index = parent
	}
}

// Offer implements Heap
func (m *minHeap) Offer(value int) {
	m.Elements = append(m.Elements, value)
	index := m.Size
	m.Size++
	BubbleUp(m.Elements, index, m.Size)
}

// Peek implements Heap
func (m *minHeap) Peek() *int {
	if m.IsEmpty() {
		return nil
	}
	return &m.Elements[0]
}

// Poll implements Heap
func (m *minHeap) Poll() *int {
	if m.IsEmpty() {
		return nil
	}

	n := m.Elements[0]
	swap(m.Elements, 0, m.Size-1)
	m.Size--
	m.Elements = m.Elements[:m.Size]
	BubbleDown(m.Elements, 0, m.Size)
	return &n
}

func NewMinHeap() Heap {
	elements := make([]int, 0)
	return &minHeap{
		Elements: elements,
		Size:     0,
	}
}
