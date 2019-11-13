package AC_test

import (
	. "github.com/onsi/ginkgo"
	. "Algorithm-Datastruct/golang/Trie/AC"
	. "github.com/onsi/gomega"
)

var _ = Describe("AC", func() {
	It("BuildFailurepointer", func() {
		ACTree := NewACNodeTree()
		ACTree.Insert("abcd")
		ACTree.Insert("bc")
		ACTree.Insert("abdc")
		ACTree.Insert("yuzhipeng")


		ACTree.BuildFailurePointer()
		Expect(ACTree.Leafs[0].FailurePointer.Value).To(Equal(rune('&')))
		Expect(ACTree.Leafs[0].Leafs[1].FailurePointer.Value).To(Equal(rune('b')))
		Expect(ACTree.Leafs[0].Leafs[1].FailurePointer.FailurePointer.Value).To(Equal(rune('&')))
		Expect(ACTree.Leafs[0].Leafs[1].Leafs[2].FailurePointer.FailurePointer.Value).To(Equal(rune('&')))
		Expect(ACTree.Leafs[0].Endchar).To(Equal(false))
		Expect(ACTree.Leafs[0].Leafs[1].Leafs[2].Leafs[3].Endchar).To(Equal(true))

		result := ACTree.FindString("dsbcfdabcdsfdsabddcyuzhipengdbccdefabddcd")
		Expect(result).To(Equal([]string{"bc", "abcd", "yuzhipeng", "bc"}))
	})
})
