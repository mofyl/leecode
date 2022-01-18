/*
 * @Author: lirui
 * @Date: 2022-01-08 16:24:52
 * @LastEditTime: 2022-01-08 19:14:35
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\backtracking\v3\46_permute_test.go
 */

package v3

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {

	nums := []int{5, 4, 6, 2}

	res := permute(nums)

	fmt.Println(res)
}

func permute(nums []int) [][]int {

	res := make([][]int, 0)
	used := make([]bool, len(nums))
	permuteHelper(nums, 0, used, []int{}, &res)

	return res
}

func permuteHelper(nums []int, idx int, used []bool, cur []int, res *[][]int) {

	if idx == len(nums) {

		tmp := make([]int, len(cur))
		copy(tmp, cur)

		*res = append(*res, tmp)

		return
	}

	for i := 0; i < len(nums); i++ {
		// 若该元素已经被使用过，那么就continue
		if used[i] {
			continue
		}

		cur = append(cur, nums[i])
		// 每次去标记当前子树使用过的元素
		used[i] = true
		permuteHelper(nums, idx+1, used, cur, res)
		used[i] = false
		cur = cur[:len(cur)-1]

	}

}
