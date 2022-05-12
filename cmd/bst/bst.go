package main

import (
	"fmt"

	"github.com/nashmaniac/golang-coding/datastructures/trees/bst"
	"github.com/nashmaniac/golang-coding/datastructures/utils"
)

func main() {
	t := bst.NewBST()

	t.Add(16)
	t.Add(20)
	t.Add(17)
	t.Add(18)
	t.Add(19)
	t.Add(15)
	t.Add(12)
	t.Add(13)
	t.Add(14)
	utils.PrintTree(t.GetRoot())
	fmt.Println("------------------------------------------------------------------------------------")
	t.Delete(16)
	utils.PrintTree(t.GetRoot())
	t.Delete(15)
	fmt.Println("------------------------------------------------------------------------------------")
	utils.PrintTree(t.GetRoot())
	t.Delete(14)
	fmt.Println("------------------------------------------------------------------------------------")
	utils.PrintTree(t.GetRoot())

}
