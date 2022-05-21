package splay

import (
	"fmt"
	"math"
	"strconv"

	"github.com/nashmaniac/golang-coding/datastructures/queue"
)

type Node interface {
	GetLeft() Node
	SetLeft(node Node)
	SetRight(node Node)
	GetRight() Node
	SetParent(node Node)
	GetParent() Node
	SetValue(value int)
	GetValue() int
}

type node struct {
	value  int
	left   Node
	right  Node
	parent Node
}

// GetLeft implements Node
func (n *node) GetLeft() Node {
	return n.left
}

// GetParent implements Node
func (n *node) GetParent() Node {
	return n.parent
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

// SetParent implements Node
func (n *node) SetParent(node Node) {
	n.parent = node
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
		value:  value,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

type SplayTree interface {
	SetRoot(node Node)
	GetRoot() Node
	SetNodeToSplay(node Node)
	GetNodeToSplay() Node
	Search(value int) Node
	Insert(value int)
	Delete(value int)
}

type splayTree struct {
	Root        Node
	NodeToSplay Node
}

func Transplant(s SplayTree, u Node, v Node) {
	parent := u.GetParent()
	if parent == nil {
		s.SetRoot(v)
	} else if parent.GetLeft() == u {
		parent.SetLeft(v)
	} else if parent.GetRight() == v {
		parent.SetRight(v)
	}
	if v != nil {
		v.SetParent(parent)
	}
}

func Remove(tree SplayTree, node Node) {
	if node == nil {
		return
	}

	if node.GetLeft() == nil {
		Transplant(tree, node, node.GetRight())
	} else if node.GetRight() == nil {
		Transplant(tree, node, node.GetLeft())
	} else {
		maxNode := MaximumNode(node.GetLeft())
		node.SetValue(maxNode.GetValue())
		Remove(tree, maxNode)
	}
}

// Delete implements SplayTree
func (s *splayTree) Delete(value int) {
	node := Search(s, s.GetRoot(), value)
	if node == nil {
		if s.GetNodeToSplay() != nil {
			Splay(s.GetNodeToSplay())
			s.SetRoot(s.GetNodeToSplay())
			s.SetNodeToSplay(nil)
		}
	}
	// find maximum to the left
	maxNode := MaximumNode(node.GetLeft())
	if maxNode != nil {
		s.SetNodeToSplay(maxNode.GetParent())
		node.SetValue(maxNode.GetValue())
	}
	Remove(maxNode)
}

// GetNodeToSplay implements SplayTree
func (s *splayTree) GetNodeToSplay() Node {
	return s.NodeToSplay
}

// SetNodeToSplay implements SplayTree
func (s *splayTree) SetNodeToSplay(node Node) {
	s.NodeToSplay = node
}

// SetRoot implements SplayTree
func (s *splayTree) SetRoot(node Node) {
	s.Root = node
}

// GetRoot implements SplayTree
func (s *splayTree) GetRoot() Node {
	return s.Root
}

func RotateLeft(x Node) {
	y := x.GetParent()
	t2 := x.GetLeft()

	y.SetRight(t2)
	if t2 != nil {
		t2.SetParent(y)
	}

	parent := y.GetParent()
	if parent != nil {
		if parent.GetLeft() == y {
			parent.SetLeft(x)
		}
		if parent.GetRight() == y {
			parent.SetRight(x)
		}
	}
	x.SetParent(parent)

	x.SetLeft(y)
	y.SetParent(x)

}

func RotateRight(x Node) {
	y := x.GetParent()
	t2 := x.GetRight()

	y.SetLeft(t2)
	if t2 != nil {
		t2.SetParent(y)
	}

	parent := y.GetParent()
	if parent != nil {
		if parent.GetLeft() == y {
			parent.SetLeft(x)
		}
		if parent.GetRight() == y {
			parent.SetRight(x)
		}
	}
	x.SetParent(parent)

	x.SetRight(y)
	y.SetParent(x)
}

func ZigZigLeft(x Node) {
	y := x.GetParent()
	z := y.GetParent()

	t3 := y.GetRight()
	t2 := x.GetRight()

	parent := z.GetParent()
	if parent != nil {
		if parent.GetLeft() == z {
			parent.SetLeft(x)
		}
		if parent.GetRight() == z {
			parent.SetRight(x)
		}
	}
	x.SetParent(parent)

	z.SetLeft(t3)
	if t3 != nil {
		t3.SetParent(z)
	}

	y.SetLeft(t2)
	if t2 != nil {
		t2.SetParent(y)
	}

	y.SetRight(z)
	z.SetParent(y)

	x.SetRight(y)
	y.SetParent(z)
}

func ZigZigRight(x Node) {
	y := x.GetParent()
	z := y.GetParent()

	t2 := y.GetLeft()
	t3 := x.GetLeft()

	parent := z.GetParent()
	if parent != nil {
		if parent.GetLeft() == z {
			parent.SetLeft(x)
		}
		if parent.GetRight() == z {
			parent.SetRight(x)
		}
	}

	z.SetRight(t2)
	if t2 != nil {
		t2.SetParent(z)
	}
	y.SetRight(t3)
	if t3 != nil {
		t3.SetParent(y)
	}

	y.SetLeft(z)
	z.SetParent(y)

	x.SetLeft(y)
	y.SetParent(x)
	x.SetParent(parent)
}

func ZigZagRight(x Node) {
	y := x.GetParent()
	z := y.GetParent()

	t2 := x.GetLeft()
	t3 := x.GetRight()

	parent := z.GetParent()
	if parent != nil {
		if parent.GetLeft() == z {
			parent.SetLeft(x)
		}
		if parent.GetRight() == z {
			parent.SetRight(x)
		}
	}
	x.SetParent(parent)

	z.SetRight(t2)
	if t2 != nil {
		t2.SetParent(z)
	}

	y.SetLeft(t3)
	if t3 != nil {
		t3.SetParent(y)
	}

	x.SetLeft(z)
	z.SetParent(x)

	x.SetRight(y)
	y.SetParent(x)

}

func ZigZagLeft(x Node) {
	y := x.GetParent()
	z := y.GetParent()

	t2 := x.GetLeft()
	t3 := x.GetRight()

	parent := z.GetParent()
	if parent != nil {
		if parent.GetLeft() == z {
			parent.SetLeft(x)
		}
		if parent.GetRight() == x {
			parent.SetRight(x)
		}
	}
	x.SetParent(parent)

	y.SetRight(t2)
	if t2 != nil {
		t2.SetParent(y)
	}

	z.SetLeft(t3)
	if t3 != nil {
		t3.SetParent(z)
	}

	x.SetLeft(y)
	y.SetParent(x)

	x.SetRight(z)
	z.SetParent(x)
}

func Splay(node Node) {
	for node.GetParent() != nil {
		var grandParent Node = nil
		parent := node.GetParent()
		if parent != nil {
			grandParent = parent.GetParent()
		}

		if grandParent == nil {
			if parent.GetRight() == node {
				// Zig Right
				RotateLeft(node)
			}

			if parent.GetLeft() == node {
				// Zig Right
				RotateRight(node)
			}
		} else {
			if grandParent.GetRight() == parent {
				if parent.GetLeft() == node {
					// zig zag right side
					ZigZagRight(node)
				} else {
					// zig zig right side
					ZigZigRight(node)
				}
			}

			if grandParent.GetLeft() == parent {
				if parent.GetLeft() == node {
					// zig zig left side
					ZigZigLeft(node)
				} else {
					// zig zag left side
					ZigZagLeft(node)
				}
			}
		}
	}
}

func Insert(s SplayTree, node Node, value int) Node {
	if node == nil {
		n := NewNode(value)
		s.SetNodeToSplay(n)
		return n
	}

	var child Node = nil
	if value <= node.GetValue() {
		child = Insert(s, node.GetLeft(), value)
		node.SetLeft(child)
	} else {
		child = Insert(s, node.GetRight(), value)
		node.SetRight(child)
	}

	if child != nil {
		child.SetParent(node)
	}

	return node
}

// Insert implements SplayTree
func (s *splayTree) Insert(value int) {
	root := Insert(s, s.GetRoot(), value)
	s.SetRoot(root)
	if s.NodeToSplay != nil {
		Splay(s.NodeToSplay)
	}
	s.SetRoot(s.NodeToSplay)
	s.NodeToSplay = nil
}

func Search(s SplayTree, root Node, value int) Node {
	if root == nil {
		return nil
	}
	s.SetNodeToSplay(root)
	if root.GetValue() == value {
		return root
	} else if value < root.GetValue() {
		return Search(s, root.GetLeft(), value)
	} else {
		return Search(s, root.GetRight(), value)
	}
}

// Search implements SplayTree
func (s *splayTree) Search(value int) Node {
	node := Search(s, s.GetRoot(), value)
	if s.GetNodeToSplay() != nil {
		Splay(s.GetNodeToSplay())
		s.SetRoot(s.GetNodeToSplay())
		s.SetNodeToSplay(nil)
	}
	return node
}

func MaximumNode(node Node) Node {
	if node == nil {
		return nil
	}
	for node.GetRight() != nil {
		node = node.GetRight()
	}
	return node
}

func NewSplayTree() SplayTree {
	return &splayTree{
		Root: nil,
	}
}

var (
	SPACE = 2
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func GetHeight(root Node) int {
	if root == nil {
		return 0
	}
	return 1 + max(GetHeight(root.GetLeft()), GetHeight(root.GetRight()))
}

func PrintTree(tree SplayTree) {
	root := tree.GetRoot()
	if root == nil {
		return
	}

	height := GetHeight(root)
	level := 0
	dummy := NewNode(math.MinInt32)
	q := queue.NewQueue[Node]()
	q.Enqueue(root)
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
			node := n.(Node)
			spaces := Pow(2, height-level) * SPACE
			// log.Println(spaces)
			fmt.Print(string(buildSingleLine(node, spaces)))
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
func buildRune(runes []rune, i int, to int, ch rune) []rune {
	for j := i; j < to; j++ {
		runes = append(runes, ch)
	}
	return runes
}
func buildSingleLine(node Node, spaces int) []rune {
	runes := make([]rune, 0)
	if node.GetValue() == math.MinInt32 {
		for i := 0; i < 2*spaces; i++ {
			runes = append(runes, ' ')
		}
		return runes
	} else {
		i := 0
		to := spaces / 2
		ch := ' '
		runes = buildRune(runes, i, to, ch)
		if node.GetLeft() != nil {
			ch = '_'
		}
		runes = buildRune(runes, i, to, ch)
		value := strconv.Itoa(node.GetValue())
		for _, c := range value {
			runes = append(runes, c)
		}
		ch = ' '
		if node.GetRight() != nil {
			ch = '_'
		}
		runes = buildRune(runes, len(value), to, ch)
		ch = ' '
		runes = buildRune(runes, i, to, ch)
	}
	// log.Println(runes)
	return runes
}
