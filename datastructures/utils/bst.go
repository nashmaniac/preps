package utils

import (
	"fmt"
	"math"
	"strconv"

	"github.com/nashmaniac/golang-coding/datastructures/models"
	"github.com/nashmaniac/golang-coding/datastructures/queue"
)

var (
	SPACE = 2
)

func GetHeight(node models.Node) int {
	if node == nil {
		return 0
	}
	return 1 + Max(GetHeight(node.GetLeft()), GetHeight(node.GetRight()))
}

func PrintTree(node models.Node) {
	height := GetHeight(node)
	level := 0
	dummyNode := models.NewTreeNode(math.MinInt64)
	nodeLocation := 0
	q := queue.NewQueue[models.Node]()
	q.Enqueue(node)
	q.Enqueue(nil)
	for !q.IsEmpty() {
		element := q.Dequeue()
		if element == nil {
			node = nil
		} else {
			node = element.(models.Node)
		}
		if node == nil {
			if !q.IsEmpty() {
				q.Enqueue(nil)
			}
			fmt.Println("")
			level++
		} else {
			nodeLocation = Pow(2, height-level) * SPACE
			fmt.Print(string(BuildSingleLine(node, nodeLocation)))
			if level < height {
				if node.GetLeft() != nil {
					q.Enqueue(node.GetLeft())
				} else {
					q.Enqueue(dummyNode)
				}

				if node.GetRight() != nil {
					q.Enqueue(node.GetRight())
				} else {
					q.Enqueue(dummyNode)
				}
			}

		}
	}
}

func BuildSingleLine(node models.Node, spaces int) []rune {
	runes := make([]rune, 0)

	i := 0
	to := spaces / 2

	if node.GetValue() == math.MinInt64 {
		for ; i < 2*spaces; i++ {
			runes = append(runes, ' ')
		}
		return runes
	}

	for ; i < to; i++ {
		runes = append(runes, ' ')
	}
	ch := ' '
	if node.GetLeft() != nil {
		ch = '_'
	}
	to += spaces / 2
	for ; i < to; i++ {
		runes = append(runes, ch)
	}
	value := strconv.Itoa(node.GetValue())
	for _, c := range value {
		runes = append(runes, c)
	}

	ch = ' '
	if node.GetRight() != nil {
		ch = '_'
	}
	to += spaces / 2
	for i += len(value); i < to; i++ {
		runes = append(runes, ch)
	}

	ch = ' '
	to += spaces / 2
	for ; i < to; i++ {
		runes = append(runes, ch)
	}
	return runes
}
