package recrusion

import "fmt"

func Print1ToN(n int) {
	if n == 0 {
		return
	}
	Print1ToN(n - 1)
	fmt.Println(n)
}

func PrintNTo1(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	PrintNTo1(n - 1)
}
