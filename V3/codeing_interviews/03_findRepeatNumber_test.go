package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {

	nums := []int{2, 3, 1, 0, 2, 5, 3}

	res := findRepeatNumber_v2(nums)

	fmt.Println(nums)
	fmt.Println(res)
}

// 数据特点： 0 <= nums[i] <= len(nums)-1
func findRepeatNumber(nums []int) int {

	tmp := make([]int, len(nums))

	for i := 0; i < len(tmp); i++ {
		tmp[i] = -1
	}

	for i := 0; i < len(nums); i++ {

		if tmp[nums[i]] != -1 {
			return nums[i]
		}

		tmp[nums[i]] = nums[i]
	}

	return -1
}

func findRepeatNumber_v2(nums []int) int {

	for i := 0; i < len(nums); {
		// 若 当前元素 和 当前 i 的值相等 说明 nums[i] 放对了
		if nums[i] == i {
			i++
			continue
		}
		// 若 nums[ nums[i] ] 和 nums[i] 的值相等 说明 出现了 重复
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}

		// 交换 将 nums[i] 交换到正确 的位置上 正确的位置就是  nums[nums[i]]
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		// 这里 i 的位置没有变，一直在循环
	}
	return -1
}
