/*
 * @Author: lirui
 * @Date: 2021-12-28 14:14:01
 * @LastEditTime: 2021-12-29 10:14:05
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\array\V3\15_threeSum_test.go
 */

package V3

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {

	nums := []int{-1, 0, 1, 2, -1, -4}

	res := threeSum(nums)

	for _, v := range res {
		fmt.Println(v)
	}
}

func threeSum(nums []int) [][]int {

	// 不够3个数
	if len(nums) < 3 {
		return nil
	}

	// 先排序
	sort.Ints(nums)

	res := make([][]int, 0)

	// 固定一个数字， 然后通过窗口来确认 另外两个数字
	for i := 0; i < len(nums); i++ {

		// 若排序后 最小的数字都比0 大 那么和一定比0 大
		if nums[i] > 0 {
			return res
		}

		// 若这些数字相同 则快速跳过,因为3个相等的数字 只能用一个
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 这里使用 i+1 是为了防止结果重复，
		// 如果 0~i 这里的数字可以用,那么在前面的结果中早就已经有了
		l := i + 1
		r := len(nums) - 1

		for l < r {

			n := nums[i] + nums[l] + nums[r]

			if n == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 防止重复
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if n > 0 {
				// 这里表示 nums[r] 太大了
				r--
			} else if n < 0 {
				// 这里表示 nums[l] 太小了
				l++
			}

		}

	}

	return res

}
