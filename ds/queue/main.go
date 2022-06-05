package main

import (
	"fmt"

	"github.com/nashmaniac/golang-coding/ds/queue/queue"
)

func queueSimulation() {
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

func heapSimulation() {
	q := queue.NewPriorityQueue[string](true)

	type input struct {
		key   int
		value string
	}
	elements := []input{
		{5, "shetu"},
		{6, "Farhan"},
		{1, "Upama"},
		{0, "Sina"},
		{10, "Satari"},
		{3, "Aksa"},
		{1, "Sweet"},
	}

	for _, element := range elements {
		q.OfferItem(element.key, element.value)
		q.Print()
	}

	for !q.IsEmpty() {
		member := q.Poll()
		fmt.Printf("Removed %v\n", member)
		q.Print()
	}
}

func main() {
	heapSimulation()
}
