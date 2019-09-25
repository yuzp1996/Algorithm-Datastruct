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
			start += forward
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
