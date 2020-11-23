package V3

import (
	"fmt"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	res := findDisappearedNumbers(nums)

	fmt.Println(res)
}

func findDisappearedNumbers(nums []int) []int {

	if len(nums) == 0 {
		return nums
	}

	for i := 0; i < len(nums); i++ {

		for nums[nums[i]-1] != nums[i] {
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		}
	}
	res := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {

		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res

}
