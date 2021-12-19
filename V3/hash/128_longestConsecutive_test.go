package hash

import (
	"fmt"
	"leecode/tools"
	"math"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {

	nums := []int{100, 4, 200, 1, 3, 2}
	//nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	res := longestConsecutive(nums)
	fmt.Println(res)
}

func longestConsecutive(nums []int) int {

	// 只有一个数字的情况
	if len(nums) < 2 {
		return len(nums)
	}

	// 先将nums  录入hash 中 加速查找

	helpMap := make(map[int]struct{}, len(nums))

	for i := 0; i < len(nums); i++ {
		helpMap[nums[i]] = struct{}{}
	}

	maxLen := math.MinInt64

	for i := 0; i < len(nums); i++ {

		_, ok := helpMap[nums[i]-1]

		if !ok {
			// 这里起始值为1，由于 nums[i] 也是序列中的一个数字,所以要算上nums[i]
			count := 1
			curNum := nums[i] + 1
			_, ok1 := helpMap[curNum]
			for ok1 {
				count++
				curNum++
				_, ok1 = helpMap[curNum]
			}

			maxLen = tools.Max(maxLen, count)
		}

	}

	return maxLen

}
