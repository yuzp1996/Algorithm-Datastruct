package BF_test

import (
	"Algorithm-Datastruct/golang/String/BF"
	"Algorithm-Datastruct/golang/String/BM"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BF", func() {

	Context("find if there is the pattern", func() {
		var (
			Main           string
			MatchPattern   string
			UnMatchPattern string
		)
		BeforeEach(func() {
			Main = "gogogogogogogoddddgggggyuzhipengshiyigedahaorenbalabaala"
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
		It("Rk can notfind the string", func() {
			exist := BF.RKfindstring(UnMatchPattern, Main)
			Expect(exist).To(Equal(false))
		})
		It("Rk can find the string", func() {
			exist := BF.RKfindstring(Main, MatchPattern)
			Expect(exist).To(Equal(true))
		})

		It("BM should find the right string", func() {
			exist := BM.BMfindstring(Main, MatchPattern)
			Expect(exist).To(Equal(true))
		})

	})
})
