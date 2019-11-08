package KMP_test

import (
	. "Algorithm-Datastruct/golang/String/KMP"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KMP", func() {
	It("kmp algorithm", func() {
		result := MainSkeleTon("abcabcabcd", "abcd")
		Expect(result).To(Equal(true))

		next := Failurefunction("zpyduzpsyuzpdyuzxpyu")
		Expect(next).To(Equal([]int{1,3,4}))

	})
})
