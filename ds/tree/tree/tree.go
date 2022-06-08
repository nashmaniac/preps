package tree

type BST interface {
	Insert(item int)
	Delete(item int)
	Find(item int)
	Traverse()
	Print()
}
