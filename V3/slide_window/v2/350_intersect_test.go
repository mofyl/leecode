package v2

import (
	"fmt"
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {

	//nums1 := []int{1, 2, 2, 1}
	//nums2 := []int{2, 2}

	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	//res := intersect(nums1, nums2)
	res := intersectV2(nums1, nums2)

	fmt.Println(res)

}

func intersect(nums1 []int, nums2 []int) []int {

	if len(nums1) < 1 || len(nums2) < 1 {
		return nil
	}

	m := make(map[int]int, 0)

	for i := 0; i < len(nums1); i++ {

		v := m[nums1[i]]
		v++
		m[nums1[i]] = v
	}

	res := make([]int, 0, len(nums2))

	for i := 0; i < len(nums2); i++ {

		v, ok := m[nums2[i]]

		if ok && v > 0 {
			res = append(res, nums2[i])
			v--
			m[nums2[i]] = v
		}

	}

	return res

}

func intersectV2(nums1 []int, nums2 []int) []int {

	if len(nums1) < 1 || len(nums2) < 1 {
		return nil
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	l1, l2 := 0, 0

	res := make([]int, 0, len(nums1))

	for l1 < len(nums1) && l2 < len(nums2) {

		if nums1[l1] < nums2[l2] {
			l1++
		} else if nums1[l1] > nums2[l2] {
			l2++
		} else {
			res = append(res, nums1[l1])
			l1++
			l2++
		}

	}
	return res
}
