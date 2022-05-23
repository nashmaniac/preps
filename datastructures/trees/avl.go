package trees

type avlTree struct {
	binarySearchTree
}

func updateHeight(node Node) {
	if node == nil {
		return
	}

	lh := GetHeight(node.GetLeft())
	rh := GetHeight(node.GetRight())
	h := max(lh, rh)
	bf := rh - lh
	node.SetHeight(h)
	node.SetBalanceFactor(bf)
}

func LeftRotation(node Node) Node {
	x := node
	y := node.GetRight()

	t2 := y.GetLeft()
	x.SetRight(t2)
	y.SetLeft(x)

	updateHeight(x)
	updateHeight(y)
	return y
}

func RightRotation(node Node) Node {
	x := node
	y := node.GetLeft()

	t2 := y.GetRight()
	x.SetLeft(t2)
	y.SetRight(x)
	updateHeight(x)
	updateHeight(y)
	return y
}

func balanceNode(node Node) Node {
	if node.GetBalanceFactor() > 1 {
		// right heavy
		if node.GetRight().GetBalanceFactor() > 0 {
			// right right heavy
			return LeftRotation(node)
		} else {
			// right left heavy
			node.SetRight(RightRotation(node.GetRight()))
			return LeftRotation(node)
		}
	}

	if node.GetBalanceFactor() < -1 {
		// left heavy
		if node.GetLeft().GetBalanceFactor() < 0 {
			// left left heavy
			return RightRotation(node)
		} else {
			// left right heavy
			node.SetLeft(LeftRotation(node.GetLeft()))
			return RightRotation(node)
		}
	}
	return node
}

func InsertAVL(node Node, value int) Node {
	if node == nil {
		node = NewNode(value)
	}

	if node.GetValue() < value {
		node.SetRight(InsertAVL(node.GetRight(), value))
	}
	if node.GetValue() > value {
		node.SetLeft(InsertAVL(node.GetLeft(), value))
	}

	updateHeight(node)
	return balanceNode(node)
}

func (a *avlTree) Insert(value int) {
	root := InsertAVL(a.GetRoot(), value)
	a.SetRoot(root)
}

func DeleteAVL(node Node, value int) Node {
	if node == nil {
		return nil
	}
	if node.GetValue() > value {
		node.SetLeft(DeleteAVL(node.GetLeft(), value))
	} else if node.GetValue() < value {
		node.SetRight(DeleteAVL(node.GetRight(), value))
	} else {
		if node.GetLeft() == nil {
			return node.GetRight()
		} else if node.GetRight() == nil {
			return node.GetLeft()
		} else {
			maxNode := GetMaxNode(node.GetLeft())
			node.SetValue(maxNode.GetValue())
			node.SetLeft(DeleteAVL(node.GetLeft(), node.GetValue()))
		}
	}
	updateHeight(node)
	return balanceNode(node)
}

func (b *avlTree) Delete(value int) {
	root := DeleteAVL(b.GetRoot(), value)
	b.SetRoot(root)
}

func NewAVLTree() BinarySearchTree {
	a := &avlTree{}
	a.SetRoot(nil)
	return a
}
