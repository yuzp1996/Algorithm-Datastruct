package Sort_test

import (
	"Algorithm-Datastruct/golang/Sort"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = Describe("Sort", func() {
	Context("quick sort should return a sorted array", func() {
		It("should return a sorted array", func() {
			gomega.Expect(Sort.Quick_Sort([]int{10, 9, 8, 7, 6, 5, 4, 3, 20, 2, 1, 2, 9})).To(gomega.Equal([]int{1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10, 20}))
		})
		It("bubble sort return a sorted array", func() {
			gomega.Expect(Sort.Select_Sort([]int{10, 8, 7, 6, 5, 4, 3, 20, 1, 2, 9})).To(gomega.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20}))
		})
		It("Merge sort return a sorted array", func() {
			gomega.Expect(Sort.Merge_Sort([]int{10, 8, 7, 6, 6, 5, 4, 100, 3, 20, 1, 2, 9})).To(gomega.Equal([]int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10, 20, 100}))
		})
		It("Merge sort return a sorted array", func() {
			gomega.Expect(Sort.Insert_Sort([]int{10, 8, 7, 9, 6, 6, 5, 4, 100, 3, 20, 1, 1, 2, 0, 300, 202})).To(gomega.Equal([]int{0, 1, 1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10, 20, 100, 202, 300}))
		})
		It("Bubble sort return a sorted array", func() {
			gomega.Expect(Sort.Buuble_Sort([]int{10, 8, 7, 9, 6, 6, 5, 4, 100, 3, 20, 1, 1, 2, 0, 300, 202})).To(gomega.Equal([]int{0, 1, 1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10, 20, 100, 202, 300}))
		})
		It("Find the num ", func() {
			gomega.Expect(Sort.FindTheNum([]int{3, 3, 1, 2, 4, 5, 5, 4, 3, 2, 2, 100, 101}, 1)).To(gomega.Equal(1))
		})

	})

})
