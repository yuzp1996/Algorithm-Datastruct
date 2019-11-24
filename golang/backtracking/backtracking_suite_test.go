package backtracking_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBacktracking(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Backtracking Suite")
}
