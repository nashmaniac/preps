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
			l = linkedlist.NewDoublyLinkedList[string]()
		})

		When("initialized", func() {
			It("size should be zero", func() {
				Expect(l.Size()).To(Equal(0))
			})
		})

		When("added an element", func() {
			It("should increment the size by 1", func() {

			})
		})

		When("remove an element", func() {
			It("should decrement the size by 1", func() {

			})
		})

	})

	Context("Empty", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewDoublyLinkedList[string]()
		})

		When("initialized", func() {
			It("return empty", func() {
				Expect(l.Empty()).To(BeTrue())
			})
		})

		When("added an element", func() {
			It("should not be empty", func() {

			})
		})

		When("remove an element from more than an element", func() {
			It("should not be empty", func() {

			})
		})

		When("remove an element with a single element", func() {
			It("should be empty", func() {

			})
		})

	})

	Context("PushFront", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewDoublyLinkedList[string]()
		})

		When("added an element with empty list", func() {
			It("the front should be equal to value and size increment by 1", func() {

			})
		})

		When("added an element with a non-empty list", func() {
			BeforeEach(func() {

			})

			It("the front should be equal to value and size increment by 1", func() {

			})
		})

	})

	Context("PushBack", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewDoublyLinkedList[string]()
		})

		When("added an element with empty list", func() {
			It("the back should be equal to value and size increment by 1", func() {

			})
		})

		When("added an element with a non-empty list", func() {
			BeforeEach(func() {

			})

			It("the back should be equal to value and size increment by 1", func() {

			})
		})

	})

	Context("Front", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewDoublyLinkedList[string]()
		})

		When("list is empty", func() {
			It("should return nil", func() {

			})
		})

		When("list is not empty", func() {
			It("should return the first value", func() {

			})
		})

	})

	Context("Back", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewDoublyLinkedList[string]()
		})

		When("list is empty", func() {
			It("should return nil", func() {

			})
		})

		When("list is not empty", func() {
			It("should return the last value", func() {

			})
		})

	})

	Context("PopFront", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewDoublyLinkedList[string]()
		})

		When("list is empty", func() {
			It("should return nil", func() {

			})
		})

		When("list is not empty", func() {
			It("should remove the first value", func() {

			})
		})

	})

	Context("PopBack", func() {
		var l linkedlist.LinkedList

		BeforeEach(func() {
			l = linkedlist.NewDoublyLinkedList[string]()
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
			l = linkedlist.NewDoublyLinkedList[string]()
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
			l = linkedlist.NewDoublyLinkedList[string]()
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
			l = linkedlist.NewDoublyLinkedList[string]()
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
