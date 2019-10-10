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
			round1 := cooridinate.NewRound(3, 2, 1, 1)
			round2 := cooridinate.NewRound(3, 2, 1, 2)
			round3 := cooridinate.NewRound(3, 3, 1, 3)

			cooridinate.AddRound(*round1)
			cooridinate.AddRound(*round2)
			cooridinate.AddRound(*round3)

			resultround := cooridinate.FindRoundsonPoint(3, 2)
			//Expect(cooridinate.XRounds).Should(Equal(
			//
			//	map[int][]Round{2:[]Round{Round{3,2,1,1}},3:[]Round{Round{3,2,1,1}}},
			//
			//	))
			//Expect(cooridinate.YRounds).Should(Equal(
			//
			//	map[int][]Round{2:[]Round{Round{3,2,1,1}},3:[]Round{Round{3,2,1,1}}},
			//
			//))
			Expect(resultround).Should(Equal(map[int]struct{}{1: {}, 2: {}, 3: {}}))

		})
	})
})
