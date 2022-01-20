package v3

import (
	"fmt"
	"leecode/tools"
	"math"
	"testing"
)

func TestCoinChange(t *testing.T) {

	// coins := []int{1, 2, 5}
	// amount := 11

	// coins := []int{2}
	// amount := 3

	coins := []int{1}
	amount := 2

	res := coinChange(coins, amount)

	fmt.Println(res)
}

func coinChange(coins []int, amount int) int {

	return coinChangeHelper(coins, amount)
}

// 返回  构成 amount 需要的硬币数
func coinChangeHelper(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	if amount < 1 {
		return -1
	}
	minRes := math.MaxInt

	for i := 0; i < len(coins); i++ {

		curRes := coinChangeHelper(coins, amount-coins[i])
		if curRes == -1 {
			continue
		}

		minRes = tools.Min(minRes, curRes+1)

	}

	if minRes == math.MaxInt {
		return -1
	}

	return minRes

}

func coinChangeDP(coins []int, amount int) int {

	if amount < 1 {
		return 0
	}
	// 还剩 i 这么多钱的时候 使用的硬币数
	dp := make([]int, amount+1)

	for i := 1; i < len(dp); i++ {

		curMin := math.MaxInt

		for j := 0; j < len(coins); j++ {

			cur := i - coins[j]

			if cur >= 0 {
				num := dp[cur]
				if num == -1 {
					continue
				}

				curMin = tools.Min(curMin, num)
			}
		}

		if curMin == math.MaxInt {
			dp[i] = -1
		} else {
			dp[i] = curMin
		}

	}

	return dp[amount]

}
