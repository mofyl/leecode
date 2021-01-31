package V1

import (
	"fmt"
	"testing"
)

func TestNumTrees(t *testing.T) {
	n := 3

	res := numTrees(n)

	fmt.Println(res)
	//
	//nums := []int{1, 2, 3}
	//
	//newNums := append(nums[:0], nums[1:]...)
	//
	//fmt.Println(nums)
	//fmt.Println(newNums)

}

func numTrees(n int) int {

	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	// 这个是规定：0个节点就是一颗空树，空树也属于二叉树

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {

		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]

}
