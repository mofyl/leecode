package codeing_interviews

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestCuttingRope_2(t *testing.T) {

	n := 120

	res := cuttingRope_v2(n)

	fmt.Println(res)
}

// 这里不对，本体和 cuttingRope 主体一样， 但是本题多了一个大数的 越界处理。
func cuttingRope_v2(n int) int {

	//return cuttingRopeHelper_2(n)

	return cuttingRope_v2DP(n)

}

func cuttingRopeHelper_2(n int) int {

	if n <= 2 {
		return 1
	}

	max := 0
	for i := 1; i < n; i++ {

		cur := (i * cuttingRopeHelper_2(n-i)) % 1000000007

		max = tools.Max(max, tools.Max(cur, i*(n-i))) % 1000000007

	}

	return max
}

func cuttingRope_v2DP(n int) int {

	dp := make([]int, n+1)

	dp[1] = 1
	dp[2] = 1

	for i := 3; i < len(dp); i++ {

		for j := 1; j < i; j++ {

			//cur := j * (i - j) % 1000000007
			cur := j * (i - j)
			cur1 := j * dp[i-j]
			dp[i] = tools.Max(dp[i], tools.Max(cur, cur1)%1000000007)

		}

	}

	return dp[n]

}
