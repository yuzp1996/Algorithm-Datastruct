package Trie_test

import (
	. "Algorithm-Datastruct/golang/Trie"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trie", func() {
	root := NewTrieTree()
	root.Insert("abdcd")
	root.Insert("bcefds")
	root.Insert("abdsfsdcd")
	result := root.Find("abdcd")
	It("Insert test the int()", func() {
		Expect(root).To(Equal(0))
		Expect(result).To(Equal(true))
	})

})
