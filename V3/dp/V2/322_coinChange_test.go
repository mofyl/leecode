package V2

import (
	"fmt"
	"leecode/tools"
	"math"
	"testing"
)

func TestCoinChange(t *testing.T) {
	//
	//coins := []int{1, 2, 5}
	//amount := 11

	//coins := []int{2}
	//amount := 3

	coins := []int{1}
	amount := 2

	res := coinChange(coins, amount)
	fmt.Println(res)
}

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	//return coinChangeHelper(coins, 0, amount)
	//return coinChangeHelperV2(coins, amount)
	//return coinChangeDpWay(coins, amount)
	return coinChangeDpV2(coins, amount)

}

// 由于硬币数量没有规定，那么每次递归就可以从头遍历硬币数组
// 递归含义： 组成amount 金额 所需要的最少硬币数量
func coinChangeHelperV2(coins []int, amount int) int {

	// 这里表示当前金额已经到了0以下 当前方案不正确
	if amount < 0 {
		return -1
	}

	// 这里表示刚好组成 目标金额 后面就不需要再进行了
	// 由于这里没有使用硬币，所以是0
	if amount == 0 {
		return 0
	}

	// 当前硬币可以不选择，也可以选择 选择就要增加硬币数量

	// 这里表示 当前选择 i枚 硬币 若 硬币面值*数量  > amount 那么就不对了
	min := math.MaxInt64

	for i := 0; i < len(coins); i++ {
		nextOp := coinChangeHelperV2(coins, amount-coins[i])
		if nextOp != -1 {
			min = tools.Min(nextOp+1, min)
		}
	}

	if min == math.MaxInt64 {
		return -1
	}

	return min

}

func coinChangeDpV2(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	// 还剩 n 这么多钱时，需要的最少硬币数
	dp := make([]int, amount+1)

	// 由于amount = 0 时 需要的个数为0 个
	for i := 1; i < len(dp); i++ {

		min := math.MaxInt64
		for j := 0; j < len(coins); j++ {
			curAmount := i - coins[j]
			if curAmount >= 0 {
				nextOp := dp[curAmount]
				if nextOp != -1 {
					min = tools.Min(min, dp[curAmount]+1)
				}
			}

		}

		if min == math.MaxInt64 {
			dp[i] = -1
		} else {
			dp[i] = min
		}

	}

	return dp[amount]

}

// 递归含义： 使用当前 coins[idx....] 之后的硬币， 组成 amount 金额 使用的最少硬币数量
func coinChangeHelper(coins []int, idx int, amount int) int {

	// 这里表示当前金额已经到了0以下 当前方案不正确
	if amount < 0 {
		return -1
	}

	// 这里表示刚好组成 目标金额 后面就不需要再进行了
	// 由于这里没有使用硬币，所以是0
	if amount == 0 {
		return 0
	}
	// 这里表示 硬币已经使用到了末尾
	if idx == len(coins) {
		// 若正好组成 amount 这么多的金额 那么这里也没有使用硬币 就返回0
		if amount == 0 {
			return 0
		} else {
			// 若没有组成 表示当前方案不对
			return -1
		}
	}

	// 当前硬币可以不选择，也可以选择 选择就要增加硬币数量

	// 这里是不选择的情况， 直接到下一个位置 金额数量不会减少
	noOp := coinChangeHelper(coins, idx+1, amount)

	// 这里表示 当前选择 i枚 硬币 若 硬币面值*数量  > amount 那么就不对了
	min := math.MaxInt64
	for i := 1; coins[idx]*i <= amount; i++ {
		nextOp := coinChangeHelper(coins, idx+1, amount-coins[idx]*i)
		if nextOp != -1 {
			min = tools.Min(nextOp+i, min)
		}
	}

	if min == math.MaxInt64 {
		return noOp
	}

	if noOp == -1 {
		return min
	}

	return tools.Min(min, noOp)

}

func coinChangeDpWay(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	//创建一个二维的 dp表, 行表示当前使用的 硬币
	// 列表示 剩余的金额

	// 这里是可以取到 len(coins) 的所以给的 coins+1
	dp := make([][]int, len(coins)+1)

	for i := 0; i < len(dp); i++ {
		// 这里是可以取到 amount的， 所以给amount
		dp[i] = make([]int, amount+1)
	}

	// 填充最后一行
	for i := 1; i < amount+1; i++ {
		dp[len(dp)-1][i] = -1
	}

	for i := len(dp) - 2; i >= 0; i-- {

		// 首先是不选当前硬币的状态

		// 这里amount == 0 都是0 所以就直接跳过 第0列
		// 递归那里 是从amount开始的
		for j := 1; j <= amount; j++ {

			noOp := dp[i+1][j]
			min := math.MaxInt64

			for k := 1; coins[i]*k <= j; k++ {
				useOp := dp[i+1][j-coins[i]*k]
				if useOp != -1 {
					min = tools.Min(useOp+k, min)
				}
			}

			if min == math.MaxInt64 {
				dp[i][j] = noOp
			} else if noOp == -1 {
				dp[i][j] = min
			} else {
				dp[i][j] = tools.Min(min, noOp)
			}

		}

	}

	return dp[0][amount]

}
