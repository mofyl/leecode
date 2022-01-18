/*
 * @Author: lirui
 * @Date: 2022-01-08 14:55:05
 * @LastEditTime: 2022-01-08 17:55:09
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\backtracking\v3\40_combinationSum2_test.go
 */

package v3

import (
	"fmt"
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {

	candi := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	// candi := []int{2, 5, 2, 1, 2}
	// target := 5

	res := combinationSum2(candi, target)

	fmt.Println(res)
}

func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)

	res := make([][]int, 0)
	combinationSum2Helper(0, candidates, target, []int{}, 0, &res)

	return res
}

func combinationSum2Helper(idx int, candidates []int, target int, cur []int, curSum int, res *[][]int) {

	if idx > len(candidates) || curSum > target {
		return
	}

	if curSum == target {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := idx; i < len(candidates); i++ {

		if target < candidates[i] {
			break
		}

		if i > idx && candidates[i] == candidates[i-1] {
			continue
		}

		cur = append(cur, candidates[i])

		combinationSum2Helper(i+1, candidates, target, cur, curSum+candidates[i], res)

		cur = cur[:len(cur)-1]
	}

}
