package backtracking

import (
	"fmt"
	"sort"
	"testing"
)

func TestPermuteUnique(t *testing.T) {

	//nums := []int{1, 1, 2}
	//nums := []int{1, 2, 3}
	nums := []int{3, 3, 0, 3}

	res := permuteUnique(nums)

	fmt.Println(res)
}

func permuteUnique(nums []int) [][]int {

	res := make([][]int, 0, len(nums))
	used := make([]bool, len(nums))
	sort.Ints(nums)
	permuteUniqueLoop(0, nums, []int{}, &res, used)
	return res
}

func permuteUniqueLoop(start int, nums []int, curNums []int, res *[][]int, used []bool) {

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

		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		used[i] = true
		curNums = append(curNums, nums[i])
		permuteUniqueLoop(start+1, nums, curNums, res, used)
		curNums = curNums[:len(curNums)-1]
		used[i] = false
	}

}
