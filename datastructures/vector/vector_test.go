package vector_test

import (
	"github.com/nashmaniac/golang-coding/datastructures/vector"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vector Test", func() {
	Context("Initialization", func() {
		When("When initialized", func() {
			It("should match the capacity to given number and size 0", func() {
				n := 1
				v := vector.NewVector[string](n)
				Expect(v.Capacity()).To(Equal(n))
				Expect(v.Size()).To(Equal(0))
			})
		})
	})

	Context("Push", func() {
		When("the size is more than the capacity", func() {
			It("should increase the capacity to double the size", func() {
				n := 1
				v := vector.NewVector[string](n)
				v.Push("Hello")
				// v.Print()
				v.Push("World")
				Expect(v.Size()).To(Equal(2))
				Expect(v.Capacity()).To(Equal(4))
				// v.Print()
			})
		})
	})

	Context("Prepend", func() {
		When("an object is pushed", func() {
			It("the size should match", func() {
				n := 1
				v := vector.NewVector[string](n)
				v.Prepend("Hello")
				v.Push("World")
				Expect(v.Size()).To(Equal(2))
				Expect(v.Capacity()).To(Equal(4))
			})
		})

		When("an object is pushed", func() {
			It("the size should match", func() {
				n := 3
				v := vector.NewVector[string](n)
				v.Push("Hello")
				v.Push("World")
				v.Push("Shetu")
				v.Prepend("Hola")
				v.Print()
				Expect(v.Size()).To(Equal(4))
				Expect(v.Capacity()).To(Equal(8))
			})
		})
	})

	Context("Pop", func() {
		When("an object is pop", func() {
			It("the size should match", func() {
				n := 4
				v := vector.NewVector[string](n)
				v.Prepend("Hello")
				// v.Print()
				v.Push("World")
				// v.Print()
				c := v.Pop()
				// v.Print()
				Expect(c).To(Equal("World"))
				Expect(v.Size()).To(Equal(1))
				Expect(v.Capacity()).To(Equal(2))
			})
		})
	})
	Context("Delete", func() {
		When("an object is deleted", func() {
			It("the size should match", func() {
				n := 4
				v := vector.NewVector[string](n)
				v.Prepend("Hello")
				// v.Print()
				v.Push("World")
				// v.Print()
				v.Delete(1)
				// v.Print()
				Expect(v.Size()).To(Equal(1))
				Expect(v.Capacity()).To(Equal(2))
			})
		})

		When("an object is deleted", func() {
			It("the size should match", func() {
				n := 4
				v := vector.NewVector[string](n)
				v.Push("Hello1")
				v.Push("Hello2")
				v.Push("Hello3")
				v.Push("Hello4")
				v.Delete(1)
				v.Print()
				Expect(v.Size()).To(Equal(3))
				Expect(v.Capacity()).To(Equal(4))
			})
		})

		When("an object is pop", func() {
			It("the size should match", func() {
				n := 4
				v := vector.NewVector[string](n)
				v.Prepend("Hello")
				// v.Print()
				v.Push("World")
				// v.Print()
				v.Delete(3)
				// v.Print()
				Expect(v.Size()).To(Equal(2))
				Expect(v.Capacity()).To(Equal(4))
			})
		})
	})

	Context("At", func() {
		When("index is less than the size", func() {
			It("passes", func() {
				n := 4
				v := vector.NewVector[string](n)
				v.Push("Hello")
				v.Prepend("Hello1")
				v.Prepend("Hello2")
				c := v.At(1)
				Expect(c).To(Equal("Hello1"))
			})

		})

		When("index is greater than the size", func() {
			It("passes", func() {
				n := 4
				v := vector.NewVector[string](n)
				v.Push("Hello")
				v.Prepend("Hello1")
				v.Prepend("Hello2")
				c := v.At(5)
				Expect(c).To(BeNil())
			})

		})

	})

	Context("Find", func() {
		When("index is less than the size", func() {
			It("passes", func() {
				n := 4
				v := vector.NewVector[string](n)
				v.Push("Hello")
				v.Prepend("Hello1")
				v.Prepend("Hello2")
				c := v.Find("Hello")
				Expect(c).To(Equal(2))
			})

		})

		When("index is greater than the size", func() {
			It("passes", func() {
				n := 4
				v := vector.NewVector[string](n)
				v.Push("Hello")
				v.Prepend("Hello1")
				v.Prepend("Hello2")
				c := v.Find("Hello4")
				Expect(c).To(Equal(-1))
			})

		})

	})

	Context("IsEmpty", func() {
		When("all the objects are removed", func() {
			It("should be empty", func() {
				n := 2
				v := vector.NewVector[string](n)
				Expect(v.IsEmpty()).To(BeTrue())
				v.Push("Hello")
				v.Push("World")
				v.Delete(0)
				v.Delete(0)
				Expect(v.IsEmpty()).To(BeTrue())
			})
		})
		When("all the objects are not removed", func() {
			It("should be empty", func() {
				n := 2
				v := vector.NewVector[string](n)
				Expect(v.IsEmpty()).To(BeTrue())
				v.Push("Hello")
				v.Push("World")
				Expect(v.IsEmpty()).ToNot(BeTrue())
			})
		})
	})

	Context("Insert", func() {
		When("when the index is within range", func() {
			It("should add the element", func() {
				n := 6
				v := vector.NewVector[string](n)
				v.Push("hello")
				v.Push("world")
				v.Insert(1, "shetu")
				Expect(v.Find("shetu")).To(Equal(1))
			})
		})

		When("when the index is outside range", func() {
			It("should add the element", func() {
				n := 3
				v := vector.NewVector[string](n)
				v.Push("hello")
				v.Push("world")
				v.Insert(5, "shetu")
				Expect(v.Find("shetu")).To(Equal(2))
			})
		})
	})

	Context("remove", func() {
		When("element is in the vector", func() {
			It("should remove", func() {
				n := 3
				v := vector.NewVector[string](n)
				v.Push("hello")
				v.Push("world")
				v.Push("dude")
				v.Remove("world")
				Expect(v.Size()).To(Equal(2))
			})
		})
		When("element is not in the vector", func() {
			It("should remove", func() {
				n := 3
				v := vector.NewVector[string](n)
				v.Push("hello")
				v.Push("world")
				v.Push("dude")
				v.Remove("world2")
				Expect(v.Size()).To(Equal(3))
			})
		})
	})
})
