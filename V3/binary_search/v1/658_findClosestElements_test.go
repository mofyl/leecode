package v1

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestFindClosestElements(t *testing.T) {

	//arr := []int{1, 2, 3, 4, 5}
	//k := 4
	//x := 3

	arr := []int{1, 2, 3, 4, 5}
	k := 4
	x := -1

	res := findClosestElements(arr, k, x)

	fmt.Println(res)
}

func findClosestElements(arr []int, k int, x int) []int {

	// 由于数据是排好序的，查找的时候 从头和末尾 两边一起找就好了，
	// 找的时候 不断的进行比较 左右数字 和 x的差值 abs。每次收缩大的那一边
	// 然后一直到 l,r 之间数字数量 刚好是k个

	if k > len(arr) {
		return nil
	}

	l, r := 0, len(arr)-1

	for (r - l + 1) > k {

		a := tools.Abs(int32(arr[l] - x))
		b := tools.Abs(int32(arr[r] - x))

		if a > b {
			l++
		} else {
			r--
		}

	}
	// 因为这里需要的是 l 和 r 的闭区间，所以要 r+1
	return arr[l : r+1]
}
