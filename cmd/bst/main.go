package main

import (
	"math/rand"

	"github.com/nashmaniac/golang-coding/datastructures/trees"
)

func main() {
	bst := trees.NewBST()

	nodes := make([]int, 0)
	// rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		nodes = append(nodes, rand.Intn(100))
	}

	for _, i := range nodes {
		bst.Insert(i)
		bst.Print()
	}

	bst.Delete(47)
	bst.Print()

	bst.Delete(59)
	bst.Print()
	bst.Delete(87)
	bst.Print()

}
