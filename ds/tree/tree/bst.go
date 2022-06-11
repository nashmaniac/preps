package tree

type bst struct {
	root Node
}

// Print implements BST
func (b *bst) Print() {
	Print(b.GetRoot())
}

// Delete implements BST
func (b *bst) Delete(item int) {
	b.Find(item)
}

// Find implements BST
func (b *bst) Find(item int) Node {
	current := b.GetRoot()
	for current != nil {
		if current.GetValue() < item {
			current = current.GetRight()
		} else if current.GetValue() > item {
			current = current.GetLeft()
		} else {
			return current
		}
	}
	return nil
}

// Insert implements BST
func (b *bst) Insert(item int) {
	stack := NewStack()
	current := b.root
	for current != nil {
		stack.Push(current)
		if current.GetValue() < item {
			current = current.GetRight()
		} else {
			current = current.GetLeft()
		}
	}
	child := NewNode(item)
	for !stack.IsEmpty() {
		parent := stack.Pop()
		if parent.GetValue() < child.GetValue() {
			parent.SetRight(child)
		} else {
			parent.SetLeft(child)
		}
		child.SetParent(parent)
		child = parent
	}
	child.SetParent(nil)
	b.root = child
}

// Traverse implements BST
func (b *bst) Traverse() {
	panic("unimplemented")
}

func (b *bst) GetRoot() Node {
	return b.root
}

func (b *bst) SetRoot(node Node) {
	b.root = node
}

func NewBST() BST {
	return &bst{
		root: nil,
	}
}
