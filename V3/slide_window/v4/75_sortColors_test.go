package v4

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {

	nums := []int{2, 0, 2, 1, 1, 0}

	sortColors(nums)

	fmt.Println(nums)
}

func swap(a, b *int) {

	if *a != *b {

		*a = *a ^ *b
		*b = *a ^ *b
		*a = *a ^ *b
	}

}

func sortColors(nums []int) {

	redArea := -1
	blueArea := len(nums)

	for i := 0; i < blueArea; {

		if nums[i] == 0 {
			redArea++
			swap(&nums[i], &nums[redArea])
		} else if nums[i] == 2 {
			blueArea--
			swap(&nums[i], &nums[blueArea])
			continue
		}
		i++
	}
}
