package BF_test

import (
	"Algorithm-Datastruct/golang/String/BF"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BF", func() {
	Context("find if there is the pattern", func() {
		It("find the main", func() {
			exist := BF.BFfindstring("yuzhipengshiyigedahaorenbalabaala", "dahaoren")
			Expect(exist).To(Equal(true))

		})
		It("find no main", func() {
			exist := BF.BFfindstring("yuzhipengs   hiyigedahaorenbalabaala", "xuyahui")
			Expect(exist).To(Equal(false))
		})
		It("Rk find the string", func() {
			exist := BF.RKfindstring("yuzhixuyaihuiiiiii", "xuyahui")
			Expect(exist).To(Equal(false))
		})
	})
})
