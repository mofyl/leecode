package codeing_interviews

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

	cur := 0
	max := -1000

	for i := 0; i < len(nums); i++ {

		if cur > 0 {
			cur += nums[i]
		} else {
			cur = nums[i]
		}

		max = tools.Max(max, cur)

	}

	return max
}
