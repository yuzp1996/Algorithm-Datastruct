package BinarySearch_test

import (
	"Algorithm-Datastruct/golang/BinarySearch"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinarySearch", func() {
	Context("should return the right index", func() {
		It("should return a sorted array", func() {
			Expect(BinarySearch.Binarysearch([]int{1,3,4,5,6,7,9,10,13,15,19},8),).To(Equal(false))
		})
	})
})
