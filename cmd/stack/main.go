package main

import (
	"log"

	"github.com/nashmaniac/golang-coding/datastructures/stack"
)

func main() {
	s := stack.NewStack()

	s.Push(1)
	log.Println(s.IsEmpty())
	log.Println(s.Peek())
	log.Println(s.Pop())
	log.Println(s.IsEmpty())
	s.Push(2)
	s.Push(3)

	log.Println(s.IsEmpty())
	log.Println(s.Peek())
	log.Println(s.Pop())
	log.Println(s.IsEmpty())
	log.Println(s.Peek())

}
