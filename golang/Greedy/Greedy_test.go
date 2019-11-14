package Greedy_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "Algorithm-Datastruct/golang/Greedy"
)

var _ = Describe("Greedy", func() {
	It("Test Greedy", func() {
		Expect(FindTheLatest(1987321)).To(Equal(187321))

	})
	It("test ConstructNum", func() {
		Expect(ConstructNum([]int{1,0,2,3,4})).To(Equal(10234))
	})

})
