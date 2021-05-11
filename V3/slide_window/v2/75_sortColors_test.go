package v2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestSortColors(t *testing.T) {

	//nums := []int{2, 0, 2, 1, 1, 0}
	//nums := []int{2, 0, 1}
	nums := []int{1, 2, 0}

	sortColors(nums)

	fmt.Println(nums)
}

/*

	荷兰国旗问题：


*/
func sortColors(nums []int) {

	if len(nums) < 2 {
		return
	}

	l := -1
	r := len(nums)
	i := 0
	for l < r && i < r {

		if nums[i] == 0 {
			l++
			tools.Swap(&nums[i], &nums[l])
		} else if nums[i] == 2 {
			for l < r && i < r && nums[i] == 2 {
				r--
				tools.Swap(&nums[i], &nums[r])
			}
			continue
		}
		i++
	}

}
