package main

import (
	"fmt"

	"github.com/nashmaniac/golang-coding/datastructures/heap"
)

func Print(h []heap.HeapElement) {
	// l := len(h)
	for _, i := range h {
		fmt.Print(i.GetValue(), " ")
	}
	fmt.Println()
}

func main() {
	numbers := []int{
		5, 1, 6, 3, 4, 7, 2, 9, 8, 10,
	}

	maxHeap := heap.MinHeap[int]()
	for _, i := range numbers {
		maxHeap.Insert(i, i)
		heapList := maxHeap.ElementList()
		Print(heapList)
	}

	for !maxHeap.IsEmpty() {
		fmt.Println(maxHeap.Remove().GetValue())
		Print(maxHeap.ElementList())
	}

}
