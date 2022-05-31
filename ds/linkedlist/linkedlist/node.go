package linkedlist

type Node interface {
	GetValue() int
	SetValue(value int)

	GetNext() Node
	GetPrev() Node

	SetNext(node Node)
	SetPrev(node Node)
}

type node struct {
	value int
	next  Node
	prev  Node
}

// SetValue implements Node
func (n *node) SetValue(value int) {
	n.value = value
}

// GetNext implements Node
func (n *node) GetNext() Node {
	return n.next
}

// GetPrev implements Node
func (n *node) GetPrev() Node {
	return n.prev
}

// GetValue implements Node
func (n *node) GetValue() int {
	return n.value
}

// SetNext implements Node
func (n *node) SetNext(node Node) {
	n.next = node
}

// SetPrev implements Node
func (n *node) SetPrev(node Node) {
	n.prev = node
}

func NewNode(value int) Node {
	return &node{
		value: value,
		next:  nil,
		prev:  nil,
	}
}
