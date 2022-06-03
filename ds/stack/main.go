package main

import (
	"fmt"

	"github.com/nashmaniac/golang-coding/ds/stack/stack"
)

func main() {
	s := stack.NewStack[int]()

	for i := 0; i < 5; i++ {
		fmt.Printf("Pushing %d\n", i)
		s.Push(i)
		// s.Print()
	}

	for !s.IsEmpty() {
		fmt.Println(s.Pop())
		s.Print()
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Pushing %d\n", i)
		s.Push(5 - i)
		// s.Print()
	}
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
		s.Print()
	}
}
