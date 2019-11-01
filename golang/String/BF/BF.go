package BF

import "log"

func BFfindstring(Main, Pattern string) bool {
	patternlength := len(Pattern)
	mainlength := len(Main)
	for i := 0; i <= mainlength-patternlength; i++ {
		if Main[i:i+patternlength] == Pattern {
			return true
		}
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

	//不能太想当然的写代码  应该先用草图实现了 自己心里真的明白了才行
	for {
		log.Printf("the matched is %v", Main[matchfirstindex:matchlastindex])
		if mainvalue != patternvalue {
			if matchlastindex < len(Main) {
				mainvalue -= int32(Main[matchfirstindex])
				mainvalue += int32(Main[matchlastindex])
				matchfirstindex++
				matchlastindex++
			} else {
				return false
			}
		} else {
			if Main[matchfirstindex:matchlastindex] == Pattern {
				return true
			} else {
				if matchlastindex < len(Main) {
					matchfirstindex++
					matchlastindex++
				} else {
					return false
				}
			}
		}
	}
}

func hash(StringData string) int32 {
	var result int32
	for _, value := range StringData {
		result += value
	}
	return result

}
