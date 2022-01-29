package v4

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMaxSubArray(t *testing.T) {

	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//nums := []int{5, 4, -1, 7, 8}

	nums := []int{-1}

	res := maxSubArray(nums)

	fmt.Println(res)
}

func maxSubArray(nums []int) int {

	maxSum := -1000000
	curSum := 0
	for r := 0; r < len(nums); r++ {

		if curSum > 0 {
			curSum += nums[r]
		} else {
			curSum = nums[r]
		}

		maxSum = tools.Max(maxSum, curSum)

	}

	return maxSum

}
