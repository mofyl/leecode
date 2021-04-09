package V2

import (
	"fmt"
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {

	nums := []int{1, 1, 1, 1, 1}
	S := 3

	res := findTargetSumWays(nums, S)

	fmt.Println(res)

}

func findTargetSumWays(nums []int, S int) int {

	if len(nums) < 1 {
		return 0
	}

	//return findTargetSumWaysHelper(nums, S, 0, 0)
	return findTargetSumWaysDp(nums, S)

}

// 返回 对于当前 nums[idx: ....] 选择符号后 累加成为  curS 的方法数
func findTargetSumWaysHelper(nums []int, S int, idx int, curS int) int {

	// 这里表示 idx 到了 nums末尾
	if idx == len(nums) {
		if curS == S {
			// 这里表示 curS 与 S相同了，那么表示是一种成功的方法
			return 1
		} else {
			// 这里表示 curS 没有与S相同，那么表示这不是一种成功的方法
			return 0
		}

	}
	// 不能有这个 因为要求一定要将 nums中的数字都用完
	//if curS == 0 {
	//	return 1
	//}
	// 不能有这个，因为 可能中间有负数的情况
	//if curS < 0 {
	//	return 0
	//}

	// 对于每个数字来说 都可以选择 正数或选择负数 , 最后累加这两种选择的 结果

	// 选择正数
	plusNum := choicePlusNum(nums[idx])
	// 得到正数的结果
	res1 := findTargetSumWaysHelper(nums, S, idx+1, curS+plusNum)
	// 选择负数
	negNum := choiceNegNum(nums[idx])
	// 得到负数的结果
	res2 := findTargetSumWaysHelper(nums, S, idx+1, curS+negNum)

	return res1 + res2

}

func findTargetSumWaysDp(nums []int, S int) int {

	if len(nums) < 1 {
		return 0
	}
	// 由于是可以到 len(nums) 的 所以需要 len(nums)+1
	// dp 含义：行表示当前使用的 nums[i], 列表示 当前已经累加到的值
	dp := make([][]int, len(nums)+1)

	for i := 0; i < len(dp); i++ {
		// 这里也是需要到 S ，所以是S+1
		dp[i] = make([]int, S+1)
	}

	dp[len(nums)][S] = 1

	for i := len(nums) - 2; i >= 0; i-- {

		for j := 0; j <= S; j++ {
			//
			//// 选择正数
			plusNum := choicePlusNum(nums[i])
			negNum := choiceNegNum(nums[i])

			dp[i][j] = dp[i+1][plusNum+j] + dp[i+1][negNum+j]
			//// 得到正数的结果
			//res1 := findTargetSumWaysHelper(nums, S, idx+1, curS+plusNum)
			//// 选择负数
			//negNum := choiceNegNum(nums[idx])
			//// 得到负数的结果
			//res2 := findTargetSumWaysHelper(nums, S, idx+1, curS+negNum)

		}

	}

	return dp[0][0]

}

func choicePlusNum(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func choiceNegNum(num int) int {
	if num <= 0 {
		return num
	}
	return -num
}
