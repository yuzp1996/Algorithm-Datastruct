package Sort_test

import (
	"Algorithm-Datastruct/golang/Sort"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = Describe("Sort", func() {
	Context("quick sort should return a sorted array", func() {
		It("should return a sorted array", func() {
			gomega.Expect(Sort.QuickSort([]int{10,9,8,7,6,5,4,3,20,2,1,2,9})).To(gomega.Equal([]int{1,2,2,3,4,5,6,7,8,9,9,10,20}))
		})
		It("bubble sort return a sorted array", func() {
			gomega.Expect(Sort.Bubble([]int{10,8,7,6,5,4,3,20,1,2,9})).To(gomega.Equal([]int{1,2,3,4,5,6,7,8,9,10,20}))
		})
	})

})
