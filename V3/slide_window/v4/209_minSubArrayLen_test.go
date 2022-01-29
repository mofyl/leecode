package v4

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {

	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}

	//target := 4
	//nums := []int{1, 4, 4}

	res := minSubArrayLen(target, nums)

	fmt.Println(res)

}

func minSubArrayLen(target int, nums []int) int {

	l, r := 0, 0

	curSum := 0
	minLen := len(nums) + 1
	for ; r < len(nums); r++ {

		curSum += nums[r]

		for curSum >= target {

			minLen = tools.Min(minLen, r-l+1)
			curSum -= nums[l]
			l++
		}

	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen

}
