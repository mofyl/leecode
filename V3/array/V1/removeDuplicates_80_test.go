package V1

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates2(t *testing.T) {

	//nums := []int{1, 1, 1, 2, 2, 3}
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	//n := removeDuplicates2(nums)
	n := removeDuplicates3(nums)
	fmt.Println(n)
	fmt.Println(nums)
}

func removeDuplicates2(nums []int) int {

	fast := 1
	slow := 0
	count := 0
	for ; fast < len(nums); fast++ {

		if nums[fast] == nums[fast-1] {
			count++
		} else {
			count = 0
		}

		if count < 2 {
			slow++
			nums[slow] = nums[fast]
		}

	}

	return slow + 1

}

func removeDuplicates3(nums []int) int {

	if len(nums) <= 2 {
		return len(nums)
	}

	slow := 2
	fast := 2

	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow

}
