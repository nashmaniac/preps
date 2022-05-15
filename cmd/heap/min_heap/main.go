package main

import (
	"fmt"

	"github.com/nashmaniac/golang-coding/datastructures/heap"
)

func printOutput(m heap.Heap) {
	n := m.HeapSize()
	elements := m.ElementList()
	for i := 0; i < n; i++ {
		fmt.Print(elements[i], " ")
	}
	fmt.Println()
}

func main() {

	m := heap.NewMinHeap()

	m.Offer(3)
	printOutput(m)
	m.Offer(5)
	printOutput(m)
	m.Offer(2)
	printOutput(m)
	m.Offer(6)
	printOutput(m)
	m.Offer(7)
	printOutput(m)
	m.Offer(1)
	printOutput(m)

	for !m.IsEmpty() {
		n := m.Poll()
		// printOutput(m)
		fmt.Println(*n)
	}

}
