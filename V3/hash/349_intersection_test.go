package hash

import (
	"fmt"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	//nums1 := []int{1, 2, 2, 1}
	//nums2 := []int{2, 2}

	nums1 := []int{9, 4, 9, 8, 4}
	nums2 := []int{4, 9, 5}

	res := intersectionV2(nums1, nums2)

	fmt.Println(res)

}

func intersectionV2(nums1 []int, nums2 []int) []int {

	if len(nums1) < 1 || len(nums2) < 1 {
		return nil
	}
	// 这里值是bool 是为了保证结果数组中没有重复的元素
	// 由于可能nums2中有重复的元素，后面遍历的时候可能对map中元素进行重复访问。
	// 所以要增加一个标记，标记第一次访问
	m := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = false
	}

	res := make([]int, 0)

	for i := 0; i < len(nums2); i++ {

		// 这里表示 nums2的元素在 nums1中出现了,而且是第一次访问
		if v, ok := m[nums2[i]]; ok && !v {
			m[nums2[i]] = true
			res = append(res, nums2[i])
		}

	}

	return res

}

func intersection(nums1 []int, nums2 []int) []int {

	if len(nums1) < 1 || len(nums2) < 1 {
		return nil
	}

	m := make(map[int]struct{})

	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {

		if nums1[i] == nums2[j] {
			m[nums1[i]] = struct{}{}

			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}

	}

	if len(m) < 1 {
		return nil
	}

	res := make([]int, 0, len(m))

	for k, _ := range m {
		res = append(res, k)
	}

	return res

}
