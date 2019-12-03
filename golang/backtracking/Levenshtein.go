package backtracking

import "fmt"

func Stringtoarray(data string)(result []uint8){
	lenstr := len(data)
	for index := 0;index<lenstr;index++{
		result = append(result,data[index])
	}
	fmt.Println(result)
	return result
}

var FirstStr = "yuzhipeng"
var LenFirst = len(FirstStr)
var SecondStr = "yuzhppeeg"
var LenSecond = len(SecondStr)
var MinStep = 1000
func Levenshtein(i int, j int, step int){
	if i == LenFirst||j == LenSecond{
		if i < LenFirst{
			step += LenSecond - j
		}
		if j < LenSecond{
			step += LenFirst - i
		}
		if step < MinStep{
			MinStep = step
		}
		fmt.Println(MinStep)
		return
	}
	if FirstStr[i] == SecondStr[j]{
		Levenshtein(i+1,j+1,step)
	}else{
		Levenshtein(i+1,j,step+1)
		Levenshtein(i,j+1,step+1)
		Levenshtein(i+1,j+1,step+1)
	}
}