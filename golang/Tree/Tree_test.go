package Linkedtree_test

import (
	. "Algorithm-Datastruct/golang/Tree"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tree", func() {
	Root := NewTree(0)
	Root.LeafAdd(1)
	Context("test the tree", func() {

		It("Insert should be successful", func() {
			Expect(Root.Left.Value).To(Equal(1))
		})
	})
})
