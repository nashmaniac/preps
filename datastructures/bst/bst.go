package bst

type BST interface {
	Insert(value int)
	InOrderWalk() []int
	PostOrderWalk() []int
	PreOrderWalk() []int
	Search(value int) *BSTNode
	Successor(value int) *BSTNode
	Predecessor(value int) *BSTNode
	Minimum(value int) *BSTNode
	MinimumNode(node *BSTNode) *BSTNode
	MaximumNode(node *BSTNode) *BSTNode
	Delete(value int)
}
