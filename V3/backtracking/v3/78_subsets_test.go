/*
 * @Author: lirui
 * @Date: 2022-01-07 17:43:36
 * @LastEditTime: 2022-01-07 17:56:01
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\backtracking\v3\78_subsets_test.go
 */

package v3

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {

	nums := []int{1, 2, 3}

	res := subsets(nums)

	fmt.Println(res)
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	// 空集合 是所有集合的子集
	res = append(res, []int{})

	if len(nums) < 1 {
		return nil
	}

	subsetsHelper(nums, 0, []int{}, &res)

	return res
}

func subsetsHelper(nums []int, l int, cur []int, res *[][]int) {

	if l == len(nums) {
		return
	}

	for i := l; i < len(nums); i++ {

		cur = append(cur, nums[i])

		tmp := make([]int, len(cur))
		copy(tmp, cur)

		*res = append(*res, tmp)

		subsetsHelper(nums, i+1, cur, res)

		cur = cur[0 : len(cur)-1]
	}

}
