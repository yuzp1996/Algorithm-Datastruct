package ListGraph_test

import (
	. "Algorithm-Datastruct/golang/Graph/ListGraph"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListGraph", func() {

	var Graph *ListGraph

	BeforeEach(func(){
		Graph = NewListGraph(6)
		Graph.InsertNode(0)
		Graph.InsertNode(1)
		Graph.InsertNode(2)
		Graph.InsertNode(3)
		Graph.InsertNode(4)
		Graph.InsertNode(5)

		Graph.AddEdge(0, 5)
		Graph.AddEdge(5, 0)
		Graph.AddEdge(1, 3)
		Graph.AddEdge(3, 1)
		Graph.AddEdge(1, 4)
		Graph.AddEdge(4, 1)
		Graph.AddEdge(3, 4)
		Graph.AddEdge(4, 3)
		Graph.AddEdge(5, 2)
		Graph.AddEdge(2, 5)
		Graph.AddEdge(2, 1)
		Graph.AddEdge(1, 2)
		Graph.AddEdge(0, 3)
		Graph.AddEdge(3, 0)
	})

	Context("first find ListGraph should be a good thing", func() {
		It("New List graph", func() {
			Graph.PrintGraph()
			Expect((*Graph).Cap).To(Equal(6))
			Expect((*Graph).Data[1].Next.Next.Value).To(Equal(4))

			exsit, path := Graph.BFS(0, 0)
			Expect(exsit).To(Equal(true))
			Expect(path).To(Equal([]int{}))

			exsit, path = Graph.BFS(0, 1)
			Expect(exsit).To(Equal(true))
			Expect(path).To(Equal([]int{0, 3, 1}))

		})
	})
	Context("second fing ListGraph should be a good thing", func() {
		It("New List graph", func() {
			exsit, path := Graph.BFS(1, 0)
			Expect(exsit).To(Equal(true))
			Expect(path).To(Equal([]int{1, 3, 0}))

		})
	})

})
