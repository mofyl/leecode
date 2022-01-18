/*
 * @Author: lirui
 * @Date: 2022-01-10 14:35:58
 * @LastEditTime: 2022-01-10 16:21:26
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\slide_window\v3\209_minSubArrayLen_test.go
 */

package v3

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {

	// target := 7
	// nums := []int{2, 3, 1, 2, 4, 3}

	// target := 4
	// nums := []int{1, 4, 4}

	target := 11
	nums := []int{1, 2, 3, 4, 5}

	n := minSubArrayLen(target, nums)

	fmt.Println(n)
}

func minSubArrayLen(target int, nums []int) int {

	l := 0
	r := 0
	cur := 0
	minLen := len(nums) + 1
	for ; r < len(nums); r++ {

		cur += nums[r]

		for cur >= target {
			minLen = Min(minLen, r-l+1)
			cur -= nums[l]
			l++
		}

	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
