package Dynamic

import "fmt"

var Path = [4][4]int{
	[4]int{1,3,5,9},
	[4]int{2,1,3,4},
	[4]int{5,2,6,7},
	[4]int{6,8,4,3},
}


var ROWS = 4
var COLUMNS = 4
var PathValue = [4][4]int{}



func MinPath(){
	PrintArray(Path)
	for row:=0;row<ROWS;row++{
		for column :=0;column<COLUMNS;column++{
			PathValue[row][column] = -1
		}
	}

	PrintArray(PathValue)
	PathValue[0][0] = Path[0][0]

	for column := 1;column<COLUMNS;column++{
		PathValue[0][column] = PathValue[0][column-1] + Path[0][column]
	}
	PrintArray(PathValue)

	for row := 1;row<ROWS;row++{
		PathValue[row][0] = PathValue[row-1][0] + Path[row][0]
	}
	PrintArray(PathValue)

	for row := 1;row < ROWS;row++{
		for column := 1; column < COLUMNS; column++{
			PathValue[row][column] = MinNum(PathValue[row-1][column],PathValue[row][column-1]) + Path[row][column]
		}
	}
	PrintArray(PathValue)
}


func MinNum(first, second int)int{
	if first>second{
		return second
	}
	return first
}

func PrintArray(Array [4][4]int){
	ROW := len(Array)
	COLUMN := len(Array[0])

	for row := 0;row<ROW;row++{
		for column := 0;column<COLUMN;column++ {
			fmt.Printf("%v ",Array[row][column])
		}
		fmt.Println()
	}

}