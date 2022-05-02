package linkedlist_test

import (
	linkedlist "github.com/nashmaniac/golang-coding/datastructures/linked_list"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SinglyLinkedList", func() {

	Context("Size", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})

		When("initialized", func() {
			It("size should be zero", func() {
				Expect(l.Size()).To(Equal(0))
			})
		})

		When("added an element", func() {
			It("should increment the size by 1", func() {
				iLength := l.Size()
				l.PushBack("hello")
				rLength := l.Size()
				Expect(rLength - iLength).To(Equal(1))
			})
		})

		When("remove an element", func() {
			It("should decrement the size by 1", func() {
				iLength := l.Size()
				l.PopBack()
				rLength := l.Size()
				Expect(iLength - rLength).To(Equal(1))
			})
		})

	})

	Context("Empty", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})

		When("initialized", func() {
			It("return empty", func() {
				Expect(l.Empty()).To(BeTrue())
			})
		})

		When("added an element", func() {
			It("should not be empty", func() {
				l.PushBack("hello")
				Expect(l.Empty()).ToNot(BeTrue())
			})
		})

		When("remove an element from more than an element", func() {
			It("should not be empty", func() {
				l.PushBack("hello")
				l.PushBack("world")
				l.PopFront()
				Expect(l.Empty()).ToNot(BeTrue())
			})
		})

		When("remove an element with a single element", func() {
			It("should be empty", func() {
				l.PushBack("hello")
				l.PopFront()
				Expect(l.Empty()).To(BeTrue())
			})
		})

	})

	Context("PushFront", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})

		When("added an element with empty list", func() {
			It("the front should be equal to value and size increment by 1", func() {
				iLength := l.Size()
				l.PushFront("Hello")
				rLength := l.Size()
				Expect(rLength - iLength).To(Equal(1))
			})
		})

		When("added an element with a non-empty list", func() {
			BeforeEach(func() {
				l.PushFront("World")
			})

			It("the front should be equal to value and size increment by 1", func() {
				iLength := l.Size()
				l.PushFront("Hello")
				rLength := l.Size()
				Expect(rLength - iLength).To(Equal(1))
			})
		})

	})

	Context("PushBack", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})

		When("added an element with empty list", func() {
			It("the back should be equal to value and size increment by 1", func() {
				iLength := l.Size()
				l.PushBack("hello")
				Expect(l.Size() - iLength).To(Equal(1))
			})
		})

		When("added an element with a non-empty list", func() {
			BeforeEach(func() {
				l.PushBack("hello")
			})

			It("the back should be equal to value and size increment by 1", func() {
				iLength := l.Size()
				l.PushBack("world")
				Expect(l.Size() - iLength).To(Equal(1))
			})
		})

	})

	Context("Front", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})

		When("list is empty", func() {
			It("should return nil", func() {
				value := l.Front()
				Expect(value).To(BeNil())
			})
		})

		When("list is not empty", func() {
			It("should return the first value", func() {
				l.PushBack("hello")
				value := l.Front()
				Expect(value).To(Equal("hello"))
			})
		})

	})

	Context("Back", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})

		When("list is empty", func() {
			It("should return nil", func() {
				value := l.Back()
				Expect(value).To(BeNil())
			})
		})

		When("list is not empty", func() {
			It("should return the last value", func() {
				l.PushBack("hello")
				l.PushBack("world")
				value := l.Back()
				Expect(value).To(Equal("world"))
			})
		})

	})

	Context("PopFront", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})

		When("list is empty", func() {
			It("should return nil", func() {
				value := l.PopFront()
				Expect(value).To(BeNil())
			})
		})

		When("list is not empty and has one item", func() {
			It("should remove the first value", func() {
				l.PushBack("hello")
				value := l.PopFront()
				Expect(value).To(Equal("hello"))
			})
		})

		When("list is not empty and has multiple items", func() {
			It("should remove the first value", func() {
				l.PushBack("hello")
				l.PushBack("world")
				value := l.PopFront()
				Expect(value).To(Equal("hello"))
			})
		})

	})

	Context("PopBack", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})

		When("list is empty", func() {
			It("should return nil", func() {

			})
		})

		When("list is not empty", func() {
			It("should remove the last value", func() {

			})
		})

	})

	Context("ValueAt", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})

		When("index is less than size", func() {
			It("should return the value", func() {

			})
		})
		When("index is greater than size", func() {
			It("should return nil", func() {

			})
		})
	})

	Context("Insert", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})

		When("index is zero", func() {
			It("should be same as pushFront", func() {

			})
		})

		When("index is greater than or equal to size", func() {
			It("should be same as pushback", func() {

			})
		})

		When("index is than the size", func() {
			It("element should be added", func() {

			})
		})
	})

	Context("Erase", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})

		When("index is zero", func() {
			It("should be same as popFront", func() {

			})
		})

		When("index is greater than or equal to size", func() {
			It("should be same as popBack", func() {

			})
		})

		When("index is than the size", func() {
			It("element should be removed", func() {

			})
		})
	})

	Context("ValueFromEnd", func() {
		When("index is greater than the size", func() {
			It("should return the mod of the size", func() {

			})
		})

		When("index is less than size", func() {
			It("should return the object", func() {

			})
		})
	})

	Context("Reverse", func() {
		It("should reverse the list", func() {

		})
	})

	Context("RemoveValue", func() {
		When("value is present", func() {
			It("remove the first occurances", func() {

			})
		})

		When("value is not present", func() {
			It("should not remove anything", func() {

			})
		})
	})

})
