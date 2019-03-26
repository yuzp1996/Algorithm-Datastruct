package BinarySearch_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBinarySearch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BinarySearch Suite")
}
