package backtracking

import "fmt"

var Path = [][]int{
	[]int{1,1,1,1,1}, //0
	[]int{1,1,1,1,1}, //1
	[]int{1,1,1,1,1}, //2
	[]int{1,1,1,1,1}, //3
	[]int{1,1,1,1,1}, //4
}

var MinPathValue int = 100

func MinPath(row int, column int,pathlength int){
	if row == 4 && column ==4{
		if pathlength < MinPathValue{
			MinPathValue = pathlength
		}
		fmt.Println(MinPathValue)
		return
	}
	if row+1 <= 4{
		MinPath(row+1,column,pathlength+Path[row+1][column])
	}
	if column+1 <= 4{
		MinPath(row,column+1,pathlength + Path[row][column+1])
	}

}