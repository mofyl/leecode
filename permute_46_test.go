package main

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	num := []int{1, 2, 3}

	fmt.Println(permute(num))
}

func backtracePermute(idx int, res *[][]int, builder *[]int, nums []int) {

	if idx == len(nums) {
		newBuilder := make([]int, len(*builder))
		copy(newBuilder, *builder)
		*res = append(*res, newBuilder)
		(*builder) = (*builder)[:]
		return
	}

	for i := 0; i < len(nums); i++ {

		// 剪枝
		if intContains(*builder, nums[i]) {
			continue
		}

		*builder = append(*builder, nums[i])
		backtracePermute(idx+1, res, builder, nums)
		*builder = (*builder)[:len(*builder)-1]
	}

}

func intContains(nums []int, num int) bool {

	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return true
		}
	}

	return false
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	builder := make([]int, 0, len(nums))

	backtracePermute(0, &res, &builder, nums)
	return res
}
