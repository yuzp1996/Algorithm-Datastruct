package Sort

func merge(Larray []int,Rarray []int)[]int{
	llen,rlen := len(Larray),len(Rarray)
	lindex, rindex := 0, 0
	result := []int{}
	for lindex<=llen-1&&rindex<=rlen-1{
		if Larray[lindex] <= Rarray[rindex]{
			result = append(result,Larray[lindex])
			lindex++
		}else{
			result = append(result,Rarray[rindex])
			rindex++
		}
	}
	if lindex == llen{
		result = append(result,Rarray[rindex:]...)
	}else{
		result = append(result,Larray[lindex:]...)
	}
	return result
}

func Merge_Sort(array []int)[]int{
	if len(array) < 2{
		return array
	}
	n := len(array)/2
	larray := Merge_Sort(array[:n])
	rarray := Merge_Sort(array[n:])
	fianlarray := merge(larray,rarray)
	return fianlarray
}


// 4 3 5 6 7 3 2 4 5 6
// 4 3 5 6 7 | 3 2 4 5 6
// 4 3 5| 6 7       |          3 2 4| 5 6
// 4 | 3 5       6 | 7          3 | 2 4      5 | 6
//     3|5  						2|4
// 4 | 3 5       6 | 7           3 | 2 4     5 | 6
// 4 3 5    6 7   3 2 4    5 6
// 4 3 5  |  6 7      3 2 4  |  5 6
//  3 4 5 6 7|2 3 4 5 6
//2 3 3 4 4 5 5 6 6 7