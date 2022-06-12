package tree

type avl struct {
	root Node
}

func updateNode(node Node) {
	lh := -1
	rh := -1
	if node.GetLeft() != nil {
		lh = node.GetLeft().GetHeight()
	} else {
		lh = 0
	}
	if node.GetRight() != nil {
		rh = node.GetRight().GetHeight()
	} else {
		rh = 0
	}
	h := 1 + Max(lh, rh)
	node.SetHeight(h)
	node.SetBalanceFactor(rh - lh)
}

func RotateLeft(node Node) Node {
	B := node.GetRight()
	t2 := B.GetLeft()

	node.SetRight(t2)
	if t2 != nil {
		t2.SetParent(node)
	}

	B.SetLeft(node)
	node.SetParent(B)

	updateNode(node)
	updateNode(B)

	return B
}

func RotateRight(node Node) Node {
	B := node.GetLeft()
	t3 := B.GetRight()

	node.SetLeft(t3)
	if t3 != nil {
		t3.SetParent(node)
	}

	B.SetRight(node)
	node.SetParent(B)

	updateNode(node)
	updateNode(B)

	return B
}
func rebalanceNode(node Node) Node {
	updateNode(node)
	if Abs(node.GetBalanceFactor()) == 2 {
		if node.GetRight().GetBalanceFactor() == 1 {
			// right heavy
			return RotateLeft(node)
		} else if node.GetRight().GetBalanceFactor() == -1 {
			// right left heavy
			rightNode := RotateRight(node.GetRight())
			rightNode.SetParent(node)
			node.SetRight(rightNode)
			return RotateLeft(node)
		} else if node.GetLeft().GetBalanceFactor() == 1 {
			// left right heavy
			leftNode := RotateLeft(node.GetLeft())
			leftNode.SetParent(node)
			node.SetLeft(leftNode)
			return RotateRight(leftNode)
		} else if node.GetLeft().GetBalanceFactor() == -1 {
			// left left heavy
			return RotateRight(node)
		} else {
			
		}
	}
	return node
}

func Insert(node Node, item int) Node {
	if node == nil {
		return NewNode(item)
	} else if node.GetValue() > item {
		child := Insert(node.GetLeft(), item)
		node.SetLeft(child)
		child.SetParent(node)
	} else {
		child := Insert(node.GetRight(), item)
		node.SetRight(child)
		child.SetParent(node)
	}
	return rebalanceNode(node)
}

func (a *avl) Insert(item int) {
	a.SetRoot(Insert(a.GetRoot(), item))
}

func MinimumNode(node Node) Node {
	for node.GetLeft() != nil {
		node = node.GetLeft()
	}
	return node
}
func Delete(node Node, item int) Node {
	if node == nil {
		return nil
	} else if node.GetValue() < item {
		rightNode := Delete(node.GetRight(), item)
		if rightNode != nil {
			rightNode.SetParent(node)
		}
		node.SetRight(rightNode)
		return rebalanceNode(node)
	} else if node.GetValue() > item {
		leftNode := Delete(node.GetLeft(), item)
		if leftNode != nil {
			leftNode.SetParent(node)
		}
		node.SetLeft(leftNode)
		return rebalanceNode(node)
	} else {
		if node.GetLeft() == nil {
			return node.GetRight()
		} else if node.GetRight() == nil {
			return node.GetLeft()
		} else {
			minNode := MinimumNode(node.GetRight())
			node.SetValue(minNode.GetValue())
			rightNode := Delete(node.GetRight(), minNode.GetValue())
			node.SetRight(rightNode)
			if rightNode != nil {
				rightNode.SetParent(node)
			}
			return node
		}
	}
}

func (a *avl) Delete(item int) {
	a.Print()
	a.SetRoot(Delete(a.GetRoot(), item))
	a.Print()
}

func (a *avl) Find(item int) Node {
	//TODO implement me
	panic("implement me")
}

func (a *avl) Traverse() {
	//TODO implement me
	panic("implement me")
}

func (a *avl) Print() {
	Print(a.GetRoot())
}

func (a *avl) GetRoot() Node {
	return a.root
}

func (a *avl) SetRoot(node Node) {
	a.root = node
}

func NewAVl() BST {
	return &avl{
		root: nil,
	}
}
