package backtracking

import (
	"fmt"
	"testing"
)

func TestSubTest(t *testing.T) {

	nums := []int{1, 2, 3}

	res := subsets(nums)

	fmt.Println(res)
}

func subsets(nums []int) [][]int {
	c := make([][]int, 0, len(nums))
	if len(nums) < 1 {
		return c
	}
	c = append(c, []int{})
	subsetsLoop(nums, 0, &[]int{}, &c)
	return c
}

func subsetsLoop(nums []int, start int, cur *[]int, res *[][]int) {

	if start > len(nums) {
		return
	}

	if len(*cur) > 0 {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
	}

	for i := start; i < len(nums); i++ {

		*cur = append(*cur, nums[i])
		subsetsLoop(nums, i+1, cur, res)
		*cur = (*cur)[:len(*cur)-1]
	}

}
