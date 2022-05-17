package main

import "github.com/nashmaniac/golang-coding/datastructures/unionfind"

func main() {
	nodeCount := 10
	u := unionfind.NewUnionFind(nodeCount)
	u.Union(2, 1)
	u.Union(4, 3)
	u.Union(8, 4)
	u.Union(9, 3)
	u.Union(6, 5)
	u.Union(1, 5)
	u.Union(2, 5)
	u.PrintMembers()
}
