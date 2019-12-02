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



//var ItemsWeight = []int{1,2,1,2,1}
//var MaxNum = 5
//var MaxLoad = 10

// 有了价值的  多的就是看看留下 如果有重复的值的话 留下价值最大的
var ItemsValue = []int{2,3,4,5,1}
var maxValueArray = [5][10]int{}

func ZeronePackageMaxValue(){
	for row := 0; row<MaxNum; row++{
		for column := 0; column<MaxLoad;column++{
			maxValueArray[row][column]=-1
		}
	}
	fmt.Println(maxValueArray)
	//处理第一个物品
	maxValueArray[0][0] = 0
	if ItemsWeight[0] <= MaxLoad{
		maxValueArray[0][ItemsWeight[0]] = ItemsValue[0]
	}
	for row := 1;row < MaxNum;row++{

		for column := 0; column < MaxLoad;column++{
			if maxValueArray[row-1][column] != -1{
				//这里只看重量 可以放进去
				if column + ItemsWeight[row] <= MaxLoad{
					// 这里要看价值了
					//不加入
					if maxValueArray[row-1][column] > maxValueArray[row][column]{
						maxValueArray[row][column] = maxValueArray[row-1][column]
					}
					// 这里要看价值了
					//加入  上一行的价值要大于下一行 加入了质量后对应的那个价值
					if 	maxValueArray[row-1][column]+ItemsValue[row] > maxValueArray[row][column + ItemsWeight[row]]{
						maxValueArray[row][column + ItemsWeight[row]] = maxValueArray[row-1][column]+ItemsValue[row]
					}
				}
			}
		}
	}

	fmt.Println(maxValueArray)
}

