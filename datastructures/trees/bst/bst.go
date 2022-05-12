package bst

import (
	"github.com/nashmaniac/golang-coding/datastructures/models"
)

type BST interface {
	Add(value int)
	GetRoot() models.Node
	Search(value int) models.Node
	Delete(value int)
}

type bst struct {
	Root models.Node
}

// GetRoot implements BST
func (b *bst) GetRoot() models.Node {
	return b.Root
}

func Search(node models.Node, value int) models.Node {
	if node == nil {
		return nil
	}
	if value == node.GetValue() {
		return node
	} else if value < node.GetValue() {
		return Search(node.GetLeft(), value)
	} else {
		return Search(node.GetRight(), value)
	}
}

func (b *bst) Search(value int) models.Node {
	return Search(b.GetRoot(), value)
}

func Transplant(u models.Node, v models.Node) models.Node {
	parent := u.GetParent()
	if parent != nil {
		if u == parent.GetLeft() {
			parent.SetLeft(v)
		}
		if u == parent.GetRight() {
			parent.SetRight(v)
		}
	}
	if v != nil {
		v.SetParent(parent)
	}

	return v
}

func MinimumNode(node models.Node) models.Node {
	for node.GetLeft() != nil {
		node = node.GetLeft()
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
			rightMinimum := MinimumNode(node.GetRight())
			node.SetValue(rightMinimum.GetValue())
			node.SetRight(Delete(node.GetRight(), rightMinimum.GetValue()))
		}
	} else if value < node.GetValue() {
		node.SetLeft(Delete(node.GetLeft(), value))
	} else {
		node.SetRight(Delete(node.GetRight(), value))
	}
	return node
}

func (b *bst) Delete(value int) {
	b.Root = Delete(b.Root, value)
}

func Add(node models.Node, value int) models.Node {
	if node == nil {
		return models.NewTreeNode(value)
	}
	if value <= node.GetValue() {
		child := Add(node.GetLeft(), value)
		node.SetLeft(child)
		child.SetParent(node)
	} else {
		child := Add(node.GetRight(), value)
		node.SetRight(child)
		child.SetParent(node)
	}
	return node
}

// Add implements BST
func (b *bst) Add(value int) {
	b.Root = Add(b.Root, value)
}

func NewBST() BST {
	return &bst{
		Root: nil,
	}
}
