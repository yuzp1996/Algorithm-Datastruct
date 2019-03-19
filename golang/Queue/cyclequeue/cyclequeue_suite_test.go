package cyclequeue_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCyclequeue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cyclequeue Suite")
}
