package V3

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	// nums := []int{1, 2}
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(nums)
	fmt.Println(n)

	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {

	left := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[left] {
			left++
			nums[left] = nums[i]
		}
	}

	return left + 1
}
