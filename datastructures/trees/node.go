package trees

import "math"

type Node interface {
	GetValue() int
	SetValue(value int)
	GetLeft() Node
	SetLeft(node Node)
	GetRight() Node
	SetRight(node Node)
	GetHeight() int
}

type node struct {
	value int
	left  Node
	right Node
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func GetHeight(n Node) int {
	if n == nil {
		return 0
	}

	return 1 + max(GetHeight(n.GetLeft()), GetHeight(n.GetRight()))
}

// GetHeight implements Node
func (n *node) GetHeight() int {
	return GetHeight(n)
}

// GetLeft implements Node
func (n *node) GetLeft() Node {
	return n.left
}

// GetRight implements Node
func (n *node) GetRight() Node {
	return n.right
}

// GetValue implements Node
func (n *node) GetValue() int {
	return n.value
}

// SetLeft implements Node
func (n *node) SetLeft(node Node) {
	n.left = node
}

// SetRight implements Node
func (n *node) SetRight(node Node) {
	n.right = node
}

// SetValue implements Node
func (n *node) SetValue(value int) {
	n.value = value
}

func NewNode(value int) Node {
	return &node{
		value: value,
		left:  nil,
		right: nil,
	}
}
