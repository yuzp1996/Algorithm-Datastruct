package arrayqueue_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestArrayqueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Arrayqueue Suite")
}
