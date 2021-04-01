package v1

import (
	"fmt"
	"sort"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 2}

	res := subsetsWithDup(nums)

	fmt.Println(res)
}

func subsetsWithDup(nums []int) [][]int {

	if len(nums) < 1 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0, len(nums))
	res = append(res, []int{})
	used := make([]bool, len(nums))
	subsetsWithDupLoop(0, nums, []int{}, &res, used)
	return res
}

func subsetsWithDupLoop(start int, nums []int, curNum []int, res *[][]int, used []bool) {

	if start > len(nums) {
		return
	}

	if len(curNum) > 0 {
		tmp := make([]int, len(curNum))
		copy(tmp, curNum)
		*res = append(*res, tmp)
	}

	for i := start; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		curNum = append(curNum, nums[i])
		used[i] = true
		subsetsWithDupLoop(i+1, nums, curNum, res, used)
		used[i] = false
		curNum = curNum[:len(curNum)-1]
	}

}
