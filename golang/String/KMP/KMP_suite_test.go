package KMP_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKMP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KMP Suite")
}
