package Linkedtree_test

import (
	. "Algorithm-Datastruct/golang/Tree"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tree", func() {
	var result []int
	Root := NewTree(0)
	Root.LeafAdd(1)
	Root.RightAdd(2)
	Root.Left.LeafAdd(3)
	Root.Left.RightAdd(4)
	Root.Right.LeafAdd(5)
	Root.Right.RightAdd(6)
	Context("test the tree", func() {
		It("Insert should be successful", func() {
			Expect(Root.Left.Value).To(Equal(1))
		})
		It("Insert right should be successful", func() {
			Expect(Root.Right.Value).To(Equal(2))
		})
	})
	Context("Traversing with fornt", func() {
		result = []int{0,1,3,4,2,5,6}
		It("preorder traversal", func() {
			Expect(Root.Preorder([]int{})).To(Equal(result))
		})
	})
	Context("Traversing with middle", func() {
		middleresult := []int{3,1,4,0,5,2,6}
		It("middle traversal", func() {
			Expect(Root.Middleorder([]int{})).To(Equal(middleresult))
		})
	})
})
