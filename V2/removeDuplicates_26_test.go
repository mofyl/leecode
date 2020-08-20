package V2

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	nums := []int{1, 1, 2}

	res := removeDuplicates(nums)

	fmt.Println(res)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {

	if nums == nil {
		return 0
	}

	if len(nums) < 1 {
		return 1
	}
	l := 0
	num := nums[l]
	for i := 1; i < len(nums); i++ {
		for i < len(nums) && num == nums[i] {
			i++
		}
		if l+1 < len(nums) && i < len(nums) {
			l++
			nums[l] = nums[i]
			num = nums[i]

		}
	}

	return l + 1

}
