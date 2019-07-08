package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSkiplist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Skiplist Suite")
}
