package V3

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)

	fmt.Println(res)
}

func maxSubArray(nums []int) int {
	sum := 0
	ans := nums[0]

	for i := 0; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
