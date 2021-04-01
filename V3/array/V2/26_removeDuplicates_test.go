package V2

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}

	res := removeDuplicates(nums)
	fmt.Println(res)
}

func removeDuplicates(nums []int) int {

	if len(nums) < 1 {
		return 0
	}

	l := 0

	for i := 1; i < len(nums); i++ {

		if nums[i] != nums[l] {
			l++
			nums[l] = nums[i]
		}

	}

	return l + 1

}
