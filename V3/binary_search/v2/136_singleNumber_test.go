package v2

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {

	nums := []int{4, 1, 2, 1, 2}
	res := singleNumber(nums)

	fmt.Println(res)
}

func singleNumber(nums []int) int {

	res := 0

	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res

}
