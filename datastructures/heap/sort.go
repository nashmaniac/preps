package heap

func AscendingSort(elements []int) {

}

func DescendingSort(elements []int) []int {
	h := NewMinHeap()
	n := len(elements)

	for i := 0; i < n; i++ {
		h.Offer(elements[i])
	}
	size := n - 1
	for i := 0; i < n; i++ {
		swap(h.ElementList(), 0, size)
		BubbleDown(h.ElementList(), 0, size)
		size--
	}
	return h.ElementList()
}
