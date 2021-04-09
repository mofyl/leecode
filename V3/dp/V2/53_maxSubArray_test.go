package V2

import (
	"fmt"
	"leecode/tools"
	"math"
	"testing"
)

func TestMaxSubArray(t *testing.T) {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	//res := maxSubArray(nums)
	res := maxSubArrayV2(nums)

	fmt.Println(res)
}

func maxSubArray(nums []int) int {

	if len(nums) < 0 {
		return 0
	}

	res := math.MinInt64
	curSum := 0
	for i := 0; i < len(nums); i++ {

		if curSum > 0 {
			curSum += nums[i]
		} else {
			// 若 curSum 已经小于0 那么就表示 curSum对 nums[i] 没有增益效果
			// 那么就应该 舍弃 curSum， 将 curSum = nums[i]
			curSum = nums[i]
		}
		// 这里一直去取最大值
		res = tools.Max(res, curSum)

	}

	if res == math.MinInt64 {
		return 0
	}

	return res
}

// 递归版
func maxSubArrayV2(nums []int) int {

	if len(nums) < 0 {
		return 0
	}

	return maxSubArrayHelper(nums, 0, 0)
}

// 表示 获取 nums[idx ... ]  连续子数组的 最大和
func maxSubArrayHelper(nums []int, idx int, value int) int {

	if idx == len(nums) {
		return 0
	}

	// 选择策略为：可以选择从 当前元素开始 ，也可以选择 从value继续走

	res1 := maxSubArrayHelper(nums, idx+1, nums[idx])
	res2 := maxSubArrayHelper(nums, idx+1, value+nums[idx])

	return tools.Max(res1, res2)
}
