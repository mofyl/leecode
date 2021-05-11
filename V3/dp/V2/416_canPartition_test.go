package V2

import (
	"fmt"
	"testing"
)

func TestCanPartition(t *testing.T) {

	//nums := []int{1, 5, 11, 5}
	//nums := []int{1, 2, 3, 5}
	nums := []int{1, 2, 5}
	//nums := []int{3, 3, 3, 4, 5}
	//nums := []int{2, 2, 3, 5}
	//nums := []int{14, 9, 8, 4, 3, 2}

	res := canPartition(nums)

	fmt.Println(res)
}

func canPartition(nums []int) bool {

	if len(nums) < 2 {
		return false
	}

	// 这里是将求子集和 相同问题，转换为 一个数组的和 是否等于 整个数组和的一半
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	// 这里 当idx== 0 的时候， curSum 要给0，因为curSum可能选择 也可能不选择
	//return canPartitionHelper(nums, 0, 0, target)
	// 这里的DP 改失败了！！！！
	return canPartitionDpWay(nums, target)
}

func canPartitionHelper(nums []int, idx int, curSum int, target int) bool {

	if idx == len(nums) {
		if curSum == target {
			return true
		} else {
			return false
		}
	}

	if curSum == target {
		return true
	}

	// 一个分支是  当前数组不选择， 那就只能另一个数组选择了
	// 另一个分支是 当前数组选择，另外一个数组就不能选择了
	return canPartitionHelper(nums, idx+1, curSum, target) ||
		canPartitionHelper(nums, idx+1, curSum+nums[idx], target)

}

func canPartitionHelperCache(nums []int, idx int, curSum int, target int, dp [][]bool) bool {

	if idx == len(nums) {
		if curSum == target {
			return true
		} else {
			return false
		}
	}

	if curSum == target {
		return true
	}

	// 一个分支是  当前数组不选择， 那就只能另一个数组选择了
	// 另一个分支是 当前数组选择，另外一个数组就不能选择了

	dp[idx+1][curSum] = canPartitionHelperCache(nums, idx+1, curSum, target, dp)
	dp[idx+1][curSum+nums[idx]] = canPartitionHelperCache(nums, idx+1, curSum+nums[idx], target, dp)

	return dp[idx+1][curSum] || dp[idx+1][curSum+nums[idx]]

}

func canPartitionDpWay(nums []int, target int) bool {

	if len(nums) < 2 {
		return false
	}

	// 这里是将求子集和 相同问题，转换为 一个数组的和 是否等于 整个数组和的一半
	N := len(nums)
	dp := make([][]bool, N+1)

	for i := 0; i < N+1; i++ {
		dp[i] = make([]bool, target+1)
	}
	// 这里是最后一行
	//for i := 1; i <= target; i++ {
	//	dp[len(nums)-1][i] = false
	//}

	// 这里是最后一列
	for i := 0; i < N+1; i++ {
		dp[i][target] = true
	}

	// 这里 len(nums) -1 位置已经填完了
	for i := N - 1; i >= 0; i-- {

		if nums[i] == target {
			return true
		}
		for j := 0; j <= target; j++ {
			//tmp := j - nums[i]
			//if tmp < 0 {
			//	// 直接给false
			//	dp[i][j] = dp[i+1][j]
			//} else {
			//	dp[i][tmp] = dp[i+1][j] || dp[i+1][tmp]
			//}

			tmp := j + nums[i]
			if tmp > target {
				// 直接给false
				dp[i][j] = dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j] || dp[i+1][tmp]
			}

		}

	}

	return dp[0][0]
}
