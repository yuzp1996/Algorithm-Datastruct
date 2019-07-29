package matrixGraph_test

import (
	. "github.com/onsi/ginkgo"
	. "Algorithm-Datastruct/golang/Graph/matrixGraph"
	. "github.com/onsi/gomega"
	)

var _ = Describe("Graph", func() {
	Context("graph insert vertex and edge should success", func() {
		It("insert vertex should success", func() {
			listgraph := NewListGraph(4)
			listgraph.AddVertex(1)
			listgraph.AddVertex(2)
			listgraph.AddVertex(3)
			listgraph.AddVertex(4)
			Expect(listgraph.Value).To(Equal([]int{1,2,3,4}))
			listgraph.AddWeightEdge(1,2,5)
			listgraph.AddWeightEdge(3,2,4)
			Expect(listgraph.Edgelist).To(Equal([][]int{
				{0,0,0,0},
				{0,0,5,0},
				{0,0,0,0},
				{0,0,4,0},
			}))

		})
	})
})
