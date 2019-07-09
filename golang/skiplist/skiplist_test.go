package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "Algorithm-Datastruct/golang/skiplist"
)

var _ = Describe("Skiplist", func() {
	Context("Skiplist can find the number ", func() {
		It("If I delete some number It will not find this number", func() {
			skiplist := NewSkipList()
			for i:=1;i<20;i++{
				skiplist.Insert(i)
			}
			skiplist.Delete(12)
			Expect(skiplist.Search(12)).To(Equal(false))
			Expect(skiplist.Search(13)).To(Equal(true))
			skiplist.Insert(12)

			Expect(skiplist.Search(12)).To(Equal(true))
		})
	})

})































