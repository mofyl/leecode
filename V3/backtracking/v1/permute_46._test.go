package v1

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}

	res := permute(nums)

	fmt.Println(res)
}

func permute(nums []int) [][]int {

	res := make([][]int, 0, len(nums))
	used := make([]bool, len(nums))
	permuteLoop(0, nums, []int{}, &res, used)
	//permuteLoopV2(nums, []int{}, &res)
	return res
}

func permuteLoopV2(nums []int, curNums []int, res *[][]int) {

	if len(nums) < 1 {
		tmp := make([]int, len(curNums))
		copy(tmp, curNums)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {

		curNums = append(curNums, nums[i])
		// 这里的思想就是 1 和 [2 , 3] 进行组合排列
		// 这里降低了 递归的深度
		tmp := []int{}
		tmp = append(tmp, nums[:i]...)
		tmp = append(tmp, nums[i+1:]...)
		permuteLoopV2(tmp, curNums, res)
		curNums = curNums[:len(curNums)-1]
	}

}

func permuteLoop(start int, nums []int, curNums []int, res *[][]int, used []bool) {

	if start == len(nums) {
		tmp := make([]int, len(curNums))
		copy(tmp, curNums)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		curNums = append(curNums, nums[i])
		// 这里不能使用 i+1 因为 当我们判断到 3 的时候 其实还可以向下递归，但是这里由于 i 已经超过了 len(nums) 所以被挡掉了
		//permuteLoop(i+1, nums, curNums, res, used)
		permuteLoop(start+1, nums, curNums, res, used)
		curNums = curNums[:len(curNums)-1]
		used[i] = false
	}

}
