package V2

import (
	"testing"
)

func TestThreeSum(t *testing.T) {

}

//func threeSum(nums []int) [][]int {
//
//	if len(nums) < 1 {
//		return nil
//	}
//
//	sort.Ints(nums)
//
//	for i := 0; i < len(nums); i++ {
//		// 由于我们已经排序了，最小的数字一定在前面
//		// 若第一个数字就 >0 则 一定不会出现结果为0的组
//		if nums[i] > 0 {
//			continue
//		}
//		// 排除重复元素
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//
//	}
//
//}
