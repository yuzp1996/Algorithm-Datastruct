package BinarySearch

// if exist,return the first index
// if not, return 0
func Binarysearch(searcharray []int,num int)int{
	last := len(searcharray)-1
	per:=0
	if searcharray[last]<num || searcharray[per]>num{
		return -1
	}
	for last >= per{
		mid := (last+per)/2
		if searcharray[mid]>num{
			last = mid-1
		}else if searcharray[mid]<num{
			per = mid+1
		}else if searcharray[mid] == num{
			return mid
		}
	}
	return -1
}

// find 5.5

// 1      3 4 5 6
//per(0)       last(4)

//       5       6
//		per(3) last(4)

//       5       6
//			  last(4)
//			  per(4)


//       5       6
//		last(3)	       XXXXX
//			  per(4)   XXXXX
//      退出循环






func RecursionBinarySearch(array []int,num int)int{
	last := len(array)-1
	per:=0
	if array[last]<num || array[per]>num{
		return -1
	}
	result := recursion(array,per,last,num)
	return result
}

func recursion(array []int, per,last,num int)int{
	//这个是终止条件
	if per>last{
		return -1
	}
	mid := (per+last)/2
	if array[mid] == num{
		return mid
	}else if array[mid]>num{
		return recursion(array,per,mid-1,num)
	}else if array[mid]<num{
		return recursion(array,mid+1,last,num)
	}else{
		return -1
	}
}