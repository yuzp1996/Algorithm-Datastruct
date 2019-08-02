package ListGraph_test

import (
	"github.com/onsi/ginkgo/reporters"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestListGraph(t *testing.T) {
	RegisterFailHandler(Fail)
	report := reporters.NewJUnitReporter("listgraph.xml")
	RunSpecsWithCustomReporters(t,"listgraph", []Reporter{report})
}
