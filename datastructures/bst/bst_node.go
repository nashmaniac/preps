package bst

type BSTNode struct {
	value  int
	key    int
	left   *BSTNode
	right  *BSTNode
	parent *BSTNode
}

func NewBSTNodeWithParent(value int, parent *BSTNode) *BSTNode {
	return &BSTNode{
		value:  value,
		key:    value,
		left:   nil,
		right:  nil,
		parent: parent,
	}
}

func NewBSTNode(value int) *BSTNode {
	return &BSTNode{
		value:  value,
		key:    value,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

func (bstnode *BSTNode) GetValue() int {
	return bstnode.value
}

func (bstnode *BSTNode) GetKey() int {
	return bstnode.key
}

func (bstnode *BSTNode) GetLeft() *BSTNode {
	return bstnode.left
}

func (bstnode *BSTNode) GetRight() *BSTNode {
	return bstnode.right
}

func (bstnode *BSTNode) GetParent() *BSTNode {
	return bstnode.parent
}

func (bstnode *BSTNode) SetValue(value int) *BSTNode {
	bstnode.value = value
	return bstnode
}

func (bstnode *BSTNode) SetKey(key int) *BSTNode {
	bstnode.key = key
	return bstnode
}

func (bstnode *BSTNode) SetLeft(left *BSTNode) *BSTNode {
	bstnode.left = left
	return bstnode
}

func (bstnode *BSTNode) SetRight(right *BSTNode) *BSTNode {
	bstnode.right = right
	return bstnode
}

func (bstnode *BSTNode) SetParent(parent *BSTNode) *BSTNode {
	bstnode.parent = parent
	return bstnode
}
