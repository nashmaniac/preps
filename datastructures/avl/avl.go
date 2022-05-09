package avl

type AVLTree interface {
	Insert(value int)
	GetRoot() *AVLNode
}

type AVLNode struct {
	value         int
	right         *AVLNode
	left          *AVLNode
	parent        *AVLNode
	height        int
	balanceFactor int
}

func (avlnode *AVLNode) GetValue() int {
	return avlnode.value
}

func (avlnode *AVLNode) GetRight() *AVLNode {
	return avlnode.right
}

func (avlnode *AVLNode) GetLeft() *AVLNode {
	return avlnode.left
}

func (avlnode *AVLNode) GetParent() *AVLNode {
	return avlnode.parent
}

func (avlnode *AVLNode) GetHeight() int {
	return avlnode.height
}

func (avlnode *AVLNode) GetBalanceFactor() int {
	return avlnode.balanceFactor
}

func (avlnode *AVLNode) SetValue(value int) *AVLNode {
	avlnode.value = value
	return avlnode
}

func (avlnode *AVLNode) SetRight(right *AVLNode) *AVLNode {
	avlnode.right = right
	return avlnode
}

func (avlnode *AVLNode) SetLeft(left *AVLNode) *AVLNode {
	avlnode.left = left
	return avlnode
}

func (avlnode *AVLNode) SetParent(parent *AVLNode) *AVLNode {
	avlnode.parent = parent
	return avlnode
}

func (avlnode *AVLNode) SetHeight(height int) *AVLNode {
	avlnode.height = height
	return avlnode
}

func (avlnode *AVLNode) SetBalanceFactor(balanceFactor int) *AVLNode {
	avlnode.balanceFactor = balanceFactor
	return avlnode
}

func NewAvlNode(value int) *AVLNode {
	return &AVLNode{
		value:         value,
		parent:        nil,
		left:          nil,
		right:         nil,
		height:        0,
		balanceFactor: 0,
	}
}
