package ListGraph_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestListGraph(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ListGraph Suite")
}
