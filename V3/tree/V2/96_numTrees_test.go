package V2

import (
	"fmt"
	"testing"
)

func TestNumTrees(t *testing.T) {

	fmt.Println(numTreeV2(3))
}

func numTrees(n int) int {
	if n < 2 {
		return 1
	}

	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {

		for j := 1; j <= i; j++ {
			dp[i] += dp[i-j] * dp[j-1]
		}

	}

	return dp[n]
}

// 卡塔兰数
// 上面我们得到的 G(n) 公式就是 卡塔兰数
// 我们可以利用 卡塔兰公式 直接得出结果
func numTreeV2(n int) int {

	if n < 2 {
		return 1
	}

	res := 1

	for i := 0; i < n; i++ {
		res = res * 2 * (2*i + 1) / (i + 2)
	}

	return res

}
