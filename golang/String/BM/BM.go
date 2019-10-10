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

	if patternlength > mainlength {
		return false
	}
	mainindex := patternlength - 1

	for {
		if mainindex < mainlength {
			//find good suffix
			patternindex := patternlength - 1
			for index := mainindex; index > mainindex-patternlength; index-- {
				if Main[index] == Pattern[patternindex] {
					patternindex--
				} else {
					goodsuffixlength := patternlength - (patternindex + 1)
					goodsuffix := Main[index+1 : index+goodsuffixlength+1]
					forwdstep = Findforwdstep(Pattern, goodsuffix)
					break
				}
				if patternindex < 0 {
					return true
				}
			}
		} else {
			return false
		}
		mainindex += forwdstep
	}

}

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
		// good suffix should be equal to prefix

	}
	//dsfdafabcdefabc
	//bcdefabc

	return len(pattern)
}

func Getneededarray(pattern string) (childpatternwithlength []string, suffix []int, prefix []bool) {
	lengthpattern := len(pattern)
	childpatternwithlength = make([]string, lengthpattern-1)
	suffix = make([]int, lengthpattern)
	prefix = make([]bool, lengthpattern)
	j := 0
	for i := lengthpattern - 1; i > 0; i-- {
		childpatternwithlength[j] = pattern[i:]
		j++
	}

	return childpatternwithlength, suffix, prefix
}
