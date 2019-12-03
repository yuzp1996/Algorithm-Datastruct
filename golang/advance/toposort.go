package advance

import (
	"Algorithm-Datastruct/golang/Graph/ListGraph"
	"fmt"
)

//1 -> 2 ->  3
//
//4 -> 5
//

func ReadyforData()*ListGraph.ListGraph{
	graphlist := ListGraph.NewListGraph(10)
	graphlist.InsertNode(0)
	graphlist.InsertNode(1)
	graphlist.InsertNode(2)
	graphlist.InsertNode(3)
	graphlist.InsertNode(4)
	graphlist.InsertNode(5)
	graphlist.AddEdge(0,1)
	graphlist.AddEdge(1,2)
	graphlist.AddEdge(2,3)
	graphlist.AddEdge(4,5)
	graphlist.AddEdge(0,4)
	//graphlist.AddEdge(5,0)

	return graphlist
}


func Start(){

	graphlist := ReadyforData()
	graphlist.PrintGraph()

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

	for {
		if(FindIfExsitInOverOne(Intime)) { //仍然有入度大于1的
			for i,_ := range Intime{
				if Intime[i] == 0{
					fmt.Printf(" %d ->",graphlist.Data[i].Value)
					next := graphlist.Data[i].Next
					for(next != nil){
						Intime[next.Value] -= 1
						next = next.Next
					}
				}
			}
		}else {
			return
		}
	}
}


func FindIfExsitInOverOne(Intime []int)(bool){
	for i, _ := range Intime{
		if Intime[i] > 0{
			return  true
		}
	}
	return false
}
