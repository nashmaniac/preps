package avl_test

import (
	"github.com/nashmaniac/golang-coding/datastructures/avl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AVLTree", func() {
	Context("Insert", func() {

		Context("when the tree is right heavy", func() {
			var l avl.AVLTree

			When("1 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
				})
				It("passes", func() {
					value := 1
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(value))
					Expect(root.GetLeft()).To(BeNil())
					Expect(root.GetRight()).To(BeNil())
					Expect(root.GetParent()).To(BeNil())
				})
			})

			When("2 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
					l.Insert(1)
				})
				It("passes", func() {
					value := 2
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(1), "root value should be the same")
					Expect(root.GetHeight()).To(Equal(1), "root height should be 1")
					Expect(root.GetLeft()).To(BeNil(), "root left should be empty")
					Expect(root.GetRight()).ToNot(BeNil(), "root right should not be empty")
					Expect(root.GetRight().GetValue()).To(Equal(value), "value should be 2")
					Expect(root.GetParent()).To(BeNil(), "value for the root parent is empty")
				})
			})

			When("3 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
					l.Insert(1)
					l.Insert(2)
				})
				It("passes", func() {
					value := 3
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(2), "root value should be the same")
					Expect(root.GetHeight()).To(Equal(1), "root height should be 1")
					Expect(root.GetLeft()).ToNot(BeNil(), "root left should be empty")
					Expect(root.GetRight()).ToNot(BeNil(), "root right should not be empty")
					Expect(root.GetLeft().GetValue()).To(Equal(1), "value should be 1")
					Expect(root.GetRight().GetValue()).To(Equal(value), "value should be 3")
					Expect(root.GetParent()).To(BeNil(), "value for the root parent is empty")
				})
			})

		})

		Context("when the tree is right left heavy", func() {
			var l avl.AVLTree

			When("60 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
				})
				It("passes", func() {
					value := 60
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(value))
					Expect(root.GetLeft()).To(BeNil())
					Expect(root.GetRight()).To(BeNil())
					Expect(root.GetParent()).To(BeNil())
				})
			})

			When("70 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
					l.Insert(60)
				})
				It("passes", func() {
					value := 70
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(60), "root value should be the same")
					Expect(root.GetHeight()).To(Equal(1), "root height should be 1")
					Expect(root.GetLeft()).To(BeNil(), "root left should be empty")
					Expect(root.GetRight()).ToNot(BeNil(), "root right should not be empty")
					Expect(root.GetRight().GetValue()).To(Equal(value), "value should be 70")
					Expect(root.GetParent()).To(BeNil(), "value for the root parent is empty")
				})
			})

			When("65 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
					l.Insert(60)
					l.Insert(70)
				})
				It("passes", func() {
					value := 65
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(value), "root value should be the same")
					Expect(root.GetHeight()).To(Equal(1), "root height should be 1")
					Expect(root.GetLeft()).ToNot(BeNil(), "root left should be empty")
					Expect(root.GetRight()).ToNot(BeNil(), "root right should not be empty")
					Expect(root.GetLeft().GetValue()).To(Equal(60), "value should be 60")
					Expect(root.GetRight().GetValue()).To(Equal(70), "value should be 70")
					Expect(root.GetParent()).To(BeNil(), "value for the root parent is empty")
				})
			})

		})

		Context("when the tree is left heavy", func() {
			var l avl.AVLTree

			When("30 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
				})
				It("passes", func() {
					value := 30
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(value))
					Expect(root.GetLeft()).To(BeNil())
					Expect(root.GetRight()).To(BeNil())
					Expect(root.GetParent()).To(BeNil())
				})
			})

			When("20 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
					l.Insert(30)
				})
				It("passes", func() {
					value := 20
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(30), "root value should be the same")
					Expect(root.GetHeight()).To(Equal(1), "root height should be 1")
					Expect(root.GetLeft()).ToNot(BeNil(), "root left should be empty")
					Expect(root.GetRight()).To(BeNil(), "root right should not be empty")
					Expect(root.GetLeft().GetValue()).To(Equal(value), "value should be 20")
					Expect(root.GetParent()).To(BeNil(), "value for the root parent is empty")
				})
			})

			When("10 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
					l.Insert(30)
					l.Insert(20)
				})
				It("passes", func() {
					value := 10
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(20), "root value should be the same")
					Expect(root.GetHeight()).To(Equal(1), "root height should be 1")
					Expect(root.GetLeft()).ToNot(BeNil(), "root left should be empty")
					Expect(root.GetRight()).ToNot(BeNil(), "root right should not be empty")
					Expect(root.GetLeft().GetValue()).To(Equal(10), "value should be 10")
					Expect(root.GetRight().GetValue()).To(Equal(30), "value should be 30")
					Expect(root.GetParent()).To(BeNil(), "value for the root parent is empty")
				})
			})

		})

		Context("when the tree is left right heavy", func() {
			var l avl.AVLTree

			When("30 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
				})
				It("passes", func() {
					value := 30
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(value))
					Expect(root.GetLeft()).To(BeNil())
					Expect(root.GetRight()).To(BeNil())
					Expect(root.GetParent()).To(BeNil())
				})
			})

			When("10 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
					l.Insert(30)
				})
				It("passes", func() {
					value := 10
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(30), "root value should be the same")
					Expect(root.GetHeight()).To(Equal(1), "root height should be 1")
					Expect(root.GetLeft()).ToNot(BeNil(), "root left should be empty")
					Expect(root.GetRight()).To(BeNil(), "root right should not be empty")
					Expect(root.GetLeft().GetValue()).To(Equal(value), "value should be 20")
					Expect(root.GetParent()).To(BeNil(), "value for the root parent is empty")
				})
			})

			When("20 is added", func() {
				BeforeEach(func() {
					l = avl.NewAVLTree()
					l.Insert(30)
					l.Insert(10)
				})
				It("passes", func() {
					value := 20
					l.Insert(value)
					root := l.GetRoot()
					Expect(root.GetValue()).To(Equal(20), "root value should be the same")
					Expect(root.GetHeight()).To(Equal(1), "root height should be 1")
					Expect(root.GetLeft()).ToNot(BeNil(), "root left should be empty")
					Expect(root.GetRight()).ToNot(BeNil(), "root right should not be empty")
					Expect(root.GetLeft().GetValue()).To(Equal(10), "value should be 10")
					Expect(root.GetRight().GetValue()).To(Equal(30), "value should be 30")
					Expect(root.GetParent()).To(BeNil(), "value for the root parent is empty")
				})
			})

		})

	})
})
