package vector

import (
	"fmt"
)

// integer vector
type Vector interface {
	Size() int
	Capacity() int
	IsEmpty() bool
	Add(item int)
	Find(item int) int
	Search(item int) int
	Delete(index int)
	Remove(item int)
	At(index int) int
	Insert(index int, item int)
	Elements() []int
	Print()
}

type vector struct {
	elements []int
	size     int
	capacity int
}

// Add implements Vector
func (v *vector) Add(item int) {
	v.Insert(v.size, item)
}

// At implements Vector
func (v *vector) At(index int) int {
	if index < 0 || index >= v.size {
		return -1
	}
	return v.elements[index]
}

// Capacity implements Vector
func (v *vector) Capacity() int {
	return v.capacity
}

// Delete implements Vector
func (v *vector) Delete(index int) {
	if index < 0 || index >= v.size {
		return
	}
	for i := index + 1; i < v.size; i++ {
		v.elements[i-1] = v.elements[i]
	}
	v.size--
	// shrink now
	if v.size < v.capacity/2 {
		// fmt.Println("shirking")
		v.capacity = v.capacity / 2
		elements := make([]int, v.capacity)
		for i := 0; i < v.size; i++ {
			elements[i] = v.elements[i]
		}
		v.elements = elements
	}
}

// Elements implements Vector
func (v *vector) Elements() []int {
	elements := make([]int, 0)
	for i := 0; i < v.size; i++ {
		elements = append(elements, v.elements[i])
	}
	return elements
}

// Find implements Vector
func (v *vector) Find(item int) int {
	for i := 0; i < v.size; i++ {
		if v.elements[i] == item {
			return i
		}
	}
	return -1
}

// Insert implements Vector
func (v *vector) Insert(index int, item int) {
	capacity := v.capacity
	if v.size == 0 {
		capacity = 1
	} else if v.size+1 >= v.capacity {
		capacity = 2 * v.size
	}

	if capacity != v.capacity {
		elements := make([]int, capacity)
		for i := 0; i < v.size; i++ {
			elements[i] = v.elements[i]
		}
		for i := v.size - 1; i >= index; i-- {
			elements[i+1] = v.elements[i]
		}
		v.capacity = capacity
		elements[index] = item
		v.elements = elements
	} else {
		v.elements[index] = item
	}
	v.size++
	// v.Print()
}

// IsEmpty implements Vector
func (v *vector) IsEmpty() bool {
	return v.size == 0
}

// Print implements Vector
func (v *vector) Print() {
	elements := v.Elements()
	fmt.Printf("Cap: %d Size: %d Elements: %+v\n", v.capacity, v.size, elements)
}

// Remove implements Vector
func (v *vector) Remove(item int) {
	index := v.Find(item)
	if index != -1 {
		v.Delete(index)
	}
}

// Search implements Vector
func (v *vector) Search(item int) int {
	return v.Find(item)
}

// Size implements Vector
func (v *vector) Size() int {
	return v.size
}

func NewVector(n int) Vector {
	return &vector{
		size:     0,
		capacity: n,
		elements: make([]int, n),
	}
}
