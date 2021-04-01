package v1

import (
	"fmt"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	//nums := []int{10, 5, 2, 6}
	//k := 100

	nums := []int{1, 1, 1}
	k := 1 //这里没有等于 所以是 结果是 0

	res := numSubarrayProductLessThanK(nums, k)

	fmt.Println(res)
}

func numSubarrayProductLessThanK(nums []int, k int) int {

	if len(nums) < 1 || k == 0 {
		return 0
	}

	count := 0

	l, r := 0, 0
	div := 1
	for r < len(nums) {

		div *= nums[r]
		// 由于这里 l <= r
		// 跳出循环后  l > r 所以 r-l = -1 +1 = 0
		for div >= k && l <= r {
			div /= nums[l]
			l++
		}
		count += r - l + 1

		r++
	}

	return count
}
