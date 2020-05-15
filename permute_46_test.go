package main

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	num := []int{1, 2, 3}

	fmt.Println(permute(num))
}

func backtracePermute(idx int, curNum int, res *[][]int, builder *[]int, nums []int) {

	if idx == len(nums)-1 {
		newBuilder := make([]int, len(*builder))
		copy(newBuilder, *builder)
		*res = append(*res, newBuilder)
		(*builder) = (*builder)[:]
		return
	}

	for i := 0; i < len(nums); i++ {
		if curNum == nums[i] {
			continue
		}
		*builder = append(*builder, nums[i])
		backtracePermute(idx+1, nums[i], res, builder, nums)
		*builder = (*builder)[:len(*builder)-1]
	}

}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	builder := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		builder = append(builder, nums[i])
		backtracePermute(0, nums[i], &res, &builder, nums)
		builder = builder[:len(builder)-1]
	}

	return res
}
