package V1

import (
	"fmt"
	"testing"
)

func TestSignalNum(t *testing.T) {
	nums := []int{-401451, -177656, -2147483646, -473874, -814645, -2147483646, -852036, -457533, -401451, -473874, -401451, -216555, -917279, -457533, -852036, -457533, -177656, -2147483646, -177656, -917279, -473874, -852036, -917279, -216555, -814645, 2147483645, -2147483648, 2147483645, -814645, 2147483645, -216555}

	fmt.Println(singleNumber2(nums))
}

func singleNumber2(nums []int) int {

	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	if numsLen == 1 {
		return nums[0]
	}

	num := nums[0]
	count := 1
	for i := 1; i < numsLen; i++ {
		if num == nums[i] {
			count++
			continue
		}

		if count != 1 {
			num = nums[i]
			count = 1
		}

	}
	return num

}
