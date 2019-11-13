package misktake_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMisktake(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Misktake Suite")
}
