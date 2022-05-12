package bbst

import (
	"github.com/nashmaniac/golang-coding/datastructures/models"
	"github.com/nashmaniac/golang-coding/datastructures/utils"
)

type AVLTree interface {
	Add(value int)
	GetRoot() models.Node
	Delete(value int)
	Search(value int) models.Node
}

type avlTree struct {
	Root models.Node
}

func (a *avlTree) Search(value int) models.Node {
	node := a.Root
	for node != nil {
		if value == node.GetValue() {
			break
		} else if value < node.GetValue() {
			node = node.GetLeft()
		} else {
			node = node.GetRight()
		}
	}
	return node
}

func Transplant(u models.Node, v models.Node) models.Node {
	parent := u.GetParent()

	if parent != nil {
		if parent.GetRight() == u {
			parent.SetRight(v)
		}

		if parent.GetLeft() == u {
			parent.SetLeft(v)
		}
	}
	if v != nil {
		v.SetParent(parent)
	}
	return v
}

func MinimumNode(node models.Node) models.Node {
	if node == nil {
		return node
	}
	for node.GetLeft() != nil {
		node = node.GetLeft()
	}
	return node
}

func MaximumNode(node models.Node) models.Node {
	if node == nil {
		return node
	}
	for node.GetRight() != nil {
		node = node.GetRight()
	}
	return node
}

func Delete(node models.Node, value int) models.Node {
	if node == nil {
		return nil
	}
	if value == node.GetValue() {
		if node.GetLeft() == nil {
			node = Transplant(node, node.GetRight())
		} else if node.GetRight() == nil {
			node = Transplant(node, node.GetLeft())
		} else {
			if node.GetLeft().GetHeight() > node.GetRight().GetHeight() {
				// left is longer
				next := MaximumNode(node.GetLeft())
				node.SetValue(next.GetValue())
				child := Delete(node.GetLeft(), next.GetValue())
				node.SetLeft(child)
				if child != nil {
					child.SetParent(node)
				}
			} else {
				next := MinimumNode(node.GetRight())
				node.SetValue(next.GetValue())
				child := Delete(node.GetRight(), next.GetValue())
				node.SetRight(child)
				if child != nil {
					child.SetParent(node)
				}
			}
		}
	} else if value < node.GetValue() {
		child := Delete(node.GetLeft(), value)
		node.SetLeft(child)
		if child != nil {
			child.SetParent(node)
		}
	} else {
		child := Delete(node.GetRight(), value)
		node.SetRight(child)
		if child != nil {
			child.SetParent(node)
		}
	}

	Update(node)
	return BalanceNode(node)
}

func (a *avlTree) Delete(value int) {
	a.Root = Delete(a.Root, value)
}

func (a *avlTree) GetRoot() models.Node {
	return a.Root
}

func Height(node models.Node) int {
	if node == nil {
		return 0
	}
	return 1 + utils.Max(Height(node.GetLeft()), Height(node.GetRight()))
}

func Update(node models.Node) {
	if node == nil {
		return
	}
	lh := -1
	rh := -1

	if node.GetLeft() != nil {
		lh = Height(node.GetLeft())
	}

	if node.GetRight() != nil {
		rh = Height(node.GetRight())
	}
	height := 1 + utils.Max(lh, rh)
	node.SetHeight(height)
	node.SetBalanceFactor(rh - lh)
}

func LeftRotation(A models.Node) models.Node {
	B := A.GetRight()

	A.SetRight(B.GetLeft())
	if B.GetLeft() != nil {
		B.GetLeft().SetParent(A)
	}

	B.SetLeft(A)
	A.SetParent(B)

	Update(A)
	Update(B)

	return B
}

func RightRotation(A models.Node) models.Node {
	B := A.GetLeft()

	A.SetLeft(B.GetRight())
	if B.GetRight() != nil {
		B.GetRight().SetParent(A)
	}

	B.SetRight(A)
	A.SetParent(B)

	Update(A)
	Update(B)

	return B
}

func BalanceNode(node models.Node) models.Node {
	if node == nil {
		return nil
	}

	if node.GetBalanceFactor() == -2 {
		if node.GetLeft().GetBalanceFactor() <= 0 {
			return RightRotation(node)
		} else {
			node.SetLeft(LeftRotation(node.GetLeft()))
			return RightRotation(node)
		}
	}

	if node.GetBalanceFactor() == 2 {
		if node.GetRight().GetBalanceFactor() >= 0 {
			return LeftRotation(node)
		} else {
			node.SetRight(RightRotation(node.GetRight()))
			return LeftRotation(node)
		}
	}

	return node
}

func Add(node models.Node, value int) models.Node {
	if node == nil {
		node = models.NewTreeNode(value)
	} else if value <= node.GetValue() {
		child := Add(node.GetLeft(), value)
		node.SetLeft(child)
		child.SetParent(node)
	} else {
		child := Add(node.GetRight(), value)
		node.SetRight(child)
		child.SetParent(node)
	}
	Update(node)
	return BalanceNode(node)
}

func (a *avlTree) Add(value int) {
	a.Root = Add(a.Root, value)
}

func NewAVL() AVLTree {
	return &avlTree{
		Root: nil,
	}
}
