package stack

type Element interface {
	int | string | interface{}
}

type Node interface {
	GetValue() Element
	SetValue(value Element)

	SetNext(node Node)
	GetNext() Node

	SetPrev(node Node)
	GetPrev() Node
}

type node struct {
	value Element
	next  Node
	prev  Node
}

// GetNext implements Node
func (n *node) GetNext() Node {
	return n.next
}

// GetPrev implements Node
func (n *node) GetPrev() Node {
	return n.prev
}

// SetNext implements Node
func (n *node) SetNext(node Node) {
	n.next = node
}

// SetPrev implements Node
func (n *node) SetPrev(node Node) {
	n.prev = node
}

// SetValue implements Node
func (n *node) SetValue(value Element) {
	n.value = value
}

// GetValue implements Node
func (n *node) GetValue() Element {
	return n.value
}

func NewNode[e Element](value e) Node {
	return &node{
		value: value,
		next:  nil,
		prev:  nil,
	}
}
