package KMP_test

import (
	. "Algorithm-Datastruct/golang/String/KMP"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KMP", func() {
	It("kmp algorithm", func() {
		result := MainSkeleTon("dsfdsfdsfdsfdsfsdabababzabababsfdsfdsabababzabababdsfdsfdsabababzababab", "abababzababab")
		Expect(result).To(Equal(true))

	})
	It("test failuerfuntion", func() {
		pattern := "abababzababab"
		next := FailuerFunction(pattern)
		Expect(next).To(Equal([]int{-1,-1,0,1,2,3,-1,0,1,2,3,4}))
	})
})
