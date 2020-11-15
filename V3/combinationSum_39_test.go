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
	combinationSumRecursion(0, candidates, 0, target, []int{}, &res)

	return res
}

func combinationSumRecursion(start int, candidates []int, cur int, target int, curUsed []int, res *[][]int) {

	if cur == target {
		t := make([]int, len(curUsed))
		copy(t, curUsed)
		*res = append(*res, t)
		return
	}

	if cur > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		curUsed = append(curUsed, candidates[i])
		// 这里是i 因为可以重复使用 用几次都可以
		combinationSumRecursion(i, candidates, cur+candidates[i], target, curUsed, res)
		curUsed = curUsed[:len(curUsed)-1]
	}
}
