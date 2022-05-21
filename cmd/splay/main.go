package main

import "github.com/nashmaniac/golang-coding/datastructures/trees/splay"

func main() {
	tree := splay.NewSplayTree()

	numbers := []int{
		//8, 5, 6,
		// 5, 10, 7,
	}

	for _, i := range numbers {
		tree.Insert(i)
		splay.PrintTree(tree)
	}
}
