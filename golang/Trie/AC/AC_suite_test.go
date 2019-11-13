package AC_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AC Suite")
}
