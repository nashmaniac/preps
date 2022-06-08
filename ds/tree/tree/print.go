package tree

import (
	"fmt"
	"math"
	"strconv"
)

const (
	SPACES      = 1
	DUMMY_VALUE = math.MinInt16
)

func Print(node Node) {
	if node == nil {
		return
	}
	q := NewQueue[Node]()
	dummy := NewNode(DUMMY_VALUE)
	height := node.GetHeight()
	level := 0
	q.Offer(node)
	q.Offer(nil)

	for !q.IsEmpty() {
		node = q.Pop()
		if node == nil {
			if !q.IsEmpty() {
				q.Offer(nil)
			}
			fmt.Println()
			level++
			continue
		}
		node = node.(Node)
		spacesCount := Pow(2, height-level) * SPACES
		s := getSingleString(node, spacesCount)
		fmt.Print(s)
		if level < height {
			if node.GetLeft() == nil {
				q.Offer(dummy)
			} else {
				q.Offer(node.GetLeft())
			}

			if node.GetRight() == nil {
				q.Offer(dummy)
			} else {
				q.Offer(node.GetRight())
			}
		}
	}
}

func appendChars(bytes []rune, start int, end int, ch rune) []rune {
	for i := start; i < end; i++ {
		bytes = append(bytes, ch)
	}
	return bytes
}

func getSingleString(node Node, spacesCount int) string {
	bytes := make([]rune, 0)
	if node.GetValue() == DUMMY_VALUE {
		bytes = appendChars(bytes, 0, 2*spacesCount, ' ')
	} else {
		to := spacesCount / 2
		ch := ' '
		bytes = appendChars(bytes, 0, to, ch)
		if node.GetLeft() != nil {
			ch = '-'
		}
		bytes = appendChars(bytes, 0, to, ch)
		valueString := strconv.Itoa(node.GetValue())
		for _, c := range valueString {
			bytes = append(bytes, c)
		}
		ch = ' '
		if node.GetRight() != nil {
			ch = '-'
		}
		bytes = appendChars(bytes, len(valueString), to, ch)
		ch = ' '
		bytes = appendChars(bytes, 0, to, ch)
	}
	return string(bytes)
}
