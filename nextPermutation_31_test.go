package main

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	// 12385764
	nums := []int{1, 2, 3, 8, 7, 6, 4}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {

	/*

		解题思路：
			若该数组是最大序列 则直接将该数组按照升序重新排列即可
			若该数不是最大序列。需要将后面的大数和前面的小数交换，还要保证增加的幅度尽可能的小，
				那么就要大数尽可能的小,所以我们就需要从后向前遍历
			交换完成后还需要将后面的数按升序重新排列，这样排序过后的数就是最小的
			找下一个最大的序列步骤
				先找到需要更换的地方i
				然后找到需要更换的数字k ，交换k ， i
				将 [j , end) 升序排序
	*/
	lNums := len(nums)
	i, j, k := lNums-2, lNums-1, lNums-1
	// 找到需要更换的地方，就是 nums[i] < nums[j] ，比如 1，2，3。2就是需要交换的位置，

	// 若是最大序列 则 这里 i会变为-1
	// 这里结束以后 i指向的就是 需要更换的地方
	// 这里其实可以保证 [j , end) 是降序的
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		// 这里表示,不是最大序列
		// 寻找第一个比nums[i] 大的数 nums[k] ,然后将nums[k] 和 nums[i] 交换

		for i >= 0 && nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 由于是降序的，则可以直接交换，将最小的和最大的进行交换
	for i, j := j, lNums-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

}
