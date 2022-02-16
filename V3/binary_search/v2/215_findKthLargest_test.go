package v2

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFindKthLargest(t *testing.T) {

	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2

	//nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	//k := 4

	//nums := []int{1}
	//k := 1

	res := findKthLargest(nums, k)

	fmt.Println(res)
}

func findKthLargest(nums []int, k int) int {

	rand.Seed(time.Now().UnixNano())
	// 这里要找的就是 排序后 下标为 targetIdx 的数字
	targetIdx := len(nums) - k

	l := 0
	r := len(nums) - 1

	for l < r {

		mid := findKthLargestHelper(nums, l, r)

		if mid == targetIdx {
			return nums[mid]
		}

		if mid > targetIdx {
			r = mid - 1
		} else if mid < targetIdx {
			l = mid + 1
		}

	}

	return nums[l]
}

func findKthLargestHelper(nums []int, l, r int) int {

	idx := l + rand.Intn(r-l+1)
	midN := nums[idx]
	nums[l], nums[idx] = nums[idx], nums[l]
	//begin := l
	//end := r
	for l < r {

		for l < r && nums[r] > midN {
			r--
		}

		if l < r {
			nums[l] = nums[r]
			l++
		}

		for l < r && midN > nums[l] {
			l++
		}

		if l < r {
			nums[r] = nums[l]
			r--
		}

	}

	nums[l] = midN

	return l

}
