package V1

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	target := 1
	res := threeSumClosest(nums, target)
	fmt.Println(res)
}

func threeSumClosest(nums []int, target int) int {

	if nums == nil || len(nums) < 3 {
		return -1
	}

	sort.Ints(nums)
	res := 0
	abs := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		// 这里也是要 跳过重复项，应为重复项 计算出来的值都一样
		// 就没有什么意义
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {

			tmp := nums[i] + nums[l] + nums[r]
			if tmp == target {
				return tmp
			}
			// target - tmp的值越小 ，表示越接近 target
			// 而 res 就需要 最接近 target的那个数字
			a := getAbs(target - tmp)
			if a < abs {
				abs = a
				res = tmp
			}

			// 这里也是为了逼近 target
			if tmp > target {
				r--
			} else {
				l++
			}
		}

	}

	return res

}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
