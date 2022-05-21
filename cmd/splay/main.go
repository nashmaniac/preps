package main

import "github.com/nashmaniac/golang-coding/datastructures/trees/splay"

func main() {
	tree := splay.NewSplayTree()

	numbers := []int{
		//8, 5, 6,
		8, 9, 10, 4, 5, 6, 3, 7,
	}

	for _, i := range numbers {
		tree.Insert(i)
		splay.PrintTree(tree)
	}

	// tree.Insert(8)
	// splay.PrintTree(tree)
	// tree.Insert(9)
	// splay.PrintTree(tree)
	// tree.Insert(10)
	// splay.PrintTree(tree)
	// tree.Insert(4)
	// splay.PrintTree(tree)
	// tree.Search(11)
	// splay.PrintTree(tree)
	// tree.Search(7)
	// splay.PrintTree(tree)
	// tree.Search(4)
	// splay.PrintTree(tree)
}
