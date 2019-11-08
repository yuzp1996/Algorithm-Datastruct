package KMP

import "log"

func MainSkeleTon(main,pattern string)bool{
	mainlength := len(main)
	failurearray := Failurefunction(pattern)
	patternlength := len(pattern)
	var forward = 1
	for mainindex := 0; mainindex <= mainlength - patternlength; mainindex += forward{
		for patternindex:= 0;patternindex < patternlength;patternindex++{
			if main[mainindex + patternindex] == pattern[patternindex]{
				if patternindex == patternlength-1{
					return true
				}
				continue
			}

			forward = failurearray[patternindex]
			break
		}
	}
	return false
}

func Failurefunction(pattern string)[]int{
	// if the index value is 0  it should be 1  because forward need use it
	log.Printf("len(pattern) is %v", len(pattern))
	next := []int{len(pattern)}
	var maxlength int
	for index := 1; index < len(pattern); index++ {
		//每次一个数的过滤


		//如果前一个值的有最长匹配后缀 && 要匹配的这个值 和 前一个值最长匹配后缀的下一个值不同
		for (maxlength > 0 && pattern[index]!=pattern[maxlength]){
			//往回找那种可以匹配的子集
			maxlength = next[maxlength-1]
		}
		//看要匹配的这个值和最长匹配的下一个值是否相同 相同就加一
		if pattern[index] == pattern[maxlength]{
			maxlength++
		}
		next[index] = maxlength


	}
	return next
}



//"ababab"  最长可匹配后缀
//0 "a"      -1
//1 "ab"	   -1
//2 "aba"	   2
//3 "abab"   2
//4 "ababa"  2
//
//[]int{-1,-1,2,2,2}


