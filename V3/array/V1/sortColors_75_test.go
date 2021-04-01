package V1

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	//nums := []int{2, 0, 2, 1, 1, 0}
	nums := []int{2, 0, 1}
	sortColors(nums)

	fmt.Println(nums)
}

func sortColors(nums []int) {

	left := 0
	i := 0
	right := len(nums) - 1

	for i <= right {
		if nums[i] == 0 {
			swap(nums, i, left)
			left++
			i++
		} else if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			swap(nums, i, right)
			right--
		}
	}

}
func swap(nums []int, a, b int) {
	tmp := nums[a]
	nums[a] = nums[b]
	nums[b] = tmp
}
