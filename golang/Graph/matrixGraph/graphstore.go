package matrixGraph

import "fmt"

//  [1]--3-->[2]
//   | 		  |
//   4        9
//   |		  |
//  [3]--2-->[4]


//  0  1  2  3
//0
//1
//2
//3



type ListGraph struct {
	Num int
	Cap int
	Value []int   // the vertex of this graph
	Edgelist [][]int   // edge between vertex
}

func NewListGraph(cap int)*ListGraph{
	Edgelist := [][]int{}

	for j := 0; j<cap; j++{
		siglearray := make([]int, cap, cap)
		Edgelist = append(Edgelist,siglearray)
	}

	return &ListGraph{
		Num:0,
		Cap:cap,
		Value:[]int{},
		Edgelist:Edgelist,
	}
}

func(graph *ListGraph)AddVertex(value int){
	if graph.Num>=graph.Cap{
		return
	}
	graph.Value = append(graph.Value,value)
	graph.Num++
}

// i and j is same as compute index
func(graph *ListGraph)AddWeightEdge(i,j,weight int){
	fmt.Println(graph.Edgelist)
	fmt.Println(i,j)
	fmt.Println(graph.Edgelist[i][j])
	graph.Edgelist[i][j] = weight
	fmt.Println(graph.Edgelist)
}

