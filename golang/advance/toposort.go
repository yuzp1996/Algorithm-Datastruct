package advance

import (
	"Algorithm-Datastruct/golang/Graph/ListGraph"
	"fmt"
	"Algorithm-Datastruct/golang/Queue/arrayqueue"
)

//0 -> 1 -> 2 ->  3
//
//
// 4 < - 5
//

func ReadyforData()*ListGraph.ListGraph{
	graphlist := ListGraph.NewListGraph(10)
	graphlist.InsertNode(0)
	graphlist.InsertNode(1)
	graphlist.InsertNode(2)
	graphlist.InsertNode(3)
	graphlist.InsertNode(4)
	graphlist.InsertNode(5)
	graphlist.AddEdge(0,4)
	graphlist.AddEdge(1,2)
	graphlist.AddEdge(2,3)
	graphlist.AddEdge(4,5)

	return graphlist
}


func Start(){

	graphlist := ReadyforData()

	//整理入度
	Intime := make([]int, 6,6)
	for index,_ := range graphlist.Data{
		next := graphlist.Data[index].Next
		for{
			if next != nil{
				Intime[next.Value] += 1
				next = next.Next
			}else{
				break
			}
		}
	}


	//把入度为0的先塞进来
	zeroqueue := arrayqueue.NewArrayQueue(10)
	//result是某个为0的index
	result := FindZeroIndex(Intime)
	for _,v := range result{
		zeroqueue.Enqueue(v)
	}

	//处理入度为0的 并且不停的填充队列 直到队列为空
	for {
		if(zeroqueue.Candequeue()) {
			zerovalue := zeroqueue.Dequeue() //入度为0的
			fmt.Printf(" %d ->",graphlist.Data[zerovalue.(int)].Value)
			next := graphlist.Data[zerovalue.(int)].Next
			if(next != nil){
				Intime[next.Value] -= 1
				if Intime[next.Value] == 0{
					zeroqueue.Enqueue(next.Value)
				}
			}

		}else {
			break
		}
	}
}


func FindZeroIndex(Intime []int)([]int){
	result := []int{}
	for i, _ := range Intime{
		if Intime[i] == 0{
			result = append(result,i)
		}
	}
	return result
}





func DFStoposort(){
	graphlist := ReadyforData()
	graphlist.PrintGraph()

	newgraphlist := ListGraph.NewListGraph(10)
	newgraphlist.InsertNode(0)
	newgraphlist.InsertNode(1)
	newgraphlist.InsertNode(2)
	newgraphlist.InsertNode(3)
	newgraphlist.InsertNode(4)
	newgraphlist.InsertNode(5)

	for _,value := range graphlist.Data{
		if value.Next !=nil{
			newgraphlist.AddEdge(value.Next.Value,value.Value)
		}
	}
	newgraphlist.PrintGraph()

	visited := []bool{
		false,false,false,false,false,false,
	}

	for i, _ := range newgraphlist.Data{
		if !visited[i]{
			visited[i] = true
			dfs(i,newgraphlist,visited)
		}
	}
}


func dfs(i int,listgraph *ListGraph.ListGraph,visited []bool){
	if !visited[i]{
		visited[i] = true
		dfs(listgraph.Data[i].Value,listgraph,visited)
	}
	fmt.Printf("visited is %v",visited)
	fmt.Printf(" %d ->", i)
}