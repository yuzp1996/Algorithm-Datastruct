package divide_test

import (
	"Algorithm-Datastruct/golang/divide"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = Describe("Divide", func() {
	It("test merge", func() {
		gomega.Expect(divide.MergeSort([]int{3,2,4,5,6,1,7,8,4})).To(gomega.Equal([]int{1,2,3,4,4,5,6,7,8}))


		gomega.Expect(divide.MergeSort([]int{3,2,4,234,5465,23,54354,5,5435,6,45646,1,654654,7,8,4})).To(gomega.Equal([]int{1, 2, 3, 4, 4, 5, 6, 7, 8, 23, 234, 5435, 5465, 45646, 54354, 654654}))
		gomega.Expect(divide.MergeSort([]int{3,2})).To(gomega.Equal([]int{2,3}))

	})
})
