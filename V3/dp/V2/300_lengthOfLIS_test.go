package V2

import (
	"fmt"
	"leecode/tools"
	"math"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//nums := []int{0, 1, 0, 3, 2, 3}
	//nums := []int{7, 7, 7, 7, 7, 7, 7}
	//res := lengthOfLIS(nums)
	res := lengthOfLISDpWay(nums)

	fmt.Println(res)
}

func lengthOfLIS(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		// 这里+1 是因为 子序列中 nums[i] 这个数字也要包括的
		max = tools.Max(lengthOfLISHelperV2(nums, i)+1, max)
	}
	return max
	//return lengthOfLISHelper(nums, 0, math.MinInt64)

}

func lengthOfLISDpWay(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	dp := make([]int, len(nums)+1)

	max := math.MinInt64
	// len(dp)-1 位置已经填了  所以就是-2
	for i := len(dp) - 2; i >= 0; i-- {

		for j := i + 1; j < len(dp)-1; j++ {

			if nums[j] > nums[i] {
				dp[i] = tools.Max(dp[i], dp[j]+1)
			}

		}

		// 这里+1 是因为 子序列中 nums[i] 这个数字也要包括的
		max = tools.Max(dp[i]+1, max)

	}
	return max

}

// 递归函数定义: 找到 从 curPos ... len(nums) 长度中，
// 满足 比 nums[curPos] 严格递增的 子序列的长度
func lengthOfLISHelperV2(nums []int, curPos int) int {

	if curPos == len(nums) {
		return 0
	}

	max := math.MinInt64
	for i := curPos; i < len(nums); i++ {

		if nums[i] > nums[curPos] {
			res := lengthOfLISHelperV2(nums, i) + 1
			max = tools.Max(res, max)
		}

	}

	if max == math.MinInt64 {
		return 0
	}

	return max

}
