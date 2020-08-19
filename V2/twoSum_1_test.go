package V2

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))
	res := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		v, ok := m[num]
		if ok {
			res = append(res, i)
			res = append(res, v)
			return res
		}
		m[nums[i]] = i
	}

	return res

}
