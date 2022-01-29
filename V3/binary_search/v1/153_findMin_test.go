package v1

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {

	//nums := []int{3, 4, 5, 1, 2}
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{11, 13, 15, 17}
	res := findMin(nums)
	fmt.Println(res)
}

func findMin(nums []int) int {

	// 题目规定了 nums的长度会 大于等于1
	if len(nums) < 2 {
		return nums[0]
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}

	// 若整个数组中都是 nums[i] > nums[i-1] 那么就证明整个数组有序
	// 最小的就是第一个元素
	return nums[0]
}

func findMinV2(nums []int) int {

	// 题目规定了 nums的长度会 大于等于1
	if len(nums) < 2 {
		return nums[0]
	}
	// 由于数据未旋转的时候 是整体有序的。而且每个数字不相同
	// 在旋转之后，后面的数字放到了前面, 1,2,3,4,5 旋转之后 3,4,5,1,2
	// 可以发现 前面一段一定是大于 nums[0]的。后面一段一定是小于nums[0]的
	// 若全部元素都大于nums[0]，那么就表示nums[0]是最小的
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[0] {
			// 也就说明了  nums[i]是拐点
			return nums[i]
		}
	}

	// 若整个数组中都是 nums[i] > nums[i-1] 那么就证明整个数组有序
	// 最小的就是第一个元素
	return nums[0]
}
