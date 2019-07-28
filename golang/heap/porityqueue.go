package heap

import (
	"strconv"
)

func(heap *Heap)MergeArray(data [][]int)string{
	grouplength := len(data)
	HeapCapcity := 6
	// build the SmallHeap
	for arrayindex:=0;arrayindex<grouplength;arrayindex++{
		for i := 0; i<HeapCapcity/grouplength;i++{
			heap.SmallTopInsert(data[arrayindex][i])
			data[arrayindex][i] = 0
		}
	}

	// 把拍好的数字放到result中
	var MergeString string
	var arraynotemptyindex int
	for {
		MergeString += strconv.Itoa(heap.Data[1])
		// 应该需要更好的方法来判断应该使用哪个数组
		if heap.Data[1]%2 == 0{
			heap.RemoveSmallTop()
			heap.SmallTopInsert(data[0][ArrayIndex(data[0])])
			data[0][ArrayIndex(data[0])] = 0
		}else {
			heap.RemoveSmallTop()
			heap.SmallTopInsert(data[1][ArrayIndex(data[1])])
			data[1][ArrayIndex(data[1])] = 0
		}
		if stop,index := ArearrayEmpty(data);stop == true{
			arraynotemptyindex = index
			break
		}
	}

	// 把小顶堆的放入result
	heap.Sort()
	for _, data:=range heap.Data[1:]{
		value := strconv.Itoa(data)
		MergeString += value
	}

	// 把有剩余元素的数组放进去
	for _,value := range data[arraynotemptyindex]{
		if value != 0{
			data := strconv.Itoa(value)
			MergeString += data
		}
	}
	return MergeString
}


func ArrayIndex(array []int)int{
	for i:= 0;i<len(array);i++{
		if array[i] != 0{
			// not null before zero
			return i
		}
	}
	return -1
}

// find the unempty array and it's index
func ArearrayEmpty(data [][]int)(bool, int){
	results:=[]int{}
	datalength := len(data)
	for i:= 0;i<datalength;i++{
		//if the data is empty
		if data[i][len(data[i])-1]==0{
			results = append(results, 0)
		}else{
			results = append(results, 1)
		}
	}
	unemptynum := 0
	unemptyindex := 0
	for index,result :=  range results{
		if result != 0{
			unemptynum++
			unemptyindex = index
		}
	}
	if unemptynum <= 1{
		return true,unemptyindex
	}
	return false, -1
}

