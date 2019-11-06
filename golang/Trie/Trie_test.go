package Trie_test

import (
	. "Algorithm-Datastruct/golang/Trie"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Trie", func() {
	root := NewTrieTree()
	root.Insert("abcd")
	result := root.Find("abcd")
	It("Insert test the int()", func() {

		var Array = []PointTester{
			PointTester{Name: 1},
			PointTester{Name: 2},
			PointTester{Name: 3},
		}

		result = TestArray(Array)
		Expect(result).To(Equal(true))

		//var PointerArray = &[]PointTester{
		//	PointTester{Name:1},
		//	PointTester{Name:2},
		//	PointTester{Name:3},
		//}
		//
		//result = TestPointerArray(PointerArray)
		//Expect(result).To(Equal(true))

		log.Printf("\n\n\n\n I am good one")
		var ArrayPointer = []*PointTester{
			&PointTester{Name: 1},
			&PointTester{Name: 2},
			&PointTester{Name: 3},
		}
		arrayointerresult := TestArrayPointer(ArrayPointer)
		Expect(arrayointerresult).To(Equal(true))


	})
	It("Test the Golang traps", func() {
		Test12()
	})

	It("Test18", func() {
		Test18()
	})

	It("Test45", func() {
		Test45()
	})
	It("Test47", func() {
		Test47()
	})
	It("Test14", func() {
		Test14()
	})
	It("Test31", func() {
		Test31()
	})
	It("Test311", func() {
		Test311()
	})
	It("Test33", func() {
		Test33()
	})
	It("Test34", func() {
		Test34()
	})
	It("Test53", func() {
		Test53()
	})

})
