package v4

import (
	"fmt"
	"testing"
)

func Test_TwoSum(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {

	set := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {

		t := target - nums[i]

		if idx, ok := set[t]; ok {
			return []int{i, idx}
		} else {
			/*
				target - 被减数 = 减数
				target - 减数 = 被减数
				这里记录的是 当前数字 -> 下标
				因为 每次使用 (target - nums[i]) 直接到map中查询
			*/
			set[nums[i]] = i
		}
	}
	return nil
}
