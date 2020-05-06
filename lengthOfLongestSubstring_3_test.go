package main

import (
	"fmt"
	"strings"
)

func main1() {

	str := "abcabcbb"
	// str = "pwwkew"
	// str = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring_v1(s string) int {

	resStr := ""

	maxLen := 0
	curLen := 0

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
		if curLen > maxLen {
			maxLen = curLen
		}

	}

	return maxLen
}

func lengthOfLongestSubstring(s string) int {
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
