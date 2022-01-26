package v2

import (
	"fmt"
	"testing"
)

func TestCanPartition(t *testing.T) {

	nums := []int{1, 5, 11, 5}
	// nums := []int{1, 2, 3, 5}

	res := canPartition(nums)

	fmt.Println(res)
}

func canPartition(nums []int) bool {

	// 这里是将 整个元素分成 和相等的两个子集，
	// 这两个子集必须包括所有的元素。

	// 这里将问题转换为 求一个子集的和 是否等于 所有元素和的一半。

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	// 这里若是奇数 ，则不能分为两个子集
	if sum&1 == 1 {
		return false
	}

	// 求 nums 中的 某个子集 是否 是 sum 的 一半
	// return canPartitionHelper(nums, sum>>1, 0)

	return canPartitionDP(nums, sum>>1)
}

func canPartitionHelper(nums []int, sum, idx int) bool {

	if sum == 0 {
		return true
	}

	if idx >= len(nums) {
		return false
	}

	if sum < 0 {
		return false
	}

	// 可以选择 当前元素， 也可以不选择
	noOp := canPartitionHelper(nums, sum, idx+1)
	if sum >= nums[idx] {
		choice := canPartitionHelper(nums, sum-nums[idx], idx+1)
		return noOp || choice
	}

	return noOp
}

// row 表示 共有  nums 个位置
// col 表示 当前 剩余的 容量
func canPartitionDP(nums []int, sum int) bool {

	dp := make([]bool, sum+1)

	// for i := 0; i < len(dp); i++ {
	// 	dp[i] = make([]bool, sum+1)
	// }

	dp[0] = true

	for i := 0; i < len(nums); i++ {

		for j := sum; j >= nums[i]; j-- {

			dp[j] = dp[j] || dp[j-nums[i]]

		}

	}

	return dp[sum]

}
