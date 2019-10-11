package PointontheGraphics_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "Algorithm-Datastruct/golang/PointontheGraphics"
)

var _ = Describe("PointontheGraphics", func() {
	Context("check if the point on the graphics", func() {

		It("Add Round should be right", func() {
			cooridinate := NewCoordinates(5, 5)
			round1 := cooridinate.NewRound(3, 2, 0.7, 1)
			round2 := cooridinate.NewRound(4, 2, 3, 2)
			round3 := cooridinate.NewRound(3, 3, 0.7, 3)
			round4 := cooridinate.NewRound(4, 3, 1, 4)

			cooridinate.AddRound(*round1)
			cooridinate.AddRound(*round2)
			cooridinate.AddRound(*round3)
			cooridinate.AddRound(*round4)

			resultround := cooridinate.FindRoundsonPoint(3, 2)

			Expect(resultround).Should(Equal(map[int]struct{}{1: {}, 2: {}}))


			resultround1 := cooridinate.FindRoundsonPoint(3.9, 2.8)
			Expect(resultround1).Should(Equal(map[int]struct{}{4: {}, 2: {}}))

			resultround2 := cooridinate.FindRoundsonPoint(1, 1)
			Expect(resultround2).Should(Equal(map[int]struct{}{}))


		})
	})
})
