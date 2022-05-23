package trees

type bst struct {
	binarySearchTree
}

func InsertBST(node Node, value int) Node {
	if node == nil {
		return NewNode(value)
	}

	if node.GetValue() > value {
		node.SetLeft(InsertBST(node.GetLeft(), value))
	}
	if node.GetValue() < value {
		node.SetRight(InsertBST(node.GetRight(), value))
	}
	return node
}

func (b *bst) Insert(value int) {
	root := InsertBST(b.GetRoot(), value)
	b.SetRoot(root)
}

func DeleteBST(node Node, value int) Node {
	if node == nil {
		return nil
	}
	if node.GetValue() > value {
		node.SetLeft(DeleteBST(node.GetLeft(), value))
	} else if node.GetValue() < value {
		node.SetRight(DeleteBST(node.GetRight(), value))
	} else {
		if node.GetLeft() == nil {
			return node.GetRight()
		} else if node.GetRight() == nil {
			return node.GetLeft()
		} else {
			maxNode := GetMaxNode(node.GetLeft())
			node.SetValue(maxNode.GetValue())
			node.SetLeft(DeleteBST(node.GetLeft(), node.GetValue()))
		}
	}
	return node
}

func (b *bst) Delete(value int) {
	root := DeleteBST(b.GetRoot(), value)
	b.SetRoot(root)
}

func NewBST() BinarySearchTree {
	b := &bst{}
	b.SetRoot(nil)
	return b
}
