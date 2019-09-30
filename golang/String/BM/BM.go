package BM

import "fmt"

func BMfindstring(Main, Pattern string) bool {
	mainlength := len(Main)
	patternlength := len(Pattern)
	start := 0

	if len(Main) < len(Pattern) {
		return false
	}

	for {
		if start+patternlength < mainlength {
			fmt.Printf("start is %v", start)
			fmt.Printf("Main[start:start+patternlength] is %v  Pattern is %v\n", Main[start:start+patternlength], Pattern)
			if Main[start:start+patternlength] == Pattern {
				return true
			}
			forward := FindForward(Main[start:start+patternlength], Pattern)
			if forward < 0 {
				start += 1
			} else {
				start += forward
			}

		} else {
			break
		}
	}

	return false
}

func FindForward(Main, Pattern string) int {

	var notmatchindex, notmatcharinpattern int
	for i := len(Pattern) - 1; i > 0; i-- {
		if Main[i] != Pattern[i] {
			notmatchindex = i
			break
		}
	}
	notmatchchar := Main[notmatchindex]

	for j := len(Pattern) - 1; j > 0; j-- {
		if string(Pattern[j]) == string(notmatchchar) {
			notmatcharinpattern = j
			break
		}
	}

	fmt.Printf("%v - %v", notmatchindex, notmatcharinpattern)
	return notmatchindex - notmatcharinpattern
}

func GoodSuffix(Main, Pattern string) bool {

	var forwdstep int
	mainlength := len(Main)
	patternlength := len(Pattern)

	fmt.Printf("mainlength is %d\n", mainlength)
	fmt.Printf("patternlength is %d\n", patternlength)

	if patternlength > mainlength {
		return false
	}
	matchendindex := patternlength - 1

	for {
		if matchendindex < mainlength {
			//find good suffix
			patternindex := patternlength - 1
			for index := matchendindex; index > matchendindex-patternlength+1; index-- {
				if patternindex == -1 {
					return true
				}
				if Main[index] == Pattern[patternindex] {
					fmt.Printf("Main[index] is %s  and Patter[pattern] is %s \n", string(Main[index]), string(Pattern[patternindex]))
					patternindex--
				} else {
					goodsuffiixlength := patternlength - (patternindex + 1)
					goodsuffix := Main[index+1 : index+goodsuffiixlength+1]
					fmt.Printf("goodsuffix is %s and goodsuffixlength is %d  matchendindex is %d\n", goodsuffix, goodsuffiixlength, matchendindex)
					forwdstep = Findforwdstep(Pattern, goodsuffix)
					break
				}
			}
		} else {
			return false
		}
		matchendindex += forwdstep
	}

}

//abcdef
//acbde

func Findforwdstep(pattern, goodsuffix string) int {
	if len(goodsuffix) == 0 {
		return 1
	}
	suffixlen := len(goodsuffix)
	startindex := len(pattern) - 1
	for index := startindex; index-suffixlen >= 0; index-- {
		//should return the latter pattern index
		if pattern[index-suffixlen:index] == goodsuffix {
			return startindex - index
		}
	}
	return len(pattern)
}
