package str

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	//s := "A man, a plan, a canal: Panama"
	s := "race a car"

	res := isPalindrome(s)

	fmt.Println(res)

}

func isPalindrome(s string) bool {

	if s == "" {
		return true
	}

	left := 0
	right := len(s) - 1

	// 这里不用等号的原因是：
	for left < right {

		if !isIllegal(s[left]) {
			left++
			continue
		}

		if !isIllegal(s[right]) {
			right--
			continue
		}

		cL := convert(s[left])
		cR := convert(s[right])

		if cL != cR {
			return false
		}

		left++
		right--

	}

	return true

}

func isIllegal(c uint8) bool {

	if (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') {
		return true
	}

	return false
}

// 这里将大写统一转化为 小写
func convert(c uint8) uint8 {

	if c >= 'A' && c <= 'Z' {
		return c + 32
	}

	return c
}
