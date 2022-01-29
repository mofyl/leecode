package v4

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestFindLength(t *testing.T) {

	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}

	//nums1 := []int{0, 0, 0, 0, 0}
	//nums2 := []int{0, 0, 0, 0, 0}

	//nums1 := []int{0, 0, 0, 0, 1}
	//nums2 := []int{1, 0, 0, 0, 0}

	//nums1 := []int{5, 14, 53, 80, 48}
	//nums2 := []int{50, 47, 3, 80, 83}

	res := findLength(nums1, nums2)

	fmt.Println(res)
}

func findLength(nums1 []int, nums2 []int) int {

	// 固定A  然后去移动B
	// 然后在 反过来 求一遍， 用最大的那个
	n := len(nums1)
	m := len(nums2)
	max := 0
	// 这里是 固定 nums2 ，遍历 nums1
	for i := 0; i < n; i++ {

		minLen := tools.Min(m, n-i)
		cur := findLengthHelper(nums1, nums2, i, 0, minLen)
		max = tools.Max(max, cur)
	}

	// 这里是 固定 nums1 , 遍历 nums2
	for i := 0; i < m; i++ {

		minLen := tools.Min(n, m-i)

		cur := findLengthHelper(nums1, nums2, 0, i, minLen)
		max = tools.Max(max, cur)

	}

	return max

}

func findLengthHelper(a, b []int, aStart, bStart int, step int) int {

	cur := 0
	max := 0
	for i := 0; i < step; i++ {

		if a[aStart+i] == b[bStart+i] {
			cur++
			max = tools.Max(max, cur)
		} else {
			cur = 0
		}

	}

	return max

}
