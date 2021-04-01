package V1

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

	if len(nums) <= 1 {
		return len(nums)
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}

	}

	return slow + 1
}
