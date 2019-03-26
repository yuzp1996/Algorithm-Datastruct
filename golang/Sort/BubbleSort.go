package Sort

func Buuble_Sort(array []int)[]int {
	lenarray := len(array)
	for index := lenarray - 1; index >= 1; index-- {
		//需要记住 这个 array[:index] 就包括了 0->index-1 的所有值
		for sindex, _ := range array[:index] {
			if array[sindex] > array[sindex+1] {
				array[sindex], array[sindex+1] = array[sindex+1], array[sindex]
			}
		}
	}
	return array
}