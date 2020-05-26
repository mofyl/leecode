package main

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

}

func removeDuplicates(nums []int) int {
	i := 0

	for j := 1; j < len(nums); j++ {

		if nums[i] == nums[j] {
			continue
		}
		i++
		nums[i] = nums[j]

	}
	// fmt.Println(nums)
	return i + 1
}
