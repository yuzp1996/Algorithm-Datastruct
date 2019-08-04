package BF_test

import (
	"Algorithm-Datastruct/golang/String/BF"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BF", func() {

	Context("find if there is the pattern", func() {
		var (
			Main  string
			MatchPattern string
			UnMatchPattern string
		)
		BeforeEach(func() {
			Main = "yuzhipengshiyigedahaorenbalabaala"
			MatchPattern = "yuzhipeng"
			UnMatchPattern = "xuyahui"
		})
		It("find the main", func() {
			exist := BF.BFfindstring(Main, MatchPattern)
			Expect(exist).To(Equal(true))

		})
		It("find no main", func() {
			exist := BF.BFfindstring(Main, UnMatchPattern)
			Expect(exist).To(Equal(false))
		})
		It("Rk find the string", func() {
			exist := BF.RKfindstring(Main, UnMatchPattern)
			Expect(exist).To(Equal(false))
		})
	})
})
