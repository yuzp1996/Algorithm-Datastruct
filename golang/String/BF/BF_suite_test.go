package BF_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBF(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BF Suite")
}
