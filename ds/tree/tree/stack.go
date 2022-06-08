package tree

type StackNode interface {
	GetNode() Node
	SetNode(node Node)

	SetPrev(s StackNode)
	GetPrev() StackNode

	SetNext(s StackNode)
	GetNext() StackNode
}

type stackNode struct {
	node Node
	next StackNode
	prev StackNode
}

// GetNext implements StackNode
func (s *stackNode) GetNext() StackNode {
	return s.next
}

// GetNode implements StackNode
func (s *stackNode) GetNode() Node {
	return s.node
}

// GetPrev implements StackNode
func (s *stackNode) GetPrev() StackNode {
	return s.prev
}

// SetNext implements StackNode
func (s *stackNode) SetNext(node StackNode) {
	s.next = s
}

// SetNode implements StackNode
func (s *stackNode) SetNode(node Node) {
	s.node = node
}

// SetPrev implements StackNode
func (s *stackNode) SetPrev(node StackNode) {
	s.prev = node
}

func NewStackNode(node Node) StackNode {
	return &stackNode{
		node: node,
		next: nil,
		prev: nil,
	}
}

type Stack interface {
	IsEmpty() bool
	Push(item Node)
	Pop() Node
}

type stack struct {
	head StackNode
	tail StackNode
	size int
}

// IsEmpty implements Stack
func (s *stack) IsEmpty() bool {
	return s.size == 0
}

// Pop implements Stack
func (s *stack) Pop() Node {
	if s.IsEmpty() {
		return nil
	}
	tail := s.tail
	newTail := tail.GetPrev()
	if newTail != nil {
		newTail.SetNext(nil)
	}
	s.tail = newTail
	s.size--
	return tail.GetNode()
}

// Push implements Stack
func (s *stack) Push(item Node) {
	node := NewStackNode(item)
	if s.IsEmpty() {
		s.tail = node
		s.head = node
	} else {
		node.SetPrev(s.tail)
		s.tail.SetNext(node)
		s.tail = node
	}
	s.size++
}

func NewStack() Stack {
	return &stack{
		head: nil,
		tail: nil,
		size: 0,
	}
}
