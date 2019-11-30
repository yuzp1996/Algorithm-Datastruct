package Dynamic

import "fmt"

var ItemsWeight = []int{1,2,1,2,1}
var MaxNum = 5
var MaxLoad = 10
var Array = [][]bool{
	[]bool{false,false,false,false,false,false,false,false,false,false,false},
	[]bool{false,false,false,false,false,false,false,false,false,false,false},
	[]bool{false,false,false,false,false,false,false,false,false,false,false},
	[]bool{false,false,false,false,false,false,false,false,false,false,false},
	[]bool{false,false,false,false,false,false,false,false,false,false,false},
}

func ZerOnePackage(){
	Array[0][0] = true
	if ItemsWeight[0] <= MaxLoad{
		Array[0][ItemsWeight[0]] = true
	}
	for row:= 1;row<MaxNum;row++{
		for column := 0;column<MaxLoad;column++{
			if Array[row-1][column]{
				Array[row][column] = true
				if column+ItemsWeight[row]<=MaxLoad{
					Array[row][column+ItemsWeight[row]] = true
				}
			}
		}
	}
	fmt.Println(Array[MaxNum-1])
	return
}



var simplearray = []bool{
	false,false,false,false,false,false,false,false,false,false,false,
}
func ZerOnePackageSimple(){
	simplearray[0] = true
	if ItemsWeight[0] <= MaxLoad{
		simplearray[ItemsWeight[0]] = true
	}
	for i:= 1;i < MaxNum;i++{
		for j := MaxLoad;j>=0;j--{
			if simplearray[j] && j+ ItemsWeight[i] <= MaxLoad{
				simplearray[j+ ItemsWeight[i]] = true
			}
		}
	}
	fmt.Println(simplearray)

}