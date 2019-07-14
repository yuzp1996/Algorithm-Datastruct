package Recursion_test

import (
	"Algorithm-Datastruct/golang/Recursion"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Recursion", func() {
	Context("how many staris should I use", func() {
		It("how many stairs should use", func() {
			Expect(Recursion.Ways(29)).To(Equal(832040))
			Expect(Recursion.WaysIteration(29)).To(Equal(832040))

		})
	})
	Context("Fibonacci ", func() {
		It("Fibonacci 0 should return 0", func() {
			Expect(Recursion.Fibonacci(0)).To(Equal(0))
			Expect(Recursion.Fibonacci(8)).To(Equal(21))
			Expect(Recursion.Factorial(4)).To(Equal(24))
			Expect(Recursion.Factorial(5)).To(Equal(120))

		})
	})
	Context("Arrangement ", func() {
		It("Arrangement 0 should return 0", func() {
			Expect(Recursion.Arrangement(1)).To(Equal(1))
			Expect(Recursion.Arrangement(2)).To(Equal(2))
			Expect(Recursion.Arrangement(3)).To(Equal(6))

		})
	})

})
