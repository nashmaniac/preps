package main

import (
	"github.com/nashmaniac/golang-coding/ds/tree/tree"
)

func main() {
	nums := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	t := tree.NewAVl()
	for _, i := range nums {
		t.Insert(i)
		//t.Print()
	}
	t.Delete(8)
	//node := t.Find(3)
	t.Delete(4)
	t.Delete(5)
}
