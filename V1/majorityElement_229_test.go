package V1

import (
	"fmt"
	"testing"
)

func TestMajorityEles(t *testing.T) {
	// nums := []int{3, 2, 3}
	nums := []int{2, 2}

	fmt.Println(majorityElements(nums))
}

func majorityElements(nums []int) []int {

	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}

	if numsLen == 1 {
		return nums
	}
	cend1 := nums[0]
	count1 := 0
	cend2 := nums[0]
	count2 := 0
	for i := 0; i < numsLen; i++ {
		if cend1 == nums[i] {
			count1++
			continue
		}

		if cend2 == nums[i] {
			count2++
			continue
		}

		if count1 == 0 {
			cend1 = nums[i]
			count1 = 1
			continue
		}

		if count2 == 0 {
			cend2 = nums[i]
			count2 = 1
			continue
		}

		count1--
		count2--

	}
	res := make([]int, 0, 2)
	// 统计阶段
	// 统计 cend1
	if count1 > 0 {
		if sumNum(nums, cend1) > numsLen/3 {
			res = append(res, cend1)
		}
	}
	if count2 > 0 {
		if sumNum(nums, cend2) > numsLen/3 {
			res = append(res, cend2)
		}
	}
	return res

}

func sumNum(nums []int, num int) int {
	sum := 0
	for _, v := range nums {
		if num == v {
			sum++
		}
	}

	return sum
}
