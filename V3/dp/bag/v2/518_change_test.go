package v2

import (
	"fmt"
	"testing"
)

func TestChange(t *testing.T) {

	coins := []int{1, 2, 5}
	amount := 5

	//coins := []int{10}
	//amount := 10

	res := change(amount, coins)

	fmt.Println(res)
}

func change(amount int, coins []int) int {

	//return changeHelper(coins, amount, 0)

	//return changeDp(coins, amount)

	return changeDpV2(coins, amount)

}

func changeHelper(coins []int, amount int, idx int) int {

	if amount == 0 {
		return 1
	}

	if idx >= len(coins) {
		return 0
	}

	if amount < 0 {
		return 0
	}

	// 可以选择当前 硬币， 也可以不选择

	// 不选择
	noOp := changeHelper(coins, amount, idx+1)

	// 选择 当前硬币
	//choice := 0

	for i := 1; amount >= i*coins[idx]; i++ {
		noOp += changeHelper(coins, amount-i*coins[idx], idx+1)
	}

	return noOp

}

func changeDp(coins []int, amount int) int {

	dp := make([][]int, len(coins)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}

	// init dp
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}

	for i := len(dp) - 2; i >= 0; i-- {

		for j := 1; j <= amount; j++ {

			choice := dp[i+1][j]
			for k := 1; j >= k*coins[i]; k++ {

				choice += dp[i+1][j-k*coins[i]]

			}

			dp[i][j] = choice

		}

	}

	return dp[0][amount]

}

func changeDpV2(coins []int, amount int) int {

	dp := make([]int, amount+1)

	dp[0] = 1

	for i := 0; i < len(coins); i++ {

		for j := coins[i]; j <= amount; j++ {

			// 我们这里是循环数组，计算 下一次结果时， 上一次结果已经在数组中了，所以这里不用额外的 循环
			//choice := dp[j]
			//for k := 1; j >= k*coins[i]; k++ {
			//	choice += dp[j-k*coins[i]]
			//}
			//
			//dp[j] = choice

			dp[j] += dp[j-coins[i]]

		}

	}

	return dp[amount]

}
