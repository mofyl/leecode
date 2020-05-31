package main

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {

	num1 := make([]int, 5)
	num1[0] = 0
	// num1[1] = 5

	num2 := make([]int, 2)
	num2[0] = 1
	// num2[1] = 4
	merge(num1, 0, num2, 1)
	fmt.Println(num1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	/*

		已知 nums1 和 nums2 都是有序的， 那么无论从0~m-1 或者从 m-1~0
			他们都是有序的。
		那么我们就可以从后向前遍历，
			若nums1[m-1]> nums2[n-1] 那么就应该将该 nums[m-1] 放到 nums[m+n-1] 的位置上去
			若nums1[m-1] <= nums2[n-1] 那么就应该将 nums[n-1] 放到 nums[m+n-1] 的位置上
		最后处理若 num2 最后有剩余元素，那么就直接将剩下的元素放到nums1[n-1] 的位置上

	*/
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	if n > 0 {
		for n > 0 {
			nums1[n-1] = nums2[n-1]
			n--
		}
	}

}

type Str struct {
	A int
	B int
}

func (s *Str) Func1() {
	s.A = 11
	fmt.Println(111)
}

func (s Str) Func2() {
	s.A = 1
	fmt.Println(222)
}

func TestTemp(t *testing.T) {

	s1 := Str{}
	s2 := new(Str)

	s1.Func2()
	s2.Func2()

	fmt.Println(s1)
	fmt.Println(s2)
}
