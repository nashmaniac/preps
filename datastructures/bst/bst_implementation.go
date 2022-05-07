package bst

type bstTree struct {
	Root *BSTNode
}

// MaximumNode implements BST
func (*bstTree) MaximumNode(node *BSTNode) *BSTNode {
	current := node.GetRight()
	for current.GetRight() != nil {
		current = current.GetRight()
	}
	return current
}

// MinimumNode implements BST
func (*bstTree) MinimumNode(node *BSTNode) *BSTNode {
	current := node.GetLeft()
	for current.GetLeft() != nil {
		current = current.GetLeft()
	}
	return current
}

func (b *bstTree) Transplant(u *BSTNode, v *BSTNode) {
	if u.GetParent() == nil {
		b.Root = v
	} else if u == u.GetParent().GetLeft() {
		u.GetParent().SetLeft(v)
	} else {
		u.GetParent().SetRight(v)
	}

	if u.GetParent() != nil {
		v.SetParent(u.GetParent())
	}
}

// Delete implements BST
func (b *bstTree) Delete(value int) {
	node := b.Search(value)
	if node == nil {
		return
	}

	if node.GetLeft() == nil {
		b.Transplant(node, node.GetRight())
	} else if node.GetRight() == nil {
		b.Transplant(node, node.GetLeft())
	} else {
		y := b.MinimumNode(node.GetRight())
		if y.GetParent() != node {
			b.Transplant(y, y.GetRight())
			y.SetRight(node.GetRight())
			node.GetRight().SetParent(y)
		}
		b.Transplant(node, y)
		y.SetLeft(node.GetLeft())
		node.GetLeft().SetParent(y)
	}
}

func inOrderWalk(root *BSTNode, values []int) {
	if root == nil {
		return
	}
	inOrderWalk(root.GetLeft(), values)
	values = append(values, root.GetValue())
	inOrderWalk(root.GetRight(), values)
}

func preOrderWalk(root *BSTNode, values []int) {
	if root == nil {
		return
	}
	values = append(values, root.GetValue())
	preOrderWalk(root.GetLeft(), values)
	preOrderWalk(root.GetRight(), values)
}

func postOrderWalk(root *BSTNode, values []int) {
	if root == nil {
		return
	}
	postOrderWalk(root.GetLeft(), values)
	postOrderWalk(root.GetRight(), values)
	_ = append(values, root.GetValue())

}

// InOrderWalk implements BST
func (b *bstTree) InOrderWalk() []int {
	values := make([]int, 0)
	inOrderWalk(b.Root, values)
	return values
}

// Insert implements BST
func (b *bstTree) Insert(value int) {
	var y *BSTNode = nil
	x := b.Root
	node := NewBSTNode(value)
	for x != nil {
		y = x
		if x.GetValue() <= value {
			x = x.GetLeft()
		} else {
			x = x.GetRight()
		}
	}

	if y == nil {
		b.Root = node
	} else if y.GetValue() <= value {
		y.SetLeft(node)
	} else {
		y.SetRight(node)
	}
	node.SetParent(y)

}

// Maximum implements BST
func (b *bstTree) Maximum(value int) *BSTNode {
	node := b.Search(value)
	if node != nil {
		return b.MaximumNode(node)
	}
	return node
}

// Minimum implements BST
func (b *bstTree) Minimum(value int) *BSTNode {
	node := b.Search(value)
	if node != nil {
		return b.MinimumNode(node)
	}
	return node
}

// PostOrderWalk implements BST
func (b *bstTree) PostOrderWalk() []int {
	values := make([]int, 0)
	postOrderWalk(b.Root, values)
	return values
}

// PreOrderWalk implements BST
func (b *bstTree) PreOrderWalk() []int {
	values := make([]int, 0)
	preOrderWalk(b.Root, values)
	return values
}

// Predecessor implements BST
func (b *bstTree) Predecessor(value int) *BSTNode {
	node := b.Search(value)
	if node != nil {
		return node.GetParent()
	}
	return nil
}

// Search implements BST
func (b *bstTree) Search(value int) *BSTNode {
	x := b.Root
	for x != nil {
		if x.GetValue() == value {
			return x
		} else if x.GetValue() < value {
			x = x.GetRight()
		} else {
			x = x.GetLeft()
		}
	}
	return x
}

// Successor implements BST
func (b *bstTree) Successor(value int) *BSTNode {
	node := b.Search(value)
	if node != nil {
		if node.GetRight() != nil {
			return b.MinimumNode(node.GetRight())
		} else {
			y := node.GetParent()
			for y != nil && y.GetRight() == node {
				y = y.GetParent()
				node = y
			}
			return y
		}
	}
	return nil
}

func NewBST() BST {
	return &bstTree{
		Root: nil,
	}
}
