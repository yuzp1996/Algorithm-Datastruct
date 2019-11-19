package divide

func MergeSort(data []int)[]int{

	length := len(data)
	if length <= 1{
		return data
	}
	index := length/2

	larray := MergeSort(data[:index])
	rarray := MergeSort(data[index:])
	result := Merge(larray,rarray)
	return result
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