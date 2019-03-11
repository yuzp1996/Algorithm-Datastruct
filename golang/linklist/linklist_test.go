package main

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Testlinklist(t *testing.T){
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Linklist")
}

func Createlinklist()*LinkList{
	testlinklist := NewLinkList(1)
	testlinklist.InsterInTail(2)
	testlinklist.InsterInTail(3)
	testlinklist.InsterInTail(4)
	testlinklist.InsterInTail(5)
	return testlinklist
}

var _ = Describe("Shopping cart", func() {
	testlinklist := Createlinklist()

	Context("Find the middle element of a linklist", func() {
		It("Find the middle element",func(){
		 	Expect(testlinklist.FindMiddlenode().Getvalue()).Should(Equal(3))
		})
	})
	Context("reserver the link list", func() {
		It("should reserver the linklist", func() {
			testlinklist.reserve()
			Expect(testlinklist.findfianlNode().Getvalue()).Should(Equal(0))
			Expect(testlinklist.FindwithIndex(1).Getvalue()).Should(Equal(5))
		})
		It("recursion reserver ", func() {
			//recursionresrver()
			//Expect(testlinklist.findfianlNode().Getvalue()).Should(Equal(0))
			//Expect(testlinklist.FindwithIndex(1).Getvalue()).Should(Equal(5))
		})
		AfterEach(func() {
			testlinklist = Createlinklist()
		})

	})

})
