package Sort

//比较的应该是某个索引下面的值  不是某个值  应该抓到索引
func Select_Sort(array []int)[]int{
	for index,_ := range array{
		for sindex,_ := range array[index+1:]{
			if array[index] >= array[sindex+index+1]{
				array[sindex+index+1],array[index] = array[index],array[sindex+index+1]
			}
		}
	}
	return array
}

//从右边选最小的，放到左边来
// 2 3 7 5 4 3 2

// |
// 2   3 7 5 4 3 2

//   |
// 2 3   7 5 4 3 2
// 2 2   7 5 4 3 3

//    |
//2 2 7   5 4 3 3
//2 2 3   5 4 7 3

//      |
//2 2 3 5   4 7 3
//2 2 3 3   4 7 5

//        |
//2 2 3 3 4   7 5
//2 2 3 3 4   7 5

//          |
//2 2 3 3 4 7   5
//2 2 3 3 4 5   7
