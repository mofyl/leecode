package V3

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	// candidates := []int{2, 3, 5}
	// candidates := []int{1, 2}
	target := 7
	res := combinationSum(candidates, target)
	fmt.Println(res)
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0, len(candidates))
	tmp := make([]int, 0, 3)
	combinationSumRecursion(candidates, 0, target, tmp, &res)

	return res
}

func combinationSumRecursion(candidates []int, cur int, target int, tmp []int, res *[][]int) {

	if cur == target {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*res = append(*res, t)
		return
	}

	if cur > target {
		return
	}

	for i := 0; i < len(candidates); i++ {
		tmp = append(tmp, candidates[i])
		combinationSumRecursion(candidates[i:], cur+candidates[i], target, tmp, res)
		tmp = tmp[:len(tmp)-1]
	}
}
