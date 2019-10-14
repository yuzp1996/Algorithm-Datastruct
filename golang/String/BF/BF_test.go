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
			Main1          string
			MatchPattern1  string
			UnMatchPattern string
		)
		BeforeEach(func() {
			Main = "aaaassssdsdsdsds1zhipengyuzhipeng"
			MatchPattern = "zhipengyuzhipeng"
			UnMatchPattern = "xuyahui"

			Main1 = "xxxwwwwabcdwwwqqqaaaaaxxxabcdabcabcdab"
			MatchPattern1 = "abcdabc"
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

		It("GoodSuffix can find if there is the matching string", func() {
			exsit := BM.GoodSuffix(Main1, MatchPattern1)
			Expect(exsit).To(Equal(true))
		})

		It("GetNeededArray should return the right three array", func() {
			Pattern := "abcdefabc"
			patternwithlength, suffix, _ := BM.Getneededarray(Pattern)

			Expect(patternwithlength).To(Equal([]string{"c", "bc", "abc", "fabc", "efabc", "defabc", "cdefabc", "bcdefabc"}))
			Expect(suffix).To(Equal([]int{2, 1, 0}))
			//Expect(prefix).To(Equal([]string{"c","bc","abc","fabc","efabc","defbac","cdefabc","bcdefabc"}))

		})


		It("find the right string before . ", func() {
			NameWithVersion := "GoLangBuild.1.10.9"
			ExpectString := "GoLangBuild"
			result := BM.GetNameWithOutVersion(NameWithVersion)
			Expect(result).To(Equal(ExpectString))
		})

		It("find version with RE", func() {
			NameWithVersion := "GoLangBuilder2.7.1.10.9"
			result := BM.GetNameWithRE(NameWithVersion)
			Expect(result).To(Equal("GoLangBuilder2.7"))
		})

	})
})
