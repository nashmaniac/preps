package main

import (
	"fmt"

	"github.com/nashmaniac/golang-coding/ds/unionfind/unionfind"
)

func main() {
	n := 13

	u := unionfind.NewUnionFind(n)

	edges := [][]int{
		{6, 9},
		{0, 1},
		{4, 10},
		{3, 2},
		{3, 1},
		{3, 4},
		{5, 7},
		{8, 11},
		{5, 8},
	}

	for _, edge := range edges {
		u.Add(edge[0], edge[1])
	}

	u.Add(1, 9)
	fmt.Println(u.IsConnected(0, 3))
	fmt.Println(u.IsConnected(0, 6))

}
