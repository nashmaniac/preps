package vector

import "fmt"

type Element interface {
	string | int | float32 | int64 | interface{}
}

type Vector interface {
	Size() int
	Capacity() int
	IsEmpty() bool
	At(index int) Element
	Push(e Element)
	Insert(index int, e Element)
	Prepend(e Element)
	Pop() Element
	Delete(index int)
	Remove(e Element)
	Find(e Element) int
	resize(new_capacity int)
	prune()
	extend()
	Print()
}

type vectorObj struct {
	Elements []Element
	capacity int
	size     int
}

// At implements Vector
func (v *vectorObj) At(index int) Element {
	if index >= v.size {
		return nil
	}
	return v.Elements[index]
}

// Capacity implements Vector
func (v *vectorObj) Capacity() int {
	return v.capacity
}

// Delete implements Vector
func (v *vectorObj) Delete(index int) {
	if index >= v.size {
		return
	}
	j := index
	for i := index + 1; i < v.size; i++ {
		v.Elements[j] = v.Elements[i]
		j++
	}
	v.size--
	v.prune()

}

// Find implements Vector
func (v *vectorObj) Find(e Element) int {
	for i := 0; i < v.size; i++ {
		if e == v.Elements[i] {
			return i
		}
	}
	return -1
}

// Insert implements Vector
func (v *vectorObj) Insert(index int, e Element) {
	if index >= v.size {
		v.Push(e)
	} else {
		j := v.size
		for i := v.size - 1; i >= index; i-- {
			v.Elements[j] = v.Elements[i]
			j--
		}
		v.Elements[index] = e
	}
	v.size++
	v.extend()
}

// IsEmpty implements Vector
func (v *vectorObj) IsEmpty() bool {
	return v.size == 0
}

// Pop implements Vector
func (v *vectorObj) Pop() Element {
	c := v.Elements[v.size-1]
	v.Elements[v.size-1] = nil
	v.size--
	v.prune()
	return c
}

// Prepend implements Vector
func (v *vectorObj) Prepend(e Element) {
	for i := v.size - 1; i >= 0; i-- {
		v.Elements[i+1] = v.Elements[i]
	}
	v.Elements[0] = e
	v.size++
	v.extend()
}

// Push implements Vector
func (v *vectorObj) Push(e Element) {
	v.Elements[v.size] = e
	v.size++
	v.extend()
}

// Remove implements Vector
func (v *vectorObj) Remove(e Element) {
	index := v.Find(e)
	if index != -1 {
		v.Delete(index)
	}
}

// Size implements Vector
func (v *vectorObj) Size() int {
	return v.size
}

// resize implements Vector
func (v *vectorObj) resize(new_capacity int) {
	temp := make([]Element, new_capacity)
	v.capacity = new_capacity
	for i := 0; i < v.size; i++ {
		temp[i] = v.Elements[i]
	}
	v.Elements = temp
}

func (v *vectorObj) prune() {
	if v.size < v.capacity/2 {
		newSize := v.capacity / 2
		v.resize(newSize)
	}
}

func (v *vectorObj) extend() {
	if v.size > v.capacity/2 {
		newSize := v.size * 2
		v.resize(newSize)
	}
}

func (v *vectorObj) Print() {
	fmt.Printf("Capacity: %d\n", v.capacity)
	fmt.Printf("Size: %d\n", v.size)
	fmt.Printf("Elements: %+v\n", v.Elements)
}

// NewVector returns a vector instance with the size initialized
func NewVector[e Element](size int) Vector {
	var elements = make([]Element, size)
	return &vectorObj{
		Elements: elements,
		capacity: size,
		size:     0,
	}
}
