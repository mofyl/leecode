package v1

import (
	"fmt"
	"math"
	"testing"
)

func TestMinWindow(t *testing.T) {

	s := "ADOBECODEBANC"

	s1 := "ABC"

	//res := minWindow(s, s1)
	res := minWindow_v2(s, s1)

	fmt.Println(res)
}

func minWindow(s string, t string) string {

	if len(s) < 1 || len(t) < 1 {
		return ""
	}

	table := [128]int{}
	cnt := len(t)
	for i := 0; i < len(t); i++ {
		table[t[i]]++
	}

	left := 0
	right := 0
	minLen := math.MaxInt64
	start := 0
	for right < len(s) {

		if table[s[right]] > 0 {
			cnt--
		}
		table[s[right]]--

		if cnt == 0 {

			// 收缩左边界 直到遇到t中的字符, 因为要求最小值，一定要去掉多余的字符
			for left < right && table[s[left]] < 0 {
				left++
				table[s[left]]++
			}

			// 记录本次的长度
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			// 这是一个优化，我们只需要对 当前遇到的t中字符进行判断就好
			table[s[left]]++
			cnt++
			left++

		}

		right++
	}

	if minLen == math.MaxInt64 {
		return ""
	}
	return s[start : start+minLen]

}

func minWindow_v2(s string, t string) string {

	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// 这里要先有一个128 的字符表， ascii码就128个

	table := [128]int{}
	cnt := len(t)
	// 这里 为目标串添加 数量 表示
	for i := 0; i < len(t); i++ {
		table[t[i]]++
	}

	left := 0
	right := 0
	start := 0
	maxLen := math.MaxInt64

	for right < len(s) {

		if table[s[right]] > 0 {
			cnt--
		}
		table[s[right]]--

		if cnt == 0 {

			// 这时 开始清理没用的字符, 同时收缩左边界
			for left < right && table[s[left]] < 0 {
				table[s[left]]++
				left++
			}
			// 当 循环结束时 left一定指向t 中的某一个字符
			if right-left+1 < maxLen {
				maxLen = right - left + 1
				start = left
			}

			// 此时所有的t 都已经在table中了，那么我们下次遍历的时候
			// 只需要遍历 当前 left 指向的t 中的该字符就好 这里是一个优化
			table[s[left]]++
			cnt++
			left++
		}
		right++
	}
	// 这里表示没有找到
	if maxLen == math.MaxInt64 {
		return ""
	}

	return s[start : start+maxLen]
}
