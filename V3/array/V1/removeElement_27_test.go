package V1

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	// nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// nums := []int{2}
	target := 3
	res := removeElement(nums, target)
	fmt.Println(res, nums)
}

func removeElement(nums []int, val int) int {

	i := 0
	n := len(nums)

	for i < n {

		if nums[i] == val {
			n--
			nums[i] = nums[n]
		} else {
			i++
		}
	}

	return i

}
