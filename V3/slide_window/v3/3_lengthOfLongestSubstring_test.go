/*
 * @Author: lirui
 * @Date: 2022-01-09 14:52:41
 * @LastEditTime: 2022-01-09 15:46:09
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\slide_window\v3\3_lengthOfLongestSubstring_test.go
 */

package v3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	s := "abcabcbb"
	// s := "au"
	// s := "aab"
	// s := "abba"
	// s := "dvdf"

	res := lengthOfLongestSubstring(s)

	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {

	if len(s) < 2 {
		return len(s)
	}

	l := 0
	r := l
	max := -1
	data := make([]int, 256)

	for i := 0; i < len(data); i++ {
		data[i] = -1
	}

	for ; r < len(s); r++ {

		v := data[s[r]]

		if v != -1 {
			// 说明发生了重复
			l = tools.Max(v, l)
		}

		// 每次都来检查 cur 和 max
		// 这里进行+1 是因为 若发生重复 我们前面已经将l 移动了，
		max = tools.Max(r-l+1, max)
		// 记录下一个不重复的位置，这里假设 r+1 位置不会重复
		data[s[r]] = r + 1
	}

	return max

}
