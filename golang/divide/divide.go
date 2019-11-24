package divide

var count int

func MergeSort(data []int)( []int, int){


	length := len(data)
	if length <= 1{
		return data,count
	}
	index := length/2
	larray,_ := MergeSort(data[:index])
	rarray,_ := MergeSort(data[index:])
	result := Merge(larray,rarray)
	return result, count
}

func Merge(larray []int, rarray []int)[]int{
	result := []int{}
	lindex, rindex := 0,0
	llenght, rlength := len(larray), len(rarray)
	for lindex < llenght && rindex < rlength{
		if larray[lindex] < rarray[rindex]{
			result = append(result,larray[lindex])
			lindex++
		}else{
			result = append(result,rarray[rindex])
			count += llenght - lindex
			rindex++
		}
	}
	if rindex < rlength{
		result = append(result,rarray[rindex:]...)
	}else{
		result = append(result,larray[lindex:]...)
	}
	return result
}