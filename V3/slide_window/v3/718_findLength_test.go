/*
 * @Author: lirui
 * @Date: 2022-01-10 16:23:48
 * @LastEditTime: 2022-01-10 16:39:04
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\slide_window\v3\718_findLength_test.go
 */

package v3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestFindLength(t *testing.T) {
	a := []int{1, 2, 3, 2, 1}
	b := []int{3, 2, 1, 4, 7}

	res := findLength(a, b)

	fmt.Println(res)
}

// 滑动窗口法， 固定一个
func findLength(nums1 []int, nums2 []int) int {

	n := len(nums1)
	m := len(nums2)

	res := -1
	// 固定 n 使用 m 去找

	for i := 0; i < m; i++ {

		minLen := tools.Min(n, m-i)
		cur := findLengthHelper(nums1, nums2, 0, i, minLen)
		res = tools.Max(res, cur)

	}

	for i := 0; i < n; i++ {

		minLen := tools.Min(m, n-i)
		cur := findLengthHelper(nums1, nums2, i, 0, minLen)
		res = tools.Max(res, cur)
	}

	return res

}

func findLengthHelper(a, b []int, aStart, bStart int, step int) int {

	res := 0
	cur := 0
	for i := 0; i < step; i++ {

		if a[aStart+i] == b[bStart+i] {
			cur++
			res = tools.Max(cur, res)
		} else {
			cur = 0
		}

	}

	return res

}
