package main

import (
	"fmt"
	"math/rand"

	"github.com/nashmaniac/golang-coding/datastructures/trees"
)

func main() {

	nodes := make([]int, 0)
	// rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		nodes = append(nodes, rand.Intn(100))
	}
	fmt.Println(nodes)

	avl := trees.NewAVLTree()
	for _, i := range nodes {
		avl.Insert(i)
		avl.Print()
	}

	avl.Delete(47)
	avl.Print()

}
