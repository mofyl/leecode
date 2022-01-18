package v3

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)

	fmt.Println(res)
}

func maxSubArray(nums []int) int {

	sum := 0
	res := -100000
	for i := 0; i < len(nums); i++ {

		if sum > 0 {
			// 若当前 sum > 0 表示 sum + nums[i] 是有意义的。
			// 若 sum < 0 表示 sum + nums[i] 无意义 因为会导致最后的和 变小。
			sum += nums[i]
		} else {
			// 若sum < 0 则舍弃，然后 已当前元素为新的sum
			// 若当前 num[i] < 0 也没关系，到下一个元素的时候，该nums[i] 就会被舍弃
			sum = nums[i]
		}

		res = tools.Max(res, sum)

	}

	return res
}
