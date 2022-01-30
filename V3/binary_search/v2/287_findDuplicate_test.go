package v2

import (
	"fmt"
	"testing"
)

func TestFindDuplicate(t *testing.T) {

	//nums := []int{1, 3, 4, 2, 2}

	//nums := []int{3, 1, 3, 4, 2}

	nums := []int{1, 1, 2}

	res := findDuplicate(nums)

	fmt.Println(res)
}

// 由于  1 <= nums[i] <= n , n为nums 的长度
func findDuplicate(nums []int) int {

	tmp := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {

		if tmp[nums[i]] != 0 {
			return nums[i]
		}

		tmp[nums[i]] = nums[i]
	}

	return -1

}
