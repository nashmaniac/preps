package queue

type Element interface {
	string | int | float32 | interface{}
}

type Queue interface {
	Enqueue(e Element)
	Dequeue() Element
	IsEmpty() bool
}

type queue struct {
	Elements []Element
}

// IsEmpty implements Queue
func (q *queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

// Dequeue implements Queue
func (q *queue) Dequeue() Element {
	e := q.Elements[0]
	q.Elements = q.Elements[1:]
	return e
}

// Enqueue implements Queue
func (q *queue) Enqueue(e Element) {
	q.Elements = append(q.Elements, e)
}

func NewQueue[e Element]() Queue {
	elements := make([]Element, 0)
	return &queue{
		Elements: elements,
	}
}
