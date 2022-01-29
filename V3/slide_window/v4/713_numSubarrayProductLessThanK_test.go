package v4

import (
	"fmt"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {

	nums := []int{10, 5, 2, 6}
	k := 100

	//nums := []int{1, 1, 1}
	//k := 1

	res := numSubarrayProductLessThanK(nums, k)

	fmt.Println(res)
}

func numSubarrayProductLessThanK(nums []int, k int) int {

	if k < 1 {
		return 0
	}

	curPro := 1
	res := 0
	l, r := 0, 0

	for ; r < len(nums); r++ {

		curPro *= nums[r]

		for l <= r && curPro >= k {
			curPro /= nums[l]
			l++
		}

		res += r - l + 1

	}

	return res

}
