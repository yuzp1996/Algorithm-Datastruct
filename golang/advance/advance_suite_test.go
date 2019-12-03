package advance_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAdvance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Advance Suite")
}
