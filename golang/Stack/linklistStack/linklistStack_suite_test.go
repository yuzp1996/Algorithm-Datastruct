package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLinklistStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LinklistStack Suite")
}
