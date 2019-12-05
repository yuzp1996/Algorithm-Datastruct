package advance_test

import (
	. "github.com/onsi/ginkgo"
	"Algorithm-Datastruct/golang/advance"

)

var _ = Describe("Advance", func() {
	It("test toposort", func() {
		advance.Start()
	})
	It("DFStoposort", func() {
		advance.DFStoposort()
	})

})
