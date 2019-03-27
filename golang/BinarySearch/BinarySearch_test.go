package BinarySearch_test

import (
	"Algorithm-Datastruct/golang/BinarySearch"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinarySearch", func() {

	Context("should return the right index", func() {
		It("should return a sorted array", func() {
			Expect(BinarySearch.Binarysearch([]int{1,3,4,5,6,7,9,10,13,15,19},1)).To(Equal(0))
			Expect(BinarySearch.Binarysearch([]int{1,3,4,5,6,7,9,10,13,15,19},2)).To(Equal(-1))
			Expect(BinarySearch.Binarysearch([]int{1,3,4,5,6,7,9,10,13,15,19},20)).To(Equal(-1))
			Expect(BinarySearch.Binarysearch([]int{1,3,4,5,6,7,9,10,13,15,19},11)).To(Equal(-1))
			Expect(BinarySearch.Binarysearch([]int{1,3,4,5,6,7,9,10,13,15,19},19)).To(Equal(10))
			Expect(BinarySearch.Binarysearch([]int{1,3,4,5,6,7,9,10,13,15,19},9)).To(Equal(6))
			Expect(BinarySearch.Binarysearch([]int{1,3,4,5,6,7,9,10,13,15,19},9)).To(Equal(6))

		})
	})
	Context("Resucrsion should return the right index", func() {
		It("should return a sorted array", func() {
			Expect(BinarySearch.RecursionBinarySearch([]int{1,3,4,5,6,7,9,10,13,15,19},1)).To(Equal(0))
			Expect(BinarySearch.RecursionBinarySearch([]int{1,3,4,5,6,7,9,10,13,15,19},2)).To(Equal(-1))
			Expect(BinarySearch.RecursionBinarySearch([]int{1,3,4,5,6,7,9,10,13,15,19},20)).To(Equal(-1))
			Expect(BinarySearch.RecursionBinarySearch([]int{1,3,4,5,6,7,9,10,13,15,19},11)).To(Equal(-1))
			Expect(BinarySearch.RecursionBinarySearch([]int{1,3,4,5,6,7,9,10,13,15,19},19)).To(Equal(10))
			Expect(BinarySearch.RecursionBinarySearch([]int{1,3,4,5,6,7,9,10,13,15,19},9)).To(Equal(6))
			Expect(BinarySearch.RecursionBinarySearch([]int{1,3,4,5,6,7,9,10,13,15,19},9)).To(Equal(6))

		})
	})
	Context("Resucrsion should return the First Big right index", func() {
		It("should return a sorted array", func() {
			Expect(BinarySearch.Getfistindex([]int{1,1,3,4,5,6,7,9,10,13,15,19},1)).To(Equal(0))
			Expect(BinarySearch.Getfistindex([]int{1,3,4,5,6,7,9,10,13,15,19},2)).To(Equal(-1))
			Expect(BinarySearch.Getfistindex([]int{1,3,4,5,6,7,9,10,13,15,19},20)).To(Equal(-1))
			Expect(BinarySearch.Getfistindex([]int{1,3,4,5,6,7,9,10,13,15,19},11)).To(Equal(-1))
			Expect(BinarySearch.Getfistindex([]int{1,3,4,5,6,7,9,10,13,15,19,19},19)).To(Equal(10))
			Expect(BinarySearch.Getfistindex([]int{1,3,4,5,6,7,9,9,9,10,13,15,19},9)).To(Equal(6))
			Expect(BinarySearch.Getfistindex([]int{1,3,4,5,6,7,9,10,13,15,15,15,19},15)).To(Equal(9))

		})
	})
	Context("Resucrsion should return the Last Big right index", func() {
		It("should return a sorted array", func() {
			Expect(BinarySearch.Getlastindex([]int{1,1,3,4,5,6,7,9,10,13,15,19},1)).To(Equal(1))
			Expect(BinarySearch.Getlastindex([]int{1,3,4,5,6,7,9,10,13,15,19},2)).To(Equal(-1))
			Expect(BinarySearch.Getlastindex([]int{1,3,4,5,6,7,9,10,13,15,19},20)).To(Equal(-1))
			Expect(BinarySearch.Getlastindex([]int{1,3,4,5,6,7,9,10,13,15,19},11)).To(Equal(-1))
			Expect(BinarySearch.Getlastindex([]int{1,3,4,5,6,7,9,10,13,15,19,19},19)).To(Equal(11))
			Expect(BinarySearch.Getlastindex([]int{1,3,4,5,6,7,9,9,9,10,13,15,19},9)).To(Equal(8))
			Expect(BinarySearch.Getlastindex([]int{1,3,4,5,6,7,9,10,13,15,15,15,19},15)).To(Equal(11))
		})
	})
	Context("Resucrsion should return index of the first Bigger than the num", func() {
		It("should return a sorted array", func() {
			Expect(BinarySearch.GetFirstBigger([]int{1,1,3,4,5,6,7,9,10,13,15,19},1)).To(Equal(0))
			Expect(BinarySearch.GetFirstBigger([]int{1,3,4,5,6,7,9,10,13,15,19},2)).To(Equal(1))
			Expect(BinarySearch.GetFirstBigger([]int{1,3,4,5,6,7,9,10,13,15,19},20)).To(Equal(-1))
			Expect(BinarySearch.GetFirstBigger([]int{1,3,4,5,6,7,9,10,13,15,19},11)).To(Equal(8))
			Expect(BinarySearch.GetFirstBigger([]int{1,3,4,5,6,7,9,10,13,15,19,19},19)).To(Equal(10))
			Expect(BinarySearch.GetFirstBigger([]int{1,3,4,5,6,7,9,9,9,10,13,15,19},9)).To(Equal(6))
			Expect(BinarySearch.GetFirstBigger([]int{1,3,4,5,6,7,9,10,13,15,15,15,19},15)).To(Equal(9))
		})
	})
	Context("Resucrsion should return index of the last samller than the num", func() {
		It("should return a sorted array", func() {
			Expect(BinarySearch.Lastsamller([]int{1,1,3,4,5,6,7,9,10,13,15,19},1)).To(Equal(1))
			Expect(BinarySearch.Lastsamller([]int{1,3,4,5,6,7,9,10,13,15,19},2)).To(Equal(0))
			Expect(BinarySearch.Lastsamller([]int{1,3,4,5,6,7,9,10,13,15,19},20)).To(Equal(10))
			Expect(BinarySearch.Lastsamller([]int{1,3,4,5,6,7,9,10,13,15,19},11)).To(Equal(7))
			Expect(BinarySearch.Lastsamller([]int{1,3,4,5,6,7,9,10,13,15,19,19},19)).To(Equal(11))
			Expect(BinarySearch.Lastsamller([]int{1,3,4,5,6,7,9,9,9,10,13,15,19},9)).To(Equal(8))
			Expect(BinarySearch.Lastsamller([]int{1,3,4,5,6,7,9,10,13,15,15,15,19},15)).To(Equal(11))
		})
	})
})
