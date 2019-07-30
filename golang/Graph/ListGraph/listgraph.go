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
	Prev[source] = -1
	var result []int
	for{
		if Queue.Candequeue() {
			queuehead := Queue.Data[Queue.Head].(int)
			if queuehead != target {
				Visited[queuehead] = true
				childnodevalues := graph.GetNextList(queuehead)
				Queue.Dequeue()
				for _, value := range childnodevalues {
					if !Visited[value]{
						if Prev[value]==-1{
							Prev[value] = queuehead
						}
						Queue.Enqueue(value)
					}
				}
			}else{
				result = Getpath(Prev,source,target)
				return true, result
			}
		}
	}
}



func Getpath(Prev []int,source,target int)(result []int){
	if(Prev[target]!=-1 && target != source){
		result = Getpath(Prev,source,Prev[target])
	}
	result = append(result,target)
	return
}


func ValueInQueue(data []interface{}, value int)bool{
	for index, _ := range data{
		interval,_ := data[index].(int)
		if interval == value{
			return true
		}
	}
	return false
}



func (graph *ListGraph)GetNextList(source int)[]int{
	Result := []int{}
	for _,node := range graph.Data{
		if node.Value == source{
			for{
				node = node.Next
				if node != nil{
					Result = append(Result,node.Value)
				}else{
					break
				}
			}
		}
	}
	return Result
}
