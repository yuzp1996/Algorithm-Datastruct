package Greedy

import (
	. "Algorithm-Datastruct/golang/Stack/stackarray"
)

func FindTheLatest(data int, time int)(result int){

	stack := NewStackArray()

	resultarray := []int{}
	for data > 0{
		stack.Push(data%10)
		data = data/10
	}
	for !stack.IsEmpty(){
		resultarray = append(resultarray,stack.Pop().(int))
	}
	var really []int
	for t:=0;t<time;t++{
		really = []int{}
		for i := 0;i<len(resultarray)-1;i++{
			if resultarray[i] > resultarray[i+1]{
				really = append(really,resultarray[:i]...)
				really = append(really,resultarray[i+1:]...)
				resultarray = really
				break
			}
		}
	}
	result = ConstructNum(really)
	return result
}




func ConstructNum(data []int)int{
	tennum, result := 0,0
	for i := len(data)-1;i>=0;i--{
		result += data[i]*power(tennum)
		tennum += 1
	}
	return result
}

func power(num int)int{
	result := 1
	for i:=0;i<num;i++{
		result = result * 10
	}
	return  result
}