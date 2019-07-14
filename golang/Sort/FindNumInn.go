package Sort

//第几大
// 1 2 3 4 5 6       3
// 1 是第一大,  3 是第三大, 6 是第六大
func FindTheNum(array []int, num int) int {
	var data int
	if len(array) == 1 {
		return array[0]
	}
	bigger := GetBigger(array, array[len(array)-1])
	small := GetSmaller(array, array[len(array)-1])
	same := FindSame(array, array[len(array)-1])
	if len(small) >= num {
		data = FindTheNum(small, num)
	} else if len(small) < num && num <= len(same)+len(small) {
		data = array[len(array)-1]
	} else if num > len(same)+len(small) {
		data = FindTheNum(bigger, len(array)-len(same)-len(same))
	}
	return data
}
