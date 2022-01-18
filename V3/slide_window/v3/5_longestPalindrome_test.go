/*
 * @Author: lirui
 * @Date: 2022-01-09 17:04:55
 * @LastEditTime: 2022-01-10 14:04:43
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\slide_window\v3\5_longestPalindrome_test.go
 */

package v3

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s := "ccc"

	res := longestPalindrome(s)

	fmt.Println(res)
}

func longestPalindrome(s string) string {

	if len(s) < 2 {
		return s
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {

		l1, r1 := expendAroundCenter(s, i-1, i)

		l2, r2 := expendAroundCenter(s, i-1, i+1)

		if end-start < r1-l1 {
			start = l1
			end = r1
		}

		if end-start < r2-l2 {
			start = l2
			end = r2
		}

	}

	return s[start : end+1]

}

func expendAroundCenter(s string, l, r int) (int, int) {

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	// 这里是因为 上面边界条件为  >= 0 和  r < len(s)
	// 若 s[l] != s[r]  那么 l+1 和  r-1 就指向了  s[l] == s[r] 的位置
	// 若 s[l] == s[r] 但是 l 或 r 越界了，那么l+1 和 r-1 指向的就是没有越界的位置
	return l + 1, r - 1

}
