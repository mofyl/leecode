/*
 * @Author: lirui38
 * @Date: 2021-12-16 17:27:14
 * @LastEditTime: 2021-12-17 10:31:56
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/96_numTrees_test.go
 */

package V3

import "testing"

func TestNumTrees(t *testing.T) {

	n := 3
	res := numTrees(n)

	t.Log(res)
}

func numTrees(n int) int {

	// dp[i] 表示 当节点数为i时可以有几种 不同的二叉搜索树
	dp := make([]int, n+1)

	dp[0] = 1 // 没有节点就是一颗空树，空树也是二叉搜索树
	dp[1] = 1

	// i 从2开始 是因为 dp[1] 的结果已经确定了，就不用从1 开始了
	//  表示有i个节点时 可以有 几种不同的二叉搜索树
	for i := 2; i <= n; i++ {
		// 表示有 i个节点时， 以j为根节点可以有 几种不同的二叉搜索树
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
