package backtracking

import "fmt"

var quene = []int{0,0,0,0,0,0,0,0}

func EightQueue(row int){
	if row == 8{
		PrintEightQueue()
		return
	}
	for column :=0; column < 8; column++{
		if IsOk(row,column){
			quene[row] = column
			EightQueue(row+1)
		}
	}
}

func IsOk(rows int, column int) bool{
	lefttup,rightup := column-1, column+1
	for row := rows-1;row >= 0; row--{
		if quene[row] == column{
			return false
		}
		if lefttup>=0{
			if quene[row] == lefttup{
				return false
			}
		}
		if rightup<8{
			if quene[row] == rightup{
				return  false
			}
		}

		lefttup--
		rightup++
	}
	return true

}

func PrintEightQueue(){
	for row :=0 ;row<8;row++{
		for colume := 0;colume<8;colume++{
			if quene[row] == colume{
				fmt.Printf("  Q")
			}else{
				fmt.Printf("  *")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

var LoadMaxWeight int

func ZerOnePackage(index int, currentweight int,weightarray []int,totalnum int,loadmaxwight int){

	//装满了和全装进去了 就是最多的表现了
	if index == totalnum-1 || currentweight == loadmaxwight{
		if currentweight > LoadMaxWeight{
			LoadMaxWeight = currentweight
		}
		fmt.Println(LoadMaxWeight)
		return
	}
	//不考虑当前的这个数 直接考虑当前的下一个
	ZerOnePackage(index+1,currentweight,weightarray,totalnum,loadmaxwight)

	//考虑当前的这个数
	if currentweight+weightarray[index] <= loadmaxwight{
		ZerOnePackage(index+1, currentweight+weightarray[index],weightarray,totalnum,loadmaxwight)
	}
}




var CanloadWeight int

var PackageMaxWeight  = 10
var MaxNum = 8
var ItemsWeight = []int{2,1,1,2,2,3,4,5}

var cheap = [][]bool{
		[]bool{false,false,false,false,false,false,false,false,false,false},
		[]bool{false,false,false,false,false,false,false,false,false,false},
		[]bool{false,false,false,false,false,false,false,false,false,false},
		[]bool{false,false,false,false,false,false,false,false,false,false},
		[]bool{false,false,false,false,false,false,false,false,false,false},
		[]bool{false,false,false,false,false,false,false,false,false,false},
		[]bool{false,false,false,false,false,false,false,false,false,false},
		[]bool{false,false,false,false,false,false,false,false,false,false},
	}

//SimpleZerOnePackage  use it by SimpleZerOnePackage(0,0)
func SimpleZerOnePackage(index int,currentweight int){

	if index == MaxNum - 1||currentweight == PackageMaxWeight{
		if currentweight > CanloadWeight{
			CanloadWeight = currentweight
		}
		fmt.Println(CanloadWeight)
		fmt.Println(cheap)
		return
	}
	if cheap[index][currentweight]{
		return
	}
	cheap[index][currentweight] = true
	SimpleZerOnePackage(index+1, currentweight)
	if currentweight+ItemsWeight[index]<= PackageMaxWeight{
		SimpleZerOnePackage(index+1,currentweight+ItemsWeight[index])
	}
}



















