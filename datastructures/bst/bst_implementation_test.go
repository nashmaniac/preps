package bst_test

import (
	"github.com/nashmaniac/golang-coding/datastructures/bst"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BST", func() {

	var tree bst.BST

	BeforeEach(func() {
		tree = bst.NewBST()
	})

	Context("Insert", func() {
		When("constructed correctly", func() {
			It("passes", func() {
				inputValues := []int{
					3, 6, 5, 2, 10, 8, 1, 7, 4, 9,
				}

				for _, i := range inputValues {
					tree.Insert(i)
				}

				values := tree.InOrderWalk()
				for index, i := range values {
					Expect(index + 1).To(Equal(i))
				}
			})
		})
	})
})
