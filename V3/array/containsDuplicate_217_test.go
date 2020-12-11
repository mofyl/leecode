package array

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {

	//nums := []int{1, 2, 3, 1}
	nums := []int{}

	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {

	// 这里使用hash表来去重
	m := make(map[int]struct{}, len(nums))

	for i := 0; i < len(nums); i++ {

		_, ok := m[nums[i]]
		if !ok {
			m[nums[i]] = struct{}{}
		} else {
			return true
		}
	}
	return false
}
