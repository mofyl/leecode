package V1

import (
	"fmt"
	"sort"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	// nums := make([]int, 0, 3)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 2)

	nums := make([]int, 0, 4)
	nums = append(nums, 3)
	nums = append(nums, 3)
	nums = append(nums, 0)
	nums = append(nums, 3)

	// sort.Ints(nums)

	fmt.Println(permuteUnique(nums))
}

func backtracePermuteUnique(idx int, used []bool, res *[][]int, builder *[]int, nums []int) {

	if idx == len(nums) {
		newBuilder := make([]int, len(*builder))
		copy(newBuilder, *builder)
		*res = append(*res, newBuilder)
		(*builder) = (*builder)[:]
		return
	}

	for i := 0; i < len(nums); i++ {

		if used[i] {
			continue
		}

		// 若当前和前一个相同，并且前一个已经使用过了 那就这支就该被剪掉
		if i > 0 && nums[i] == nums[i-1] && used[i-1] {
			break
		}

		*builder = append(*builder, nums[i])
		used[i] = true
		backtracePermuteUnique(idx+1, used, res, builder, nums)
		*builder = (*builder)[:len(*builder)-1]
		used[i] = false
	}

}

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	builder := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	sort.Ints(nums)
	backtracePermuteUnique(0, used, &res, &builder, nums)
	return res
}
