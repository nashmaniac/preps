package models

type Node interface {
	GetValue() int

	GetLeft() Node
	GetRight() Node
	GetParent() Node
	GetHeight() int
	GetBalanceFactor() int

	SetLeft(node Node)
	SetRight(node Node)
	SetParent(node Node)
	SetValue(value int)
	SetHeight(value int)
	SetBalanceFactor(value int)
}

type TreeNode struct {
	value         int
	left          Node
	right         Node
	parent        Node
	height        int
	balanceFactor int
}

// GetBalanceFactor implements Node
func (t *TreeNode) GetBalanceFactor() int {
	return t.balanceFactor
}

// GetHeight implements Node
func (t *TreeNode) GetHeight() int {
	return t.height
}

// SetBalanceFactor implements Node
func (t *TreeNode) SetBalanceFactor(value int) {
	t.balanceFactor = value
}

// SetHeight implements Node
func (t *TreeNode) SetHeight(value int) {
	t.height = value
}

// GetLeft implements Node
func (t *TreeNode) GetLeft() Node {
	return t.left
}

// GetParent implements Node
func (t *TreeNode) GetParent() Node {
	return t.parent
}

// GetRight implements Node
func (t *TreeNode) GetRight() Node {
	return t.right
}

// GetValue implements Node
func (t *TreeNode) GetValue() int {
	return t.value
}

// SetLeft implements Node
func (t *TreeNode) SetLeft(node Node) {
	t.left = node
}

// SetParent implements Node
func (t *TreeNode) SetParent(node Node) {
	t.parent = node
}

// SetRight implements Node
func (t *TreeNode) SetRight(node Node) {
	t.right = node
}

// SetValue implements Node
func (t *TreeNode) SetValue(value int) {
	t.value = value
}

func NewTreeNode(value int) Node {
	return &TreeNode{
		value:  value,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}
