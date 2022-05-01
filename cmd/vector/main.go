package main

import (
	"github.com/nashmaniac/golang-coding/datastructures/vector"
)

func main() {
	n := 1
	v := vector.NewVector[string](n)
	v.Push("hello")
	v.Print()
	v.Push("World")
	v.Print()
}
