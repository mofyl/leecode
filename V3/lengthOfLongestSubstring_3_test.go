package V3

import (
	"fmt"
	"strings"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {

	resStr := ""
	maxLen := 0
	for i := 0; i < len(s); i++ {

		tmp := string(s[i])
		for strings.Contains(resStr, tmp) {
			curLen := len(resStr)
			if curLen > 0 {
				resStr = resStr[1:curLen]
			}
		}
		resStr += tmp
	}
	if len(resStr) > maxLen {
		maxLen = len(resStr)
	}
	return maxLen
}

func lengthOfLongestSubstringDP(s string) int {
	max := 0
	dp := make([]int, len(s))
	if len(s) == 1 {
		return 1
	}
	for i := 1; i < len(s); i++ {
		j := i - 1
		for ; j >= dp[i-1]; j-- {
			if s[j] == s[i] {
				break
			}
		}
		dp[i] = j + 1
	}

	fmt.Println(dp)

	for i, j := range dp {
		if max < i-j+1 {
			max = i - j + 1
		}
	}
	return max
}
