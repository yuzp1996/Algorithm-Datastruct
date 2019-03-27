package BinarySearch

func Getfistindex(array []int, target int)int{
	per := 0
	last := len(array)-1
	if array[last]<target || array[per]>target{
		return -1
	}
	for per <= last{
		mid := (per+last)/2
		if array[mid]>target{
			last = mid-1
		}else if array[mid]<target{
			per = mid+1
		}else{
			if mid==0||array[mid-1]!=target{//按照返回来分 什么时候直接返回  什么时候需要再次处理
				return mid
			}else {
				last = mid-1 //需要的值肯定在这里面
			}
		}
	}
	return -1
}

func Getlastindex(array []int, target int)int{
	per := 0
	last := len(array)-1
	if array[last]<target || array[per]>target{
		return -1
	}
	for per <= last{
		mid := (per+last)/2
		if array[mid]>target{
			last = mid-1
		}else if array[mid]<target{
			per = mid+1
		}else{ //变体的 这里是关键
			if mid==len(array)-1||array[mid+1]!=target{//按照返回来分 什么时候直接返回  什么时候需要再次处理
				return mid
			}else {
				per = mid+1 //需要的值肯定在这里面
			}
		}
	}
	return -1
}

//第一个大于等于给定值  往左找
func GetFirstBigger(array []int, target int)int{
	per := 0
	last := len(array)-1
	for per<=last{
		mid := (per+last)/2
		if array[mid]>=target{
			if mid==0||array[mid-1]<target{
				return mid
			}
			last = mid-1
		}else if array[mid]<target {
			per = mid + 1
		}
	}
	return -1
}

//最后一个小于等于给定值 往右找
func Lastsamller(array []int,target int)int{
	per := 0
	last := len(array)-1
	for per<=last{
		mid := (per+last)/2
		if array[mid]<=target{
			if mid == len(array)-1||array[mid+1]>target{
				return mid
			}
			per = mid+1
		}else if array[mid]>target{
			last = mid-1
		}
	}
	return -1
}