package tree

type Node interface {
	GetValue() int
	SetValue(item int)

	GetRight() Node
	SetRight(node Node)

	GetLeft() Node
	SetLeft(node Node)

	GetParent() Node
	SetParent(node Node)

	GetHeight() int
	SetHeight(i int)

	GetBalanceFactor() int
	SetBalanceFactor(i int)
}

type node struct {
	value  int
	right  Node
	left   Node
	parent Node
	height int
	bf     int
}

func (node *node) SetHeight(i int) {
	//TODO implement me
	node.height = i
}

func (node *node) GetBalanceFactor() int {
	//TODO implement me
	return node.bf
}

func (node *node) SetBalanceFactor(i int) {
	//TODO implement me
	node.bf = i
}

func getHeight(node Node) int {
	if node == nil {
		return 0
	}

	return 1 + Max(getHeight(node.GetLeft()), getHeight(node.GetRight()))
}

func (node *node) GetHeight() int {
	return getHeight(node)
}

func (node *node) GetValue() int {
	return node.value
}

func (node *node) GetRight() Node {
	return node.right
}

func (node *node) GetLeft() Node {
	return node.left
}

func (node *node) GetParent() Node {
	return node.parent
}

func (node *node) SetValue(value int) {
	node.value = value
}

func (node *node) SetRight(right Node) {
	node.right = right
}

func (node *node) SetLeft(left Node) {
	node.left = left
}

func (node *node) SetParent(parent Node) {
	node.parent = parent
}

func NewNode(item int) Node {
	return &node{
		value:  item,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}
