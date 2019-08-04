package BF

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

	for {
		// if it can equal you should think what is the length
		if matchlastindex < len(Main) {
			if mainvalue != patternvalue {
				// wrong here
				mainvalue -= int32(Main[matchfirstindex])
				mainvalue += int32(Main[matchlastindex])
				// index add should be done latter
				matchfirstindex++
				matchlastindex++

			} else {
				// if it has the equal val
				if Main[matchfirstindex:matchlastindex] == Pattern {
					return true
				} else {
					matchfirstindex++
					matchlastindex++
				}
			}
		} else {
			return false
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
