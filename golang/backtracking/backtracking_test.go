package backtracking_test

import (
	"Algorithm-Datastruct/golang/backtracking"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Backtracking", func() {
	It("test eight queue", func() {
		backtracking.EightQueue(0)
	})
	It("one zero package", func() {
		//index int, currentweight int,weightarray []int,totalnum int,loadwight int
		backtracking.ZerOnePackage(0,0,[]int{3,1,1,1,1,1,1,3,4,5,6},11,10)
	})
	It("test one zero package", func() {
		//index int, currentweight int,weightarray []int,totalnum int,loadwight int
		backtracking.SimpleZerOnePackage(0,0)
	})
})
