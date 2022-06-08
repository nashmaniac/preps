package tree

type QueueElement interface {
	GetNode() Node
	SetNode(node Node)

	GetPrev() QueueElement
	SetPrev(node QueueElement)

	GetNext() QueueElement
	SetNext(node QueueElement)
}

type queueElement struct {
	node Node
	next QueueElement
	prev QueueElement
}

// GetNext implements QueueElement
func (q *queueElement) GetNext() QueueElement {
	return q.next
}

// GetNode implements QueueElement
func (q *queueElement) GetNode() Node {
	return q.node
}

// GetPrev implements QueueElement
func (q *queueElement) GetPrev() QueueElement {
	return q.prev
}

// SetNext implements QueueElement
func (q *queueElement) SetNext(node QueueElement) {
	q.next = node
}

// SetNode implements QueueElement
func (q *queueElement) SetNode(node Node) {
	q.node = node
}

// SetPrev implements QueueElement
func (q *queueElement) SetPrev(node QueueElement) {
	q.prev = node
}

func NewQueueElement(node Node) QueueElement {
	return &queueElement{
		node: node,
		next: nil,
		prev: nil,
	}
}

type Queue interface {
	IsEmpty() bool
	Offer(node Node)
	Pop() Node
}

type queue struct {
	head QueueElement
	tail QueueElement
	size int
}

// IsEmpty implements Queue
func (q *queue) IsEmpty() bool {
	return q.size == 0
}

// Offer implements Queue
func (q *queue) Offer(node Node) {
	qElement := NewQueueElement(node)
	if q.IsEmpty() {
		q.head = qElement
		q.tail = qElement
	} else {
		qElement.SetPrev(q.tail)
		q.tail.SetNext(qElement)
		q.tail = qElement
	}
	q.size++
}

// Pop implements Queue
func (q *queue) Pop() Node {
	if q.IsEmpty() {
		return nil
	}
	head := q.head
	next := head.GetNext()
	if next != nil {
		next.SetPrev(nil)
	}
	q.head = next
	q.size--
	return head.GetNode()
}

func NewQueue[e Node]() Queue {
	return &queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}
