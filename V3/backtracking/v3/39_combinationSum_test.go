/*
 * @Author: lirui
 * @Date: 2022-01-08 14:47:13
 * @LastEditTime: 2022-01-08 14:53:55
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\backtracking\v3\39_combinationSum_test.go
 */

package v3

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {

	// candi := []int{2, 3, 6, 7}
	// target := 7

	// candi := []int{2, 3, 5}
	// target := 8

	candi := []int{1}
	target := 2

	res := combinationSum(candi, target)

	fmt.Println(res)
}

func combinationSum(candidates []int, target int) [][]int {

	res := make([][]int, 0)

	combinationSumHelper(0, candidates, target, 0, []int{}, &res)

	return res
}

func combinationSumHelper(idx int, candidates []int, target int, curSum int, cur []int, res *[][]int) {

	if curSum > target {
		return
	}

	if curSum == target {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := idx; i < len(candidates); i++ {

		cur = append(cur, candidates[i])
		combinationSumHelper(i, candidates, target, curSum+candidates[i], cur, res)
		cur = cur[:len(cur)-1]
	}

}
