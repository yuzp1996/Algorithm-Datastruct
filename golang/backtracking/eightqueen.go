package backtracking

import (
	"fmt"
)

var quene = []int{0,0,0,0,0,0,0,0}

// GetEightQueue get Eight queue array
func GetEightQueue(row int) {
	if row == 8{
		fmt.Println(quene)
		PrintEightQueue()
		return
	}
	for column := 0;column<8;column++{
		if IsOk(row,column){
			quene[row] = column
			GetEightQueue(row+1)
		}
	}
}

func IsOk(row int, column int) bool {

	leftup,rightup := column-1,column+1
	for r := row-1; r >= 0; r--{
			if quene[r] == leftup{
				return false
			}

			if quene[r] == rightup{
				return false
			}


		if quene[r] == column{
			return false
		}
		leftup--
		rightup++
	}
	return true
}

func PrintEightQueue() {
	for i := 0;i < 8;i++{
		for j := 0;j<8;j++{
			if quene[i]==j{
				fmt.Print("  Q",)
			}else{
				fmt.Printf("  *")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	return
}




