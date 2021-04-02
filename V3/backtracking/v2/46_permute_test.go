package v2

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {

	nums := []int{1, 1, 2}

	res := permute(nums)
	fmt.Println(res)
}

func permute(nums []int) [][]int {

	if len(nums) < 1 {
		return nil
	}

	res := make([][]int, 0)

	permuteHelper(nums, 0, &res)
	return res
}

func swap(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}

// 递归函数表示  在idx 这个位置上 所有的数字都可以放到这里1
func permuteHelper(nums []int, idx int, res *[][]int) {

	if idx >= len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := idx; i < len(nums); i++ {

		swap(nums, i, idx)

		permuteHelper(nums, idx+1, res)
		swap(nums, i, idx)
	}

}
