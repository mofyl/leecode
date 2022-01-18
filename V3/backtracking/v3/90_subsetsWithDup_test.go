/*
 * @Author: lirui
 * @Date: 2022-01-07 17:58:41
 * @LastEditTime: 2022-01-08 18:24:04
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\backtracking\v3\90_subsetsWithDup_test.go
 */

package v3

import (
	"fmt"
	"sort"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {

	nums := []int{1, 2, 2}

	res := subsetsWithDup(nums)

	fmt.Println(res)

}

// 该题要求 同一个树下， A 分支用过了，B分支就不能使用了
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	// 空集合 是所有集合的子集
	res = append(res, []int{})
	// 将相同元素放到一起
	// 要想不重复，相同元素只能用一次
	sort.Ints(nums)

	if len(nums) < 1 {
		return nil
	}

	subsetsWithDupHelper(nums, 0, []int{}, &res)

	return res
}

// visit 中 0 表示 未被访问过
// visit 中 1 表示 已经被访问过
func subsetsWithDupHelperVisit(nums []int, l int, cur []int, res *[][]int, visit []int) {

	if l == len(nums) {
		return
	}

	for i := l; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] && visit[i-1] == 0 {
			continue
		}

		cur = append(cur, nums[i])
		visit[i] = 1
		tmp := make([]int, len(cur))
		copy(tmp, cur)

		*res = append(*res, tmp)

		subsetsWithDupHelper(nums, i+1, cur, res)
		visit[i] = 0
		cur = cur[:len(cur)-1]
	}

}

// 不使用额外空间的情况
func subsetsWithDupHelper(nums []int, l int, cur []int, res *[][]int) {

	if l == len(nums) {
		return
	}

	for i := l; i < len(nums); i++ {

		if i > l && nums[i] == nums[i-1] {
			continue
		}

		cur = append(cur, nums[i])

		tmp := make([]int, len(cur))
		copy(tmp, cur)

		*res = append(*res, tmp)

		subsetsWithDupHelper(nums, i+1, cur, res)

		cur = cur[:len(cur)-1]
	}

}
