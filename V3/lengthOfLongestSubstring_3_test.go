package V3

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
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
