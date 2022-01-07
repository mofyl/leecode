/*
 * @Author: lirui
 * @Date: 2021-12-28 11:11:44
 * @LastEditTime: 2021-12-30 11:47:31
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\array\V3\1_towsum_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum(nums, target)

	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {

	tmp := make(map[int]int, len(nums))
	res := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		// c - 数字
		n1 := target - nums[i]

		idx, ok := tmp[n1]
		if ok {
			res = append(res, i, idx)
			return res
		} else {
			// 每次记录 a 和 a 的idx
			tmp[nums[i]] = i
		}
	}

	return res

}
