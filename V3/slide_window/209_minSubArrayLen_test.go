package slide_window

import (
	"fmt"
	"leecode/tools"
	"math"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	target := 7
	//nums := []int{2, 3, 1, 2, 4, 3}
	nums := []int{7, 1, 2}

	//target := 4
	//nums := []int{1, 4, 4}

	//target := 11
	//nums := []int{1, 1, 1, 1, 1, 1, 1, 1}

	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}

func minSubArrayLen(target int, nums []int) int {

	if len(nums) < 1 {
		return 0
	}

	num := 0
	minLen := math.MaxInt64
	left, right := 0, 0
	for right < len(nums) {

		num += nums[right]
		// 题目求的是大于等于的最小子数组
		for num >= target {
			minLen = tools.Min(minLen, right-left+1)
			num -= nums[left]
			left++
		}

		right++

	}

	if minLen == math.MaxInt64 {
		return 0
	}

	return minLen

}
