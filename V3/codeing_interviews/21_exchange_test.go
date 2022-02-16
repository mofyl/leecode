package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestExchange(t *testing.T) {

	nums := []int{1, 2, 3, 4}

	res := exchange(nums)

	fmt.Println(res)
}

func exchange(nums []int) []int {

	// 奇数 放在前面 偶数放在后面
	even := len(nums)

	for i := 0; i < even; {

		// 奇数
		if nums[i]&1 == 1 {
			i++
		} else {
			even--
			nums[even], nums[i] = nums[i], nums[even]
		}

	}

	return nums

}
