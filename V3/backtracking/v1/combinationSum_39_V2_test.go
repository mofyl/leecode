package v1

import (
	"fmt"
	"testing"
)

func TestCombinationSumV2(t *testing.T) {
	//
	//nums := []int{2, 3, 6, 7}
	//target := 7

	nums := []int{2, 3, 5}
	target := 8
	res := combinationSumV2(nums, target)
	fmt.Println(res)
}

func combinationSumV2(candidates []int, target int) [][]int {

	if len(candidates) < 1 {
		return nil
	}

	res := make([][]int, 0, len(candidates))

	combinationSumLoopV2(0, candidates, []int{}, 0, target, &res)
	return res
}

func combinationSumLoopV2(start int, nums []int, curNums []int, cur int, target int, res *[][]int) {

	if cur > target {
		return
	}

	if cur == target {
		tmp := make([]int, len(curNums))
		copy(tmp, curNums)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(nums); i++ {

		curNums = append(curNums, nums[i])
		// 这里给i 是因为可以重复使用
		// 当前元素可以一直使用，直到 cur == target || cur > target 才会换下一个元素
		combinationSumLoopV2(i, nums, curNums, cur+nums[i], target, res)
		curNums = curNums[:len(curNums)-1]
	}

}
