package v1

import (
	"leecode/tools"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {

}

// 滑动窗口问题
func lengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	strMap := make(map[uint8]int, len(s))

	l := 0
	maxLen := 0
	for r := 0; r < len(s); r++ {

		v, ok := strMap[s[r]]

		if ok {
			l = tools.Max(l, v)
		}

		maxLen = tools.Max(maxLen, r-l+1)
		// 这里假设不重复的字符的位置 是下一个字符
		strMap[s[r]] = r + 1
	}

	return maxLen

}
