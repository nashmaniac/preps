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
				l.PushBack("hello")
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
				value := l.PopBack()
				Expect(value).To(BeNil())
			})
		})

		When("list is not empty", func() {
			It("should remove the first value if there is only element left", func() {
				l.PushFront("Hello")
				value := l.PopBack()
				Expect(value).To(Equal("Hello"))
				Expect(l.Size()).To(Equal(0))
				Expect(l.PopBack()).To(BeNil())
			})
			It("should remove the last value", func() {
				l.PushFront("Hello")
				l.PushBack("World")
				l.PushBack("Shetu")
				value := l.PopBack()
				Expect(value).To(Equal("Shetu"))
				Expect(l.Size()).To(Equal(2))
				Expect(l.PopBack()).To(Equal("World"))
				Expect(l.PopBack()).To(Equal("Hello"))
				Expect(l.Size()).To(Equal(0))
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
				values := []string{
					"Hello",
					"World",
					"Shetu",
					"Raju",
				}

				for _, i := range values {
					l.PushBack(i)
				}
				for index, value := range values {
					Expect(l.ValueAt(index)).To(Equal(value))
				}

			})
		})
		When("index is greater than size", func() {
			It("should return nil", func() {
				values := []string{
					"Hello",
					"World",
					"Shetu",
					"Raju",
				}

				for _, i := range values {
					l.PushBack(i)
				}

				Expect(l.ValueAt(10)).To(BeNil())
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
				l.PushBack("Hello")
				l.Insert(0, "World")
				Expect(l.Front()).To(Equal("World"))
			})
		})

		When("index is greater than or equal to size", func() {
			It("should be same as pushback", func() {
				l.PushBack("Hello")
				l.PushBack("World")
				l.Insert(2, "Shetu")
				Expect(l.Size()).To(Equal(3))
				Expect(l.Back()).To(Equal("Shetu"))
				l.Insert(6, "Hi")
				Expect(l.Size()).To(Equal(4))
				Expect(l.Back()).To(Equal("Hi"))
			})
		})

		When("index is SMALLER than the size", func() {
			It("element should be added", func() {
				l.PushBack("hello")
				l.PushBack("World")
				l.Insert(1, "Shetu")
				Expect(l.Size()).To(Equal(3))
				Expect(l.ValueAt(1)).To(Equal("Shetu"))
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
				values := []string{
					"hello", "world", "shetu",
				}
				for _, i := range values {
					l.PushBack(i)
				}
				l.Erase(0)
				Expect(l.Size()).To(Equal(2))
				Expect(l.Front()).To(Equal(values[1]))
			})
		})

		When("index is greater than or equal to size", func() {
			It("should be same as popBack", func() {
				values := []string{
					"hello", "world", "shetu",
				}
				for _, i := range values {
					l.PushBack(i)
				}

				l.Erase(9)
				Expect(l.Size()).To(Equal(2))
				Expect(l.Back()).To(Equal(values[1]))
			})
		})

		When("index is than the size", func() {
			It("element should be removed", func() {
				values := []string{
					"hello", "world", "one", "two", "3", "shetu",
				}
				for _, i := range values {
					l.PushBack(i)
				}

				l.Erase(3)
				Expect(l.Size()).To(Equal(5))
				Expect(l.ValueAt(3)).To(Equal(values[4]))
			})
		})
	})

	Context("ValueFromEnd", func() {
		var l linkedlist.LinkedList
		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})
		When("index is greater than the size", func() {
			It("should return the mod of the size", func() {
				values := []string{
					"hello", "world", "raju", "hey",
				}
				for _, i := range values {
					l.PushBack(i)
				}

				value := l.ValueFromEnd(6)
				Expect(value).To(Equal(l.ValueAt(2)))
			})
		})

		When("index is less than size", func() {
			It("should return the object", func() {
				values := []string{
					"hello", "world", "raju", "hey",
				}
				for _, i := range values {
					l.PushBack(i)
				}
				value := l.ValueFromEnd(3)
				Expect(value).To(Equal(l.ValueAt(1)))
			})
		})
	})

	Context("Reverse", func() {
		var l linkedlist.LinkedList
		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})
		It("should reverse the list", func() {
			values := []string{
				"hello", "world", "shetu", "high",
			}
			for _, i := range values {
				l.PushBack(i)
			}

			l.Reverse()

			n := l.Size()
			for index, i := range values {
				value := l.ValueAt(n - index - 1)
				Expect(value).To(Equal(i))
			}
		})
	})

	Context("RemoveValue", func() {
		var l linkedlist.LinkedList
		BeforeEach(func() {
			l = linkedlist.NewSinglyLinkedList[string]()
		})
		When("value is present", func() {
			It("remove the first occurances", func() {
				values := []string{
					"hello", "hello", "hello", "shetu",
				}

				for _, i := range values {
					l.PushBack(i)
				}

				Expect(l.Size()).To(Equal(4))
				l.RemoveValue("hello")
				Expect(l.Size()).To(Equal(3))
				Expect(l.Front()).To(Equal("hello"))
				l.RemoveValue("shetu")
				Expect(l.Size()).To(Equal(2))
				Expect(l.Back()).To(Equal("hello"))
			})
		})

		When("value is not present", func() {
			It("should not remove anything", func() {
				values := []string{
					"hello", "hello", "hello", "shetu",
				}

				for _, i := range values {
					l.PushBack(i)
				}

				l.RemoveValue("welcome")
				Expect(l.Size()).To(Equal(4))
				Expect(l.Front()).To(Equal("hello"))
				Expect(l.Back()).To(Equal("shetu"))
			})
		})
	})

})
