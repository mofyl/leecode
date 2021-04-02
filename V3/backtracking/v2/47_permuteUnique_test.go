package v2

import (
	"fmt"
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

	if len(nums) < 1 {
		return nil
	}
	res := make([][]int, 0)

	permuteUniqueHelper(nums, 0, &res)

	return res
}

func permuteUniqueHelper(nums []int, idx int, res *[][]int) {

	if idx == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	visit := make(map[int]struct{})

	for i := idx; i < len(nums); i++ {

		_, ok := visit[nums[i]]

		if !ok {
			visit[nums[i]] = struct{}{}
			swap(nums, i, idx)
			permuteUniqueHelper(nums, idx+1, res)
			swap(nums, i, idx)
		}

	}

}
