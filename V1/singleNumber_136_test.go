package V1

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}
