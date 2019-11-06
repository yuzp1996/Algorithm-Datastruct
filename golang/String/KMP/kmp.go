package KMP

func MainSkeleTon(main,pattern string)bool{
	mainlength := len(main)
	failurearray := Failurefunction(pattern)
	patternlength := len(pattern)
	forward := 1
	for mainindex := 0;mainindex < mainlength - patternlength;mainindex = mainindex+forward{
		for patternindex:= 0;patternindex < patternlength;patternindex++{
			if main[mainindex + patternindex] == pattern[patternindex]{
				if mainindex + patternindex == patternlength{
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
	return []int{}
}
