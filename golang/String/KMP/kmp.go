package KMP

func MainSkeleTon(main,pattern string)bool{
	mainlength := len(main)
	patternlength := len([]rune(pattern))
	failurearray := FailuerFunction(pattern)

	j := 0
	//每次都是大的不动，小的动起来
	for index := 0;index<mainlength;index++{

		for(j>0 && main[index]!= pattern[j]){
			j = failurearray[j -1] +1
		}
		if main[index] == pattern[j]{
			j++
		}
		if j == patternlength{
			return true
		}
	}
	return false
}

func FailuerFunction(pattern string)[]int{
	lenpattern := len([]rune(pattern))
	next := []int{}
	next = append(next,-1)
	maxlength := -1
	for index := 1; index < lenpattern - 1; index++{
		for(maxlength > -1 && pattern[index] != pattern[maxlength+1]){
			maxlength = next[maxlength]
		}
		if pattern[index] == pattern[maxlength+1]{
			maxlength++
		}
		next = append(next,maxlength)
	}
	return next
}




//func TestFilureFunction(pattern string)[]int{
//	next := []int{}
//	next = append(next,-1)
//	maxlength := -1
//	lengthpattern := len([]rune(pattern))
//	for index := 1; index < lengthpattern-1;index++{
//		for(maxlength > -1 && pattern[index] != pattern[maxlength+1]){
//			maxlength = next[maxlength]
//		}
//		if pattern[maxlength+1] == pattern[index]{
//			maxlength++
//		}
//		next = append(next,maxlength)
//	}
//	return next
//}




//// next是反映的 前缀结尾字符下标：最长可匹配前缀字串结尾字符的下标
//func Failurefunction(pattern string)[]int{
//	// if the index value is 0  it should be 1  because forward need use it
//	lengthpattern := len([]rune(pattern))
//	next := []int{}
//
//	//pattern中 一个子串  前缀与后缀 匹配的最大长度
//	var maxlength = -1
//
//	//第一个肯定是没有 没有做最大最小前后缀的匹配
//	next = append(next,-1)
//	for index := 1; index < lengthpattern-1; index++ {
//
//		for (maxlength > -1 && pattern[index]!=pattern[maxlength+1]){
//			maxlength = next[maxlength]
//		}
//
//		if pattern[index] == pattern[maxlength+1]{
//			maxlength++
//		}
//		next = append(next,maxlength)
//
//		log.Println(next)
//
//	}
//	return next
//}












// next是反映的 某种长度的前缀字符 最大匹配后缀的第一个字符的位置
// if is dont match it will be -1
// the result should be
// for example:
//a   b  a  b  a  b  z  a  b  a  b  a  b
//-1 -1  0  2  2  2  0  8  8  8  8  8
//func Failfunction(pattern string)[]int{
//
//}









//func Failurefunction(pattern string)[]int{
//	// if the index value is 0  it should be 1  because forward need use it
//	lengthpattern := len([]rune(pattern))
//	next := []int{}
//
//	var maxlength int
//
//	for index := 1; index < lengthpattern; index++ {
//		//每次一个数的过滤
//		//如果前一个值的有最长匹配后缀 && 要匹配的这个值 和 前一个值最长匹配后缀的下一个值不同
//		for (maxlength > 0 && pattern[index]!=pattern[maxlength]){
//			//往回找那种可以匹配的子集
//			maxlength = next[maxlength-1]
//		}
//		//看要匹配的这个值和最长匹配的下一个值是否相同 相同就加一
//		if pattern[index] == pattern[maxlength]{
//			maxlength++
//		}
//		next = append(next,maxlength)
//
//	}
//	log.Println(next)
//	return next
//}

//a b a b a b z a b a b a b

//0 0 0 3 2 2 0 0 1



//func FailFunction(pattern string)[]int{
//	patternlength := len(pattern)
//
//}







//"ababab"  最长可匹配后缀
//0 "a"      -1
//1 "ab"	   -1
//2 "aba"	   2
//3 "abab"   2
//4 "ababa"  2
//
//[]int{-1,-1,2,2,2}


