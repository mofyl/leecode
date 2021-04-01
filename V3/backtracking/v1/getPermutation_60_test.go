package v1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestGetPermutation(t *testing.T) {
	n := 3
	k := 3
	//n := 4
	//k := 9
	//n := 3
	//k := 1
	str := getPermutation(n, k)
	fmt.Println(str)
}

func getPermutation(n int, k int) string {
	//
	//nums := make([]int, 0, n)
	//
	//for i := 1; i <= n; i++ {
	//	nums = append(nums, i)
	//}
	//curK := 0
	//str := ""

	//used := make([]bool, n)
	//getPermutationLoop(0, nums, []int{}, &curK, &str, used, k)
	//getPermutationLoopV2(nums, []int{}, &curK, &str, k)
	str := getPermutationLoopV3(n, k)
	return str
}

func getPermutationLoop(start int, num []int, curNum []int, curK *int,
	str *string, used []bool, k int) {

	if start == len(num) {
		*curK++
		if *curK == k {

			b := strings.Builder{}

			for i := 0; i < len(curNum); i++ {
				b.WriteString(strconv.Itoa(curNum[i]))
			}
			*str = b.String()
			return
		}

		return
	}

	for i := 0; i < len(num); i++ {

		if *curK == k {
			return
		}

		if used[i] {
			continue
		}

		curNum = append(curNum, num[i])
		used[i] = true
		getPermutationLoop(start+1, num, curNum, curK, str, used, k)
		used[i] = false
		curNum = curNum[:len(curNum)-1]
	}

}

func getPermutationLoopV2(num []int, curNum []int, curK *int,
	str *string, k int) {

	if len(num) < 1 {
		*curK++
		if *curK == k {
			b := strings.Builder{}
			for i := 0; i < len(curNum); i++ {
				b.WriteString(strconv.Itoa(curNum[i]))
			}
			*str = b.String()
			return
		}

		return
	}

	for i := 0; i < len(num); i++ {

		if *curK == k {
			return
		}

		curNum = append(curNum, num[i])
		tmp := []int{}
		tmp = append(tmp, num[:i]...)
		tmp = append(tmp, num[i+1:]...)
		getPermutationLoopV2(tmp, curNum, curK, str, k)
		curNum = curNum[:len(curNum)-1]
	}

}

func factorial(num int) int {
	res := 1
	for num > 1 {
		res *= num
		num--
	}
	return res
}

func getPermutationLoopV3(n int, k int) string {

	// 这里采用数学的方法
	/*
		每个数字的 全排列都是相同的为 (n-1)!  一共会有 (n-1)! * n 种 所以这里我们可以将 每个数字的排列
			看成一组。我们用k/(n-1)!  就得出了当前数字在那个组中。 然后到对应的nums数组中寻找该数字就好了。
			注意这里i是从0 开始的， k需要减一，
	*/
	res := ""
	k--

	factorialNum := make([]int, n+1)
	factorialNum[0] = 1
	nums := make([]int, n)

	for i := 1; i <= n; i++ {
		factorialNum[i] = i * factorialNum[i-1]
		nums[i-1] = i
	}

	for i := 0; i < n; i++ {
		// 首先算出 当前 n-1 的阶层是多少
		f := factorialNum[n-1-i] // 由于i是从0 开始的 n取不到
		// 这里算出 当前阶层 应该在那个组中
		a := k / f
		num := nums[k/f]
		// 从 nums中去掉这个数字
		nums = append(nums[:a], nums[a+1:]...)
		res = strings.Join([]string{res, strconv.Itoa(num)}, "")

		k %= f
	}

	return res
}
