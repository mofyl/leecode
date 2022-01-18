/*
 * @Author: lirui
 * @Date: 2022-01-08 14:26:31
 * @LastEditTime: 2022-01-08 14:42:33
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\backtracking\v3\77_combine_test.go
 */

package v3

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {

	n := 4
	k := 2

	res := combine(n, k)

	fmt.Println(res)
}

func combine(n int, k int) [][]int {

	res := make([][]int, 0)

	combineHelper(1, n, k, []int{}, &res)

	return res
}

func combineHelper(idx, n, k int, cur []int, res *[][]int) {

	// len(cur) > k  表示里面的元素已经大于 k个 那么就没有必要进行本次递归了
	// len(cur) + (n-idx) < k 表示  k 太大了，当前 cur 中的元素 + 剩余的元素  还是不够k个。也没有必要进行递归
	// 这里 +1 是因为 idx 是从1开始的，
	if len(cur) > k || len(cur)+(n-idx+1) < k {
		return

	}

	if len(cur) == k {
		tmp := make([]int, k)
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := idx; i <= n; i++ {

		cur = append(cur, i)

		combineHelper(i+1, n, k, cur, res)

		cur = cur[:len(cur)-1]
	}

}
