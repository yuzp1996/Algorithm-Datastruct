package Linkedtree_test

import (
	. "Algorithm-Datastruct/golang/Tree"
	"fmt"
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
		result = []int{0, 1, 3, 4, 2, 5, 6}
		It("preorder traversal", func() {
			Expect(Root.Preorder([]int{})).To(Equal(result))
		})
	})
	Context("Traversing with middle", func() {
		middleresult := []int{3, 1, 4, 0, 5, 2, 6}
		It("middle traversal", func() {
			Expect(Root.Middleorder([]int{})).To(Equal(middleresult))
		})
	})
	Context("Traversing with behind", func() {
		behindresult := []int{3, 4, 1, 5, 6, 2, 0}
		It("middle traversal", func() {
			Expect(Root.BehindOrder([]int{})).To(Equal(behindresult))
		})
	})
	Context("Traversing with level", func() {
		levelresult := []int{0, 1, 2, 3, 4, 5, 6}
		It("level traversal", func() {
			Expect(Root.LevelOrder([]int{})).To(Equal(levelresult))
		})
	})
	Context("Traversing with level1", func() {
		levelresult := []int{0, 1, 2, 3, 4, 5, 6}
		It("leve1 traversal", func() {
			Expect(Root.Level()).To(Equal(levelresult))
		})
	})


	Context("binary search tree", func() {
		BianrysearchRoot := NewTree(3)
		BianrysearchRoot.Insert(2)
		BianrysearchRoot.Insert(1)
		BianrysearchRoot.Insert(4)
		It("return the right leaf", func() {
			Expect(BianrysearchRoot.Right.Value).To(Equal(4))
			Expect(BianrysearchRoot.Left.Left.Value).To(Equal(1))
		})
		It("search the right leaf", func() {
			Expect(BianrysearchRoot.Search(2)).To(Equal(true))
		})
		It("New Delete", func() {
			Tree := NewTree(29)
			Tree.Insert(19)
			Tree.Insert(20)
			Tree.Insert(18)
			Tree.Insert(1)
			Tree.Insert(45)
			Tree.Insert(40)
			Tree.Insert(50)
			fmt.Printf("Tree is %v \n",Tree.Level())
			Expect(Tree.DeleteLeaf(0)).To(Equal(-1))
			Expect(Tree.Level()).To(Equal([]int{29,19,45,18,20,40,50,1}))
			Expect(Tree.DeleteLeaf(1)).To(Equal(1))
			Expect(Tree.Level()).To(Equal([]int{29,19,45,18,20,40,50}))
			Expect(Tree.DeleteLeaf(45)).To(Equal(45))
			Expect(Tree.Level()).To(Equal([]int{29,19,50,18,20,40}))
			Expect(Tree.DeleteLeaf(50)).To(Equal(50))
			Expect(Tree.Level()).To(Equal([]int{29,19,40,18,20}))
			Expect(Tree.DeleteLeaf(19)).To(Equal(19))
			Expect(Tree.Level()).To(Equal([]int{29,20,40,18}))
			Expect(Tree.DeleteLeaf(29)).To(Equal(29))
			Expect(Tree.Level()).To(Equal([]int{40,20,18}))
			Expect(Tree.DeleteLeaf(20)).To(Equal(20))
			Expect(Tree.Level()).To(Equal([]int{40,18}))
			Expect(Tree.DeleteLeaf(40)).To(Equal(40))
			Expect(Tree.Level()).To(Equal([]int{18}))
		})
	})

})
