package heap_test

import (
	"Algorithm-Datastruct/golang/heap"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Heap", func() {
	var Heap *heap.Heap
	var Array1 []int
	var Array2 []int
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
			Heap.RemoveBigTop()


			Expect(Heap.Data[1]).To(Equal(14))
			//Expect(Heap.Data[2]).To(Equal(14))
			//Expect(Heap.Data[3]).To(Equal(0))
			//Expect(Heap.Data[4]).To(Equal(20))
		})
	})

	Context("RemoveSmallTop should remove the top value", func() {
		It("remove the top value ", func() {
			Heap.SmallTopInsert(20)
			Heap.SmallTopInsert(15)
			Heap.SmallTopInsert(18)
			Heap.SmallTopInsert(14)
			fmt.Print(Heap.Data)
			Heap.RemoveSmallTop()
			Heap.RemoveSmallTop()
			Heap.RemoveSmallTop()
			fmt.Print(Heap.Data)

			Expect(Heap.Data[1]).To(Equal(20))
			//Expect(Heap.Data[3]).To(Equal(0))
			//Expect(Heap.Data[4]).To(Equal(20))
		})
	})


	Context("Sort should work", func() {
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

	Context("top(K) insert", func() {
		JustBeforeEach(func() {
			Heap = heap.NewHeap(5)
		})
		It("top(K) insert should hold the biggest K value ", func() {
			Heap.TopKInsert(10)
			Heap.TopKInsert(12)
			Heap.TopKInsert(4)
			Heap.TopKInsert(5)
			Heap.TopKInsert(7)
			Heap.TopKInsert(3)
			Expect(Heap.GetTopKData()).To(Equal([]int{4, 5, 10, 12, 7}))
			Heap.TopKInsert(6)
			Expect(Heap.GetTopKData()).To(Equal([]int{5, 6, 10, 12, 7}))

		})
	})
	Context("Merge sorted array", func() {
		JustBeforeEach(func() {
			Heap = heap.NewHeap(6)
			Array1 = []int{3,5,7,9,11,13,15}
			Array2 = []int{2,4,6,8,10,12,14}
		})
		It("sort should merge the arraies", func() {
			Arrays:=[][]int{Array2,Array1}
			mergestring := Heap.MergeArray(Arrays)
			Expect(mergestring).To(Equal("23456789101112131415"))

		})
	})

})
