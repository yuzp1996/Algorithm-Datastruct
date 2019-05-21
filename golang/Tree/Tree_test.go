package Linkedtree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "Algorithm-Datastruct/golang/Tree"
)

var _ = Describe("Tree", func() {
	Root := NewTree(0)
	Root.LeafAdd(1)
	Expect(Root.Left.Value).To(Equal())
})
