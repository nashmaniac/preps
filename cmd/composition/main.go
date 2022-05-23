package main

import "fmt"

type Animal interface {
	Bark()
}

func (a *animal) Bark() {
	fmt.Println("default bark")
}

type animal struct {
	Id int
}

type Dog struct {
	animal
}

func (d *Dog) Bark() {
	fmt.Println("Gheo", d.Id)
}

type Cat struct {
	animal
}

func (c *Cat) Bark() {
	fmt.Println("meow", c.Id)
}

type Mouse struct {
	animal
}

func (m *Mouse) Bark() {
	fmt.Println("chui chui")
}

func main() {
	a := animal{}
	a.Bark()

	d := Dog{}
	d.Bark()

	c := Cat{}
	c.Bark()

	m := Mouse{}
	m.Bark()
}
