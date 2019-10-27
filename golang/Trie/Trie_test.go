package Trie_test

import (
	. "github.com/onsi/ginkgo"
	"Algorithm-Datastruct/golang/Trie"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trie", func() {
	root := Trie.NewTrieTree()
	root.Insert("abcd")
	result := root.Find("abcd")
	It("Insert test the int()", func() {
		Expect(result).To(Equal(true))
	})
})
