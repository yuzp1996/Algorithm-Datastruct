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
			Main2          string
			MatchPattern2  string
		)
		BeforeEach(func() {
			Main = "aaaassssdsdsdsds1zhipengyuzhipeng"
			MatchPattern = "zhipengyuzhipeng"
			UnMatchPattern = "xuyahui"

			Main1 = "xxxwwwwabcdwwwqqqaaaaaxxxabcdabcabcdab"
			MatchPattern1 = "abcdabc"

			Main2 = "asfdlnjasdfoousdfhodsohafodsnfadsfsoodfoodsdfds"
			MatchPattern2 = "odfood"

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

		It("BM1 should find the right string", func() {
			exist := BM.BMfindstring(Main1, MatchPattern1)
			Expect(exist).To(Equal(true))
		})

		It("BM1 should find the right string", func() {
			exist := BM.BMfindstring(Main2, MatchPattern2)
			Expect(exist).To(Equal(true))
		})

		It("GoodSuffix can find if there is the matching string", func() {
			exsit := BM.GoodSuffix(Main1, MatchPattern1)
			Expect(exsit).To(Equal(true))

			exsit1 := BM.GoodSuffix(Main1, UnMatchPattern)
			Expect(exsit1).To(Equal(false))

			exsit2 := BM.GoodSuffix(Main, MatchPattern)
			Expect(exsit2).To(Equal(true))

			exsit3 := BM.GoodSuffix(Main2, MatchPattern2)
			Expect(exsit3).To(Equal(true))

		})

		It("GetNeededArray should return the right three array", func() {
			Pattern := "abcdefabc"
			patternwithlength, suffix, prefix := BM.Getneededarray(Pattern)

			Expect(patternwithlength).To(Equal([]string{"c", "bc", "abc", "fabc", "efabc", "defabc", "cdefabc", "bcdefabc"}))
			Expect(suffix).To(Equal([]int{2, 1, 0, -1, -1, -1, -1, -1}))
			Expect(prefix).To(Equal([]bool{false, false, true, false, false, false, false, false}))

		})

		It("find the right string before . ", func() {
			NameWithVersion := "GoLang1.10Build.1.10.9"
			ExpectString := "GoLang1.10Build"
			result := BM.GetNameWithOutVersion(ExpectString, NameWithVersion)
			Expect(result).To(Equal(true))
		})

		It("find version with RE", func() {
			NameWithVersion := "GoLangBuilder2.7.1.10.9"
			result := BM.GetNameWithRE(NameWithVersion)
			Expect(result).To(Equal("GoLangBuilder2.7"))

			NameWithVersion1 := "Python2.7Builder.1.10.9"
			result1 := BM.GetNameWithRE(NameWithVersion1)
			Expect(result1).To(Equal("Python2.7Builder"))

			NameWithOutVersion := "GoLangBuilder"
			result = BM.GetNameWithRE(NameWithOutVersion)
			Expect(result).To(Equal("nosuchthing"))

		})

	})
})
