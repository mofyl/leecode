/*
 * @Author: lirui
 * @Date: 2022-01-10 17:13:51
 * @LastEditTime: 2022-01-10 17:13:51
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\slide_window\v3\75_sortColors_test.go
 */

package v3

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {

	//nums := []int{2, 0, 2, 1, 1, 0}
	//nums := []int{2, 0, 1}
	//nums := []int{0}
	nums := []int{1}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {

	l := -1
	r := len(nums)
	idx := 0
	for idx < r {

		if nums[idx] == 0 {
			l++
			swap(&nums[l], &nums[idx])
		} else if nums[idx] == 1 {

		} else if nums[idx] == 2 {
			r--
			swap(&nums[r], &nums[idx])
			continue
		}
		idx++

	}

}

func swap(a, b *int) {

	tmp := *a
	*a = *b
	*b = tmp

}
