package ListGraph

import "fmt"
import "Algorithm-Datastruct/golang/Queue/arrayqueue"

type node struct {
	Value int
	Next  *node
}

func NewNode(value int) *node {
	return &node{
		Value: value,
		Next:  nil,
	}
}

type ListGraph struct {
	Cap  int
	Data []*node
}

func NewListGraph(cap int) *ListGraph {
	return &ListGraph{
		Cap:  cap,
		Data: []*node{},
	}
}

func (graph *ListGraph) InsertNode(value int) {
	if len(graph.Data) >= graph.Cap {
		return
	}
	graph.Data = append(graph.Data, NewNode(value))
}

func (graph *ListGraph) AddEdge(source, target int) {
	for index, node := range graph.Data {
		if node.Value == source {
			nextnode := NewNode(target)
			nextnode.Next = graph.Data[index].Next
			graph.Data[index].Next = nextnode
		}
	}
}

func (graph *ListGraph) PrintGraph() {
	for _, value := range graph.Data {
		for {
			if value != nil {
				fmt.Printf("%d -> ", value.Value)
				value = value.Next
			} else {
				break
			}
		}
		fmt.Print("nil\n")
	}
}

func (graph *ListGraph) BFS(source, target int) (bool, []int) {
	if source == target {
		return true, []int{}
	}
	Queue := arrayqueue.NewArrayQueue(10)
	Prev := []int{}
	Visited := []bool{}
	for i := 0; i < len(graph.Data); i++ {
		Visited = append(Visited, false)
		Prev = append(Prev, -1)
	}


	Queue.Enqueue(source)


	return true, []int{1, 2, 3}

}
