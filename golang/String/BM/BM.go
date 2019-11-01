package BM

import (
	"fmt"
	"regexp"
)

func BMfindstring(Main, Pattern string) bool {
	mainlength := len(Main)
	patternlength := len(Pattern)
	start := 0

	if len(Main) < len(Pattern) {
		return false
	}

	for {
		if start+patternlength <= mainlength {
			if Main[start:start+patternlength] == Pattern {
				return true
			}
			forward := FindForward(Main[start:start+patternlength], Pattern)
			if forward < 0 {
				start += 1
			} else {
				start += forward
			}
			fmt.Printf("forward is %d \n", forward)

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

	fmt.Printf("%v - %v = ", notmatchindex, notmatcharinpattern)
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

	substring, suffix, prefix := Getneededarray(Pattern)

	fmt.Printf("get the substring is %v suffix is %v  and the prefix is %v\n", substring, suffix, prefix)

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
					forwdstep = Findforwdstep(Pattern, goodsuffix, suffix, prefix)
					fmt.Printf("forwdstep is %v \n", forwdstep)
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

func Findforwdstep(pattern, goodsuffix string, suffix []int, prefix []bool) int {

	lensuffix := len(goodsuffix)
	if lensuffix == 0 {
		return 1
	}
	if suffix[lensuffix-1] != -1 {
		return len(pattern) - len(goodsuffix) - suffix[lensuffix-1]

	}
	for j := len(goodsuffix); j >= 0; j-- {
		if prefix[j] {
			return len(pattern) - j
		}
	}

	return len(pattern)
}

func Getneededarray(pattern string) (childpatternwithlength []string, suffix []int, prefix []bool) {
	lengthpattern := len(pattern)
	childpatternwithlength = make([]string, lengthpattern-1)
	suffix = make([]int, lengthpattern-1)
	prefix = make([]bool, lengthpattern-1)
	j := 0
	for i := lengthpattern - 1; i > 0; i-- {
		childpatternwithlength[j] = pattern[i:]
		j++
	}

	for parentindex := 0; parentindex <= lengthpattern-2; parentindex++ {
		for childindex := parentindex; childindex >= 0; childindex-- {
			if pattern[childindex] == pattern[len(pattern)-1-(parentindex-childindex)] {
				suffix[parentindex-childindex] = childindex
				if childindex == 0 {
					prefix[parentindex-childindex] = true
				}
			} else {
				continue
			}
		}
	}

	fmt.Printf("now now suffix is %v \n", suffix)
	for i := 0; i < len(suffix); i++ {
		if suffix[i] == 0 && !prefix[i] {
			suffix[i] = -1
		}
	}
	return childpatternwithlength, suffix, prefix
}

//below code just for test
//func GetNameWithOutVersion(name string)string{
//	return name[:strings.Index(name,".")]
//}

func GetNameWithRE(s string) string {
	rex, _ := regexp.Compile(`\.[\d]+\.[\d]+\.[\d]+$`)
	result := rex.FindAllStringIndex(s, -1)
	if len(result) > 0 {
		fmt.Printf("when there is match result is %v\n", s[:result[0][0]])
		return s[:result[0][0]]
	}
	return "nosuchthing"
}

func GetNameWithOutVersion(name, versionName string) bool {
	rex, _ := regexp.Compile(`\.[\d]+\.[\d]+\.[\d]+$`)
	result := rex.FindAllStringIndex(versionName, -1)
	if len(result) > 0 {
		return versionName[:result[0][0]] == name
	}
	return false
}
