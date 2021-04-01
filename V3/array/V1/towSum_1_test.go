package V1

import (
	"fmt"
	"testing"
)

func TestTowSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		v, ok := tmpMap[target-nums[i]]
		if ok {
			return []int{i, v}
		}

		tmpMap[nums[i]] = i
	}

	return []int{}
}
