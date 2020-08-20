package V2

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestSumClosest(t *testing.T) {
	fmt.Println("asda")
}

func threeSumClosest(nums []int, target int) int {
	res := 0

	numsLen := len(nums)

	sort.Ints(nums)

	finAbs := math.MaxFloat64
	// 题目中没有说 不能重复，所以不用判断是否重复
	for i := 0; i < numsLen; i++ {

		l := i + 1
		r := numsLen - 1
		cur := nums[i]
		for l < r {

			sum := nums[l] + nums[r] + cur
			if sum == target {
				return sum
			}

			sub := target - sum
			abs := float64(sub)

			if abs < 0 {
				abs = math.Abs(abs)
			}
			// 这里为了避免判断正负，直接用绝对值
			if finAbs > abs {
				finAbs = abs
				res = sum
			}

			if sum > target {
				r--
			} else if sum < target {
				l++
			}

		}

	}

	return res
}
