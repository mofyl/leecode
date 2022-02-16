package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestSingleNumbers(t *testing.T) {

	//nums := []int{4, 1, 4, 6}

	nums := []int{1, 2, 10, 4, 1, 4, 3, 3}

	res := singleNumbers(nums)

	fmt.Println(res)
}

func singleNumbers(nums []int) []int {

	// 先将两个数字 异或到一起
	n := 0

	for i := 0; i < len(nums); i++ {
		n ^= nums[i]
	}
	div := 1
	//
	/*
		根据异或的结果可以得到 两个数字从右向左第一个不同的bit，在那个位置上
		异或的结果 如果两个位 相同 则为0， 不同则 为1，
		若某一位位 1  则表示该位置不同，通过下面的方式就找到了不同的位置
	*/
	for n&div == 0 {
		div <<= 1
	}

	// 根据 该位置将这两个数字分开
	/*
		我们得到了 该位置，
		我们使用每个数字去 & 该位置，会得到两种结果。就可以将两个不同的数字分开了
			那些出现两次的数字，一定会分到同一组，那么在同一组中 我们在进行异或，就可以将重复的数字抵消掉
	*/
	a := 0
	b := 0

	for i := 0; i < len(nums); i++ {

		// 这里通过这两种结果 就可以将两个不同的数字分开了。在分开的过程中 也会将
		// 出现两次的数字分到一起，那么我们利用 异或 就可以将这些出现两次的数字 抵消掉
		if nums[i]&div == 0 {
			//
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}

	return []int{a, b}
}
