package main

import (
	"fmt"

	"github.com/nashmaniac/golang-coding/ds/queue/queue"
)

func main() {
	q := queue.NewQueue[int]()

	for i := 0; i < 10; i++ {
		fmt.Printf("Pushing %d\n", i+1)
		q.Offer(i + 1)
		q.Print()
	}

	for i := 0; i < 5; i++ {
		val := q.Poll()
		fmt.Printf("Removed %d\n", val)
		q.Print()
		fmt.Printf("Repushing %d\n", val)
		q.Offer(val)
		q.Print()
	}
}
