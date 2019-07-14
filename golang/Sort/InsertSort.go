package Sort

func Insert_Sort(array []int) []int {
	for index, _ := range array {
		for insert_index := index; insert_index > 0; insert_index-- {
			if array[insert_index] < array[insert_index-1] {
				array[insert_index], array[insert_index-1] = array[insert_index-1], array[insert_index]
			}
		}
	}
	return array
}

//插入排序
//从右边选一个 插到左边来
// 6 6 4 3 32 6 4 3

// |
// 6 6 4 3 32 6 4 3

//   |
// 6 6 4 3 32 6 4 3

//     |
// 6 6 4 3 32 6 4 3
// 4 6 6 3 32 6 4 3

//       |
// 4 6 6 3 32 6 4 3
// 3 4 6 6 32 6 4 3

//         |
// 3 4 6 6 32 6 4 3
// 3 4 6 6 32 6 4 3

//            |
// 3 4 6 6 32 6 4 3
// 3 4 6 6 6 32 4 3

//              |
// 3 4 6 6 6 32 4 3
// 3 4 4 6 6 6 32 3

//                |
// 3 4 4 6 6 6 32 3
// 3 3 4 4 6 6 6 32

//比较的应该是某个索引下面的值  不是某个值  应该抓到索引
func MayBubble(array []int) []int {
	for index, _ := range array {
		for sindex, _ := range array[index+1:] {
			if array[index] >= array[sindex+index+1] {
				array[sindex+index+1], array[index] = array[index], array[sindex+index+1]
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
