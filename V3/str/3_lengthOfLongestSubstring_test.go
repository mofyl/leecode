package str

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	//s := "abcabcbb"
	//s := "bbbb"
	s := "pwwkew"

	res := lengthOfLongestSubstring(s)
	fmt.Println(res)

}

func lengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	strMap := make(map[byte]int, len(s))

	start := 0
	maxLen := 0
	for end := 0; end < len(s); end++ {

		noRepeat, ok := strMap[s[end]]

		if ok {
			start = tools.Max(start, noRepeat)
		}
		maxLen = tools.Max(maxLen, end-start+1)
		strMap[s[end]] = end + 1
	}

	return maxLen
}
