package queue

type QueueElement interface {
	GetValue() Element
	SetValue(item Element)

	GetKey() int
	SetKey(item int)
}

type queueElement struct {
	key   int
	value Element
}

func (queueelement *queueElement) GetKey() int {
	return queueelement.key
}

func (queueelement *queueElement) GetValue() Element {
	return queueelement.value
}

func (queueelement *queueElement) SetKey(key int) {
	queueelement.key = key
}

func (queueelement *queueElement) SetValue(value Element) {
	queueelement.value = value
}

func NewQueueElement(key int, item Element) QueueElement {
	return &queueElement{
		key:   key,
		value: item,
	}
}
