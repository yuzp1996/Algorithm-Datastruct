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
			PointTester{Name:1},
			PointTester{Name:2},
			PointTester{Name:3},
		}

		result = TestArray(Array)
		Expect(result).To(Equal(true))


		var PointerArray = &[]PointTester{
			PointTester{Name:1},
			PointTester{Name:2},
			PointTester{Name:3},
		}

		result = TestPointerArray(PointerArray)
		Expect(result).To(Equal(true))

		log.Printf("\n\n\n\n I am good one")
		var ArrayPointer = []*PointTester{
			&PointTester{Name:1},
			&PointTester{Name:2},
			&PointTester{Name:3},
		}
		arrayointerresult := TestArrayPointer(ArrayPointer)
		Expect(arrayointerresult).To(Equal(true))







	})
})
