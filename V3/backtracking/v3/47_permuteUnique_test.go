/*
 * @Author: lirui
 * @Date: 2022-01-08 17:14:26
 * @LastEditTime: 2022-01-09 11:56:44
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\backtracking\v3\47_permuteUnique_test.go
 */

package v3

import (
	"fmt"
	"sort"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 1, 2}

	res := permuteUnique(nums)

	fmt.Println(res)
}

func permuteUnique(nums []int) [][]int {

	res := make([][]int, 0, len(nums))
	used := make([]bool, len(nums))
	sort.Ints(nums)
	permuteUniqueLoop(0, nums, []int{}, &res, used)
	return res
}

func permuteUniqueLoop(start int, nums []int, curNums []int, res *[][]int, used []bool) {

	if start == len(nums) {
		tmp := make([]int, len(curNums))
		copy(tmp, curNums)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {

		// i 没有被使用过
		if used[i] {
			continue
		}
		// 跳过 i-1 直接使用的i
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		used[i] = true
		curNums = append(curNums, nums[i])
		permuteUniqueLoop(start+1, nums, curNums, res, used)
		curNums = curNums[:len(curNums)-1]
		used[i] = false
	}

}
