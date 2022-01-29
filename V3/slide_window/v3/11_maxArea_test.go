/*
 * @Author: lirui
 * @Date: 2022-01-10 14:08:09
 * @LastEditTime: 2022-01-10 14:18:41
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\slide_window\v3\11_maxArea_test.go
 */

package v3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxArea(t *testing.T) {

	num := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	s := maxArea(num)

	fmt.Println(s)
}

func maxArea(height []int) int {

	if len(height) < 1 {
		return 0
	}

	l, r := 0, len(height)-1

	max := 0
	// 这里没有等于号 因为若 l == r，那么中间就没有水柱了，就不用循环了
	for l < r {

		h := tools.Min(height[l], height[r])
		w := r - l // 由题目中可知 这里 只能取 [l , r) 或者 (l , r]

		max = tools.Max(max, h*w)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return max

}
