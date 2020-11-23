package V3

import (
	"fmt"
	"sort"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{-1, 4, 2, 1, 9, 10}
	// nums := []int{3, 4, -1, 1}
	// nums := []int{7, 8, 9, 11, 12}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive_V1(nums []int) int {

	sort.Ints(nums)
	res := 1

	for i := 0; i < len(nums); i++ {

		if nums[i] == res {
			res++
		}

	}
	return res
}

func firstMissingPositive(nums []int) int {

	if len(nums) == 0 {
		return 1
	}

	// 缺失的第一个正数一定在 [1,N+1] 这个范围
	n := len(nums)
	// 先找位置
	// 这里是将 nums[i] 位置上的数 移动到 nums[i]-1 位置上 eg: nums[i]=3 那么该数字就要移动到 下标为2的位置上
	for i := 0; i < len(nums); i++ {
		// 这里首先要保证 nums[i] > 0 若不满足 则表示 该数 是负数或0 则不用考虑
		// 若是 nums[i] >= n 表示这个数 太大了 超出了范围 也不用考虑
		for nums[i] > 0 && nums[i] < n && nums[nums[i]-1] != nums[i] {
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		}
	}
	fmt.Println(nums)
	// 找到位置后就直接遍历
	for i := 0; i < n; i++ {
		// 由于找位置的时候 是将 3 保存到 下标为2 的位置，所以这里的下标需要+1 才会等于原来的值
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
