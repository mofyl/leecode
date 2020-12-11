package array

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{-2, 1, 1}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {

	if nums == nil || len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			// 这里表示若两个数同号 [-2 , 1, 1] r 和 l就是同号的 所以这个条件不能要
			// if nums[l] * nums[r] > 0 {
			// 	break
			// }
			tmp := nums[l] + nums[r] + nums[i]
			if tmp == 0 {

				res = append(res, []int{nums[l], nums[r], nums[i]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if tmp > 0 {
				r--
			} else {
				l++
			}

		}

	}

	return res

}
