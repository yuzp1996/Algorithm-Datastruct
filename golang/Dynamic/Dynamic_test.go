package Dynamic_test

import (
	"Algorithm-Datastruct/golang/Dynamic"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Dynamic", func() {
	It("test dynamic", func() {
		Dynamic.ZerOnePackage()
		Dynamic.ZerOnePackageSimple()
	})
	It("test max value", func() {
		Dynamic.ZeronePackageMaxValue()
	})
	It("minpath", func() {
		Dynamic.MinPath()
	})

})
