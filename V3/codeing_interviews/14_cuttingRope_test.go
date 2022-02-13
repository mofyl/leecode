package codeing_interviews

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestCuttingRope(t *testing.T) {

	//n := 2
	//n := 10

	//n := 3

	n := 120

	res := cuttingRope(n)

	fmt.Println(res)
}

func cuttingRope(n int) int {

	//// 最多可以被分为n 段
	//max := 0
	//for i := 2; i <= n; i++ {
	//
	//	cur := cuttingRopeHelper(i, n)
	//	max = tools.Max(max, cur)
	//}
	//
	//return max

	// 该题 可以不用管 分几段的，反正就是 n 这么长的 绳子。
	//return cuttingRopeHelperV2(n)

	return cuttingRopeDP(n)
}

// n 长度的绳子，切割 m段，m段可以获得的最大乘积
func cuttingRopeHelperV2(n int) int {

	//  长度为2时 的 最大乘积就是1  因为只能 1*1
	if n <= 2 {
		return 1
	}

	// 这里 不管能分多少段，直到绳子不能分为止
	// 这里的循环表示 当前使用的绳子的长度
	max := 0
	// 这一段不能都把 n用完,若这里取到了 n 那么下一段可以分的 长度就是0 ，长度不能为0
	for i := 1; i < n; i++ {

		/*
			这里我们有以下选择
			剩下的 n-i 不进行拆分 : i*(n-i)
			剩下的 n-i 进行拆分 : i*cuttingRopeHelperV2(n-i)
		*/
		cur := i * cuttingRopeHelperV2(n-i)
		max = tools.Max(max, tools.Max(i*(n-i), cur))
	}

	return max

}

// 定义为 切割为n段时  可以获得的最大乘积
// remainderLen 切割为n段时 剩余的长度
func cuttingRopeHelper(n int, remainderLen int) int {

	if remainderLen < 1 {
		return 0
	}
	// 切割为 1段时 最大乘积为 remainderLen
	if n == 1 {
		return remainderLen
	}
	max := 0
	// 当前可以用的 长度
	for i := 1; i <= remainderLen; i++ {

		cur := i * cuttingRopeHelper(n-1, remainderLen-i)

		max = tools.Max(max, cur)
	}

	return max

}

// 依据 cuttingRopeHelperV2 来改的
func cuttingRopeDP(n int) int {

	// 长度为 n 时的 最大乘积
	dp := make([]int, n+1)

	// 还剩 1 长度的绳子时 最大乘积 为 1
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < len(dp); i++ {

		for j := i; j > 0; j-- {

			dp[i] = tools.Max(dp[i], tools.Max(j*(i-j), j*dp[i-j]))

		}

	}

	return dp[n]

}
