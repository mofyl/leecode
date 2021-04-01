package v1

import (
	"fmt"
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {

	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	res := combinationSum2(candidates, target)

	fmt.Println(res)

}

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(candidates))
	sort.Ints(candidates)
	combinationSumRecursion2(0, candidates, 0, target, []int{}, &res, used)

	return res
}

func combinationSumRecursion2(start int, candidates []int, cur int, target int, curUsed []int, res *[][]int, used []bool) {

	if cur == target {
		tmp := make([]int, len(curUsed))
		copy(tmp, curUsed)
		*res = append(*res, tmp)
	}

	if cur > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
			continue
		}

		curUsed = append(curUsed, candidates[i])
		used[i] = true
		combinationSumRecursion2(i+1, candidates, cur+candidates[i], target, curUsed, res, used)
		used[i] = false
		curUsed = curUsed[:len(curUsed)-1]
	}

}
