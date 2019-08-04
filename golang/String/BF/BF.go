package BF

import "fmt"

func BFfindstring(Main, Pattern string) bool {
	patternlength := len(Pattern)
	mainlength := len(Main)
	for i := 0; i <= mainlength-patternlength; i++ {
		if Main[i:i+patternlength] == Pattern {
			return true
		}
	}
	for _, value := range "abcdefghijklmnopqrstuvwxyz" {
		fmt.Printf("%d\n", value-97)
	}
	return false
}

func RKfindstring(Main, Pattern string) bool {
	patternlength := len(Pattern)
	mainlength := len(Main)
	if mainlength < patternlength {
		return false
	}
	patternvalue := hash(Pattern)

	mainvalue := hash(Main[:patternlength])

	matchlastindex := patternlength
	matchfirstindex := 0

	for {
		if matchlastindex <= len(Main) {
			if mainvalue != patternvalue {
				matchfirstindex++
				matchlastindex++
				mainvalue -= int32(Main[matchfirstindex])
				mainvalue += int32(Main[matchlastindex])
			} else {
				// if it has the equal val
				if Main[matchfirstindex:matchlastindex] == Pattern {
					return true
				} else {
					matchfirstindex++
					matchlastindex++
				}
			}
		}
		return false
	}
}

func hash(StringData string) int32 {
	var result int32
	for _, value := range StringData {
		result += (value - 97)
	}
	return result

}
