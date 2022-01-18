/*
 * @Author: lirui
 * @Date: 2022-01-08 15:32:46
 * @LastEditTime: 2022-01-08 15:58:32
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\backtracking\v3\216_combinationSum3_test.go
 */

package v3

import (
	"fmt"
	"testing"
)

func TestCombinationSum3(t *testing.T) {

	k := 3
	n := 9

	res := combinationSum3(k, n)

	fmt.Println(res)
}

func combinationSum3(k int, n int) [][]int {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	res := make([][]int, 0)
	combinationSum3Helper(0, k, n, nums, []int{}, 0, &res)

	return res
}

func combinationSum3Helper(idx, k, n int, nums []int, cur []int, curSum int, res *[][]int) {

	if curSum > n ||
		idx > len(nums) ||
		len(cur) > k {
		return
	}

	if len(cur) == k && curSum == n {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := idx; i < len(nums); i++ {

		if nums[i] > n {
			break
		}

		cur = append(cur, nums[i])

		combinationSum3Helper(i+1, k, n, nums, cur, curSum+nums[i], res)

		cur = cur[:len(cur)-1]

	}

}
