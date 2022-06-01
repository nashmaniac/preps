package main

import (
	"github.com/nashmaniac/golang-coding/ds/linkedlist/linkedlist"
)

func main() {
	l := linkedlist.NewDoublyLinkedList()

	n := 5
	for i := 0; i < n; i++ {
		l.AddBack(i + 1)
		l.Print()
	}

	for i := n + 1; i < n+5; i++ {
		l.AddFront(i + 1)
		l.Print()
	}

	// for i := 0; i < 5; i++ {
	// 	value := l.PopBack()
	// 	fmt.Println(value)
	// 	l.Print()
	// }

	// for i := 0; i < 5; i++ {
	// 	value := l.PopFront()
	// 	fmt.Println(value)
	// 	l.Print()
	// }

	// for i := 0; i < 4; i++ {
	// 	l.Delete(i)
	// 	l.Print()
	// }

	l.Reverse()
	l.Print()
	// l.Remove(1)
	// l.Print()
	// l.Remove(5)
	// l.Print()
	// l.AddFront(10)
	// l.Print()
	// l.AddBack(15)
	// l.Print()
	// l.Remove(3)
	// l.Print()

	index := []int{0, 7, 4}
	for _, i := range index {
		l.Delete(i)
		l.Print()
	}
}
