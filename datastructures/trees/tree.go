package trees

import (
	"fmt"
	"math"
	"strconv"

	"github.com/nashmaniac/golang-coding/datastructures/queue"
)

type BinarySearchTree interface {
	Insert(value int)
	Delete(value int)
	GetMaxValue() int
	GetMinValue() int
	Traverse()
	Print()
	GetRoot() Node
	SetRoot(node Node)
}

type binarySearchTree struct {
	root Node
}

// Delete implements BinarySearchTree
func (b *binarySearchTree) Delete(value int) {
	panic("unimplemented")
}

func GetMaxNode(node Node) Node {
	if node == nil {
		return nil
	}
	for node.GetRight() != nil {
		node = node.GetRight()
	}
	return node
}

// GetMaxValue implements BinarySearchTree
func (b *binarySearchTree) GetMaxValue() int {
	node := GetMaxNode(b.root)
	if node == nil {
		return -1
	}
	return node.GetValue()
}

func GetMinNode(node Node) Node {
	if node == nil {
		return nil
	}
	for node.GetLeft() != nil {
		node = node.GetLeft()
	}
	return node
}

// GetMinValue implements BinarySearchTree
func (b *binarySearchTree) GetMinValue() int {
	node := GetMinNode(b.root)
	if node == nil {
		return -1
	}
	return node.GetValue()
}

// GetRoot implements BinarySearchTree
func (b *binarySearchTree) GetRoot() Node {
	return b.root
}

// Insert implements BinarySearchTree
func (b *binarySearchTree) Insert(value int) {
	panic("unimplemented")
}

func populateRunes(runes []rune, start int, end int, ch rune) []rune {
	for i := start; i < end; i++ {
		runes = append(runes, ch)
	}
	return runes
}

func buildSingleLine(node Node, gap int) []rune {
	runes := make([]rune, 0)
	if node.GetValue() == math.MinInt16 {
		runes = populateRunes(runes, 0, 2*gap, ' ')
	} else {
		start := 0
		end := gap / 2
		ch := ' '
		runes = populateRunes(runes, start, end, ch)
		if node.GetLeft() != nil {
			ch = '_'
		}
		runes = populateRunes(runes, start, end, ch)
		value := strconv.Itoa(node.GetValue())
		for _, c := range value {
			runes = append(runes, c)
		}
		ch = ' '
		if node.GetRight() != nil {
			ch = '_'
		}
		runes = populateRunes(runes, len(value), end, ch)
		ch = ' '
		runes = populateRunes(runes, start, end, ch)
	}
	return runes
}

func PrintTree(node Node) {
	if node == nil {
		return
	}
	spaces := 2
	height := node.GetHeight()
	level := 0

	q := queue.NewQueue[Node]()
	dummy := NewNode(math.MinInt16)
	q.Enqueue(node)
	q.Enqueue(nil)
	for !q.IsEmpty() {
		n := q.Dequeue()
		if n == nil {
			if !q.IsEmpty() {
				q.Enqueue(nil)
			}
			fmt.Println()
			level++
		} else {
			node = n.(Node)
			gap := pow(2, height-level) * spaces
			runes := buildSingleLine(node, gap)
			fmt.Print(string(runes))
			if level < height {
				if node.GetLeft() != nil {
					q.Enqueue(node.GetLeft())
				} else {
					q.Enqueue(dummy)
				}

				if node.GetRight() != nil {
					q.Enqueue(node.GetRight())
				} else {
					q.Enqueue(dummy)
				}
			}

		}
	}

}

// Print implements BinarySearchTree
func (b *binarySearchTree) Print() {
	PrintTree(b.GetRoot())
}

// SetRoot implements BinarySearchTree
func (b *binarySearchTree) SetRoot(node Node) {
	b.root = node
}

func inOrderTraversal(node Node) {
	if node == nil {
		return
	}
	inOrderTraversal(node.GetLeft())
	fmt.Println(node.GetValue())
	inOrderTraversal(node.GetRight())
}

// Traverse implements BinarySearchTree
func (b *binarySearchTree) Traverse() {
	inOrderTraversal(b.GetRoot())
}

func NewBinarySearchTree() BinarySearchTree {
	return &binarySearchTree{
		root: nil,
	}
}
