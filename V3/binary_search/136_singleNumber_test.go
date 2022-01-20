package binary_search

import (
	"fmt"
	"sort"
	"testing"
)

func TestSingleNumber(t *testing.T) {

	nums := []int{4, 1, 2, 1, 2}

	res := singleNumber_bs(nums)

	fmt.Println(res)
}

// 这里使用 异或的方式做， O(n) 时间复杂度
func singleNumber(nums []int) int {

	// 由于只有一个数字出现了一次，其他数字出现了两次
	// 那么异或就可以 将其他 出现两次的消除掉
	res := 0

	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res

}

// 这里使用 二分来做，O( log(n) )
func singleNumber_bs(nums []int) int {

	sort.Ints(nums)

	// 由于该 数组一定是包含 只出现一次的数字的
	// 那么就可以先排序 再二分，这里二分 一定是 一半是奇数个， 一半是偶数个
	// eg: [1,1,2,2,4]  排序后 二分 为  [1,1,2] [2,4] 或  [1,1] [2,2,4]
	// 若 aHalf[len(aHalf)-1] == bHalf[0] 那么表示 这个只出现一次的元素一定在 长度为偶数的那一半中
	// 若 aHalf[len(aHalf)-1] != bHalf[0] 那么表示 这个只出现一次的元素 一定在 长度为奇数的那一半中

	for {

		if len(nums) <= 1 {
			return nums[0]
		}

		mid := len(nums) >> 1

		aHalf := nums[:mid+1]
		bHalf := nums[mid+1:]

		if aHalf[len(aHalf)-1] == bHalf[0] {
			// 表示在长度为 偶数那一组中,这里注意偶数的时候要把 已经比较过的去掉
			if len(aHalf)&1 == 0 {
				nums = aHalf[:len(aHalf)-1]
			} else {
				nums = bHalf[1:]
			}

		} else {
			// 不相等 就在奇数的那一组中
			if len(aHalf)&1 == 1 {
				nums = aHalf
			} else {
				nums = bHalf
			}

		}

	}

}
