package avl

type avlTree struct {
	root *AVLNode
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Update(root *AVLNode) {
	lh := -1
	rh := -1

	if root.GetLeft() != nil {
		lh = root.GetLeft().GetHeight()
	}
	if root.GetRight() != nil {
		rh = root.GetRight().GetHeight()
	}

	// set height here
	root.SetHeight(1 + max(lh, rh))
	// set balance factor here
	root.SetBalanceFactor(rh - lh)
}

func RotateRight(A *AVLNode) *AVLNode {
	/***
			 A
			/
		   B
		  /
		 C
	**/
	B := A.GetLeft()
	A.SetLeft(B.GetRight())
	if B.GetRight() != nil {
		B.GetRight().SetParent(A)
	}
	A.SetParent(B)
	B.SetRight(A)

	Update(A)
	Update(B)

	return B
}

func RotateLeft(A *AVLNode) *AVLNode {
	/***
		A
		  \
		   B
		  	\
		 	 C
	**/
	B := A.GetRight()
	A.SetRight(B.GetLeft())
	if B.GetLeft() != nil {
		B.GetLeft().SetParent(A)
	}
	A.SetParent(B)
	B.SetLeft(A)

	Update(A)
	Update(B)

	return B
}

func RotateLeftRight(A *AVLNode) *AVLNode {
	/***
			A
		   /
		  B
		   \
			C
	**/
	B := A.GetLeft()
	node := RotateLeft(B)
	A.SetLeft(node)
	node.SetParent(A)
	return RotateRight(A)
}

func RotateRightLeft(A *AVLNode) *AVLNode {
	/***
		A
		  \
		   B
		  /
		 C
	**/
	B := A.GetRight()
	node := RotateRight(B)
	A.SetRight(node)
	node.SetParent(A)
	return RotateLeft(A)
}

func balance(root *AVLNode) *AVLNode {
	// balance factor is off make sure you balance the tree here
	if abs(root.GetBalanceFactor()) > 1 {
		if root.GetLeft() != nil {
			if root.GetLeft().GetBalanceFactor() == -1 {
				return RotateRight(root)
			}
			if root.GetLeft().GetBalanceFactor() == 1 {
				return RotateLeftRight(root)
			}
		}

		if root.GetRight() != nil {
			if root.GetRight().GetBalanceFactor() == 1 {
				return RotateLeft(root)
			}
			if root.GetRight().GetBalanceFactor() == -1 {
				return RotateRightLeft(root)
			}

		}
	}
	return root
}

func Insert(root *AVLNode, value int) *AVLNode {
	if root == nil {
		root = NewAvlNode(value)
	} else {
		var node *AVLNode = nil
		if value <= root.GetValue() {
			node = Insert(root.GetLeft(), value)
			root.SetLeft(node)
		} else {
			node = Insert(root.GetRight(), value)
			root.SetRight(node)
		}
		if node != nil {
			node.SetParent(root)
		}
	}
	Update(root)
	return balance(root)
}

// Insert implements AVLTree
func (avlTree *avlTree) Insert(value int) {
	root := Insert(avlTree.GetRoot(), value)
	root.SetParent(nil)
	avlTree.SetRoot(root)
}

func NewAVLTree() AVLTree {
	return &avlTree{
		root: nil,
	}
}

func (avltree *avlTree) GetRoot() *AVLNode {
	return avltree.root
}

func (avltree *avlTree) SetRoot(root *AVLNode) *avlTree {
	avltree.root = root
	return avltree
}
