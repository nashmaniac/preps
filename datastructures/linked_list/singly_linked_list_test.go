package linkedlist_test

import (
	linkedlist "github.com/nashmaniac/golang-coding/datastructures/linked_list"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SinglyLinkedList", func() {
	Context("Size", func() {
		When("initialized", func() {
			It("size is 0", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				Expect(l.Size()).To(Equal(0))
			})
		})
		When("added element", func() {
			It("size is 1", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushBack("Hello")
				Expect(l.Size()).To(Equal(1))
			})
		})
	})

	Context("Empty", func() {
		When("initialized", func() {
			It("size is 0", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				Expect(l.Empty()).To(BeTrue())
			})
		})
		When("added element", func() {
			It("size is 1", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushBack("Hello")
				Expect(l.Empty()).ToNot(BeTrue())
			})
		})
	})

	Context("PushBack", func() {
		When("the list is empty", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushBack("hello")
				Expect(l.Size()).To(Equal(1))
			})
		})

		When("the list is not empty", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushBack("hello")
				l.PushBack("world")
				Expect(l.Size()).To(Equal(2))
			})
		})
	})

	Context("PushFront", func() {
		When("the list is empty", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushFront("hello")
				Expect(l.Size()).To(Equal(1))
				Expect(l.Front()).To(Equal("hello"))
			})
		})

		When("the list is not empty", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushFront("hello")
				l.PushFront("world")
				Expect(l.Size()).To(Equal(2))
				Expect(l.Front()).To(Equal("world"))
			})
		})
	})

	Context("Front", func() {
		When("the list is empty", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				Expect(l.Size()).To(Equal(0))
				Expect(l.Front()).To(BeNil())
			})
		})

		When("the list is not empty", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushFront("hello")
				l.PushFront("world")
				Expect(l.Size()).To(Equal(2))
				Expect(l.Front()).To(Equal("world"))
			})
		})
	})

	Context("Back", func() {
		When("the list is empty", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				Expect(l.Size()).To(Equal(0))
				Expect(l.Back()).To(BeNil())
			})
		})

		When("the list is not empty", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushFront("hello")
				l.PushFront("world")
				Expect(l.Size()).To(Equal(2))
				Expect(l.Back()).To(Equal("hello"))
			})
		})
	})

	Context("PopFront", func() {

		When("the list contains no item", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				value := l.PopFront()
				Expect(l.Size()).To(Equal(0))
				Expect(value).To(BeNil())
			})
		})

		When("the list contains only item", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushBack("Hello")
				value := l.PopFront()
				Expect(l.Size()).To(Equal(0))
				Expect(value).To(Equal("Hello"))
			})
		})

		When("the list contains more than one item", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushFront("hello")
				l.PushFront("world")
				value := l.PopFront()
				Expect(l.Size()).To(Equal(1))
				Expect(value).To(Equal("world"))
				value = l.PopFront()
				Expect(l.Size()).To(Equal(0))
				Expect(value).To(Equal("hello"))

			})
		})
	})

	Context("PopBack", func() {

		When("the list contains no item", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				value := l.PopBack()
				Expect(l.Size()).To(Equal(0))
				Expect(value).To(BeNil())
			})
		})

		When("the list contains only item", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushBack("Hello")
				value := l.PopBack()
				Expect(l.Size()).To(Equal(0))
				Expect(value).To(Equal("Hello"))
			})
		})

		When("the list contains more than one item", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[string]()
				l.PushFront("hello")
				l.PushFront("world")
				value := l.PopBack()
				Expect(l.Size()).To(Equal(1))
				Expect(value).To(Equal("hello"))
				value = l.PopBack()
				Expect(l.Size()).To(Equal(0))
				Expect(value).To(Equal("world"))

			})
		})
	})

	Context("RemoveValue", func() {
		When("the element is in the list", func() {
			It("passes", func() {
				l := linkedlist.NewSinglyLinkedList[int]()
				l.PushBack(1)
				l.PushBack(2)
				l.PushBack(3)
				l.PushBack(4)
				
			})
		})


	})

})
