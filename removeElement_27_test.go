package main

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	t.Log(removeElement(nums, 3))
}

func removeElement(nums []int, val int) int {

	i := 0
	n := len(nums)

	for i < n {

		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}

	}

	return n

}
