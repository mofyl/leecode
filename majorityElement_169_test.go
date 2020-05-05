package main

import (
	"fmt"
	"testing"
)

func TestMajorityEle(t *testing.T) {
	// nums := []int{3, 2, 3}
	nums := []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Println(majorityElement(nums))

}

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	/*
		由于出现次数多的元素一定比n/2多
		那么经过排序以后相同的元素就都到了一起形成子串
		这时找n/2位置的数 就一定是最多的那个，因为若不是最多的那个数是不会有这么长的长度的

	*/
	// sort.Ints(nums)
	// return nums[len(nums)>>1]

	/*
		摩尔投票法：
			首先假设一个候选者，然后该候选者的次数初始化为1，
				然后遍历数组，若出现该候选者就将 次数+1，若出现的是别的数 就将该次数-1
				若减为0，则更换候选者

	*/

	candNum := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if candNum == nums[i] {
			count++
		} else {
			count--

			if count == 0 {
				candNum = nums[i]
				count = 1
			}
		}
	}

	return candNum

}
