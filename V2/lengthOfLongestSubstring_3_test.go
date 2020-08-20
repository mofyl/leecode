package V2

import "strings"

func lengthOfLongestSubstring(s string) int {

	resStr := ""

	curLen := 0
	maxLen := 0

	for i := 0; i < len(s); i++ {

		str := string(s[i])
		for strings.Contains(resStr, str) {

			if curLen > 0 {
				resStr = resStr[1:curLen]
				curLen--

			}
		}

		resStr += str
		curLen++

		if maxLen < curLen {
			maxLen = curLen
		}

	}

	return maxLen

}
