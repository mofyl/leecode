package v2

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {

	//s := "cbbd"

	//s := "babad"

	//s := "bb"

	s := "ac"

	res := longestPalindrome(s)

	fmt.Println(res)
}

func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	left, right := 0, 0
	for i := 0; i < len(s); i++ {

		// 假设该 回文串 是奇数个
		l, r := longestPalindromeHelper(s, i-1, i+1)

		if r-l+1 > right-left+1 {
			left = l
			right = r
		}

		// 假设该 回文串 是偶数个
		l, r = longestPalindromeHelper(s, i, i+1)
		if r-l+1 > right-left+1 {
			left = l
			right = r
		}
	}

	return s[left : right+1]

}

func longestPalindromeHelper(s string, l, r int) (int, int) {

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return l + 1, r - 1

}
