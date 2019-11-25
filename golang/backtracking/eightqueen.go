package backtracking

import "fmt"

var quene = []int{0,0,0,0,0,0,0,0}
//
//// GetEightQueue get Eight queue array
//func GetEightQueue(row int) {
//	if row == 8{
//		fmt.Println(quene)
//		PrintEightQueue()
//		return
//	}
//	for column := 0;column<8;column++{
//		if IsOk(row,column){
//			quene[row] = column
//			GetEightQueue(row+1)
//		}
//	}
//}
//
//func IsOk(row int, column int) bool {
//
//	leftup,rightup := column-1,column+1
//	for r := row-1; r >= 0; r--{
//			if quene[r] == leftup{
//				return false
//			}
//
//			if quene[r] == rightup{
//				return false
//			}
//
//
//		if quene[r] == column{
//			return false
//		}
//		leftup--
//		rightup++
//	}
//	return true
//}
//
//func PrintEightQueue() {
//	for i := 0;i < 8;i++{
//		for j := 0;j<8;j++{
//			if quene[i]==j{
//				fmt.Print("  Q",)
//			}else{
//				fmt.Printf("  *")
//			}
//		}
//		fmt.Println()
//	}
//	fmt.Println()
//	return
//}




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