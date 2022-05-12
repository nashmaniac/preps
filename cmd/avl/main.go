package main

import (
	"github.com/nashmaniac/golang-coding/datastructures/trees/bbst"
	"github.com/nashmaniac/golang-coding/datastructures/utils"
)

func main() {
	l := bbst.NewAVL()
	for i := 0; i < 15; i++ {
		l.Add(i)
		utils.PrintTree(l.GetRoot())
	}
}
