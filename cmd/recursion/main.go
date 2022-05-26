package main

import (
	"fmt"

	"github.com/nashmaniac/golang-coding/theories/recrusion"
)

func main() {
	n := 15
	recrusion.Print1ToN(n)
	fmt.Println()
	recrusion.PrintNTo1(n)
}
