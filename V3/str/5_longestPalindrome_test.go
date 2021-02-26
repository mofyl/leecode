package str

import (
	"fmt"
	"strings"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {

	//s := "babad"
	s := "ac"

	res := longestPalindrome(s)

	fmt.Println(res)
}

// 由于是找回文串，那么选择 中心扩展，以一个字符为中心向两边扩展
// 但是 s的长度为奇数还是偶数 会影响中心，所以干脆直接s变为奇数，在开始的时候
// 	预处理s，为s填充 # 号.这样出来的s就都是奇数个
// 	s = "abc"  填充后  s = "#a#b#c#"
func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	// 预处理 s
	sNew := strings.Builder{}

	for i := 0; i < len(s); i++ {
		sNew.WriteString("#")
		sNew.WriteByte(s[i])
	}
	sNew.WriteByte('#')
	s = sNew.String()
	res := ""
	maxLen := 0
	for i := 1; i < len(s); i++ {
		curLen := findCenter(s, i)

		if maxLen < curLen {
			maxLen = len(str)
			res = str
		}
	}

	return res
}

// 左边界和右边界
func findCenter(s string, center int) int {

	left := center + 1
	right := center - 1
	step := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
		step++
	}
	return step
}
