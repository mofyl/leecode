package array

import (
	"fmt"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	//nums := []int{1, 2, 3}
	//nums := []int{5, 4, 6, 7, 9, 3, 10, 9, 5, 6}
	res := findDuplicates(nums)

	fmt.Println(res)
}

func findDuplicates(nums []int) []int {

	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		// 由于题目给定了 1<= a[i] <= n 所以这里不用判断
		for nums[nums[i]-1] != nums[i] {
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		}
	}
	// 由于上面是交换 所以重复的数字永远会 占用着 不属于他的位置  这样我们就可以判断
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, nums[i])
		}
	}

	return res
}
