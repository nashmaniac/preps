package main

import "github.com/nashmaniac/golang-coding/ds/vector/vector"

func main() {
	cap := 0
	v := vector.NewVector(cap)
	n := 6
	for i := 0; i < n; i++ {
		v.Add(i + 1)
		v.Print()
	}
	v.Delete(0)
	v.Print()
	v.Delete(0)
	v.Print()
	v.Delete(0)
	v.Print()
}
