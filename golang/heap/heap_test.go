package heap_test

import (
	"Algorithm-Datastruct/golang/heap"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Heap", func() {
	var Heap *heap.Heap
	BeforeEach(func() {
		Heap = heap.NewHeap(5)
	})
	Context("BigTopheap BigTopInsert should be right ", func() {
		It("insert the right position", func() {
			Heap.BigTopInsert(10)
			Heap.BigTopInsert(5)
			Heap.BigTopInsert(6)
			Heap.BigTopInsert(7)
			Heap.BigTopInsert(12)
			Heap.BigTopInsert(13)

			Expect(Heap.Data[1]).To(Equal(12))
			Expect(Heap.Data[2]).To(Equal(10))
			Expect(Heap.Data[3]).To(Equal(6))
			Expect(Heap.Data[4]).To(Equal(5))
			Expect(Heap.Data[5]).To(Equal(7))
			Expect(Heap.Data[1]).To(Equal(12))
		})
	})
	Context("SmallTopHeap should insert the right value", func() {
		It("Insert the right value ", func() {
			Heap.SmallTopInsert(20)
			Heap.SmallTopInsert(15)
			Heap.SmallTopInsert(18)
			Heap.SmallTopInsert(14)

			Expect(Heap.Data[1]).To(Equal(14))
			Expect(Heap.Data[2]).To(Equal(15))
			Expect(Heap.Data[3]).To(Equal(18))
			Expect(Heap.Data[4]).To(Equal(20))
		})
	})

	Context("RemoveBigTop should remove the top value", func() {
		It("remove the top value ", func() {
			Heap.BigTopInsert(20)
			Heap.BigTopInsert(15)
			Heap.BigTopInsert(18)
			Heap.BigTopInsert(14)
			Heap.RemoveBigTop()
			Heap.RemoveBigTop()

			Expect(Heap.Data[1]).To(Equal(15))
			Expect(Heap.Data[2]).To(Equal(14))
			//Expect(Heap.Data[3]).To(Equal(0))
			//Expect(Heap.Data[4]).To(Equal(20))
		})
	})

	Context("Sort should work ", func() {
		It("sort should be good", func() {

			Heap.Data[1] = 7
			Heap.Data[2] = 3
			Heap.Data[3] = 9
			Heap.Data[4] = 14
			Heap.Data[5] = 16
			Heap.Count = 5

			Heap.Sort()
			for i := 1; i < 5; i++ {
				Expect(Heap.Data[i] < Heap.Data[i+1]).To(Equal(true))
			}
		})
	})

})
