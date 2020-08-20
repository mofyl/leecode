package V2

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	res := threeSum(nums)

	fmt.Println(res)
}

func threeSum(nums []int) [][]int {

	res := make([][]int, 0, len(nums))
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)
	lenNums := len(nums)

	for i := 0; i < lenNums; i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := lenNums - 1
		cur := nums[i]

		for l < r {
			tmp := cur + nums[l] + nums[r]

			if tmp == 0 {
				s := make([]int, 0, 3)
				s = append(s, cur)
				s = append(s, nums[l])
				s = append(s, nums[r])
				res = append(res, s)
				// 去掉重复的元素
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}
				// 由于上面判断的是 l+1 和r-1 若上面的循环退出的话，
				// l和r的位置还是停在 已经使用过的元素上，这样就会重复，所以这里才需要跳过
				l++
				r--

			} else if tmp < 0 {
				l++
			} else {
				r--
			}
		}

	}

	return res

}
