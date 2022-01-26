package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

const (
	INF = 100000
)

func TestNumSquares(t *testing.T) {

	n := 12

	res := numSquares(n)

	fmt.Println(res)
}

func numSquares(n int) int {

	nums := make([]int, 0)

	for i := 1; i <= n; i++ {

		squ := i * i

		if squ <= n {
			nums = append(nums, squ)
		} else {
			break
		}
	}

	//return numSquaresHelper(nums, n, 0)

	//return numSquaresDP(nums, n)
	//return numSquaresDP(nums, n)
	return numSquaresDPV2(nums, n)
}

func numSquaresHelper(nums []int, n int, idx int) int {

	if n == 0 {
		return 0
	}

	if n < 0 {
		return INF
	}

	if idx >= len(nums) {
		return INF
	}

	noOp := numSquaresHelper(nums, n, idx+1)
	choice := INF

	for i := 1; n >= i*nums[idx]; i++ {
		choice = tools.Min(choice, numSquaresHelper(nums, n-i*nums[idx], idx+1)+i)
	}

	return tools.Min(noOp, choice)

}

func numSquaresDP(nums []int, n int) int {

	dp := make([][]int, len(nums)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	// 这里一定要注意  n == 0 的时候 是0
	for i := 1; i < len(dp[0]); i++ {
		dp[len(dp)-1][i] = INF
	}

	for i := len(dp) - 2; i >= 0; i-- {
		x := nums[i]
		for j := 0; j <= n; j++ {

			dp[i][j] = dp[i+1][j]

			for k := 1; j >= k*x; k++ {

				dp[i][j] = tools.Min(dp[i][j], dp[i+1][j-k*x]+k)

			}

		}

	}
	return dp[0][n]
}

func numSquaresDPV2(nums []int, n int) int {

	dp := make([]int, n+1)

	// 这里一定要注意  n == 0 的时候 是0
	for i := 1; i < len(dp); i++ {
		dp[i] = INF
	}

	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i]
		for j := 0; j <= n; j++ {

			for k := 1; j >= k*x; k++ {

				dp[j] = tools.Min(dp[j], dp[j-k*x]+k)

			}

		}

	}
	return dp[n]
}
