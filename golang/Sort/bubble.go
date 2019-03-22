package Sort

//比较的应该是某个索引下面的值  不是某个值  应该抓到索引
func Bubble(array []int)[]int{
	for index,_ := range array{
		for sindex,_ := range array[index+1:]{
			if array[index] >= array[sindex+index+1]{
				array[sindex+index+1],array[index] = array[index],array[sindex+index+1]
			}
		}
	}
	return array
}







// 2 4 1 6 4
//