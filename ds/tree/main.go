package main

import (
	"github.com/nashmaniac/golang-coding/ds/tree/tree"
)

func main() {
	nums := []int{
		5, 8, 3, 4, 2, 1, // 10, 9, 7,
	}

	t := tree.NewBST()
	for _, i := range nums {
		t.Insert(i)
		t.Print()
	}

	node := t.Find(3)
	tree.Print(node)
}
