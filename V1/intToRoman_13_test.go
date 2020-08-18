package V1

import (
	"fmt"
	"testing"
)

var (
	romanMap = map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	romanSlice = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
)

func TestIntToRoman(t *testing.T) {
	res := intToRoman(3000)
	fmt.Println(res)
}

func compant(num int, c string, s *string) {
	for i := 0; i < num; i++ {
		(*s) += (c)
	}
}

func intToRoman(num int) string {

	roman, ok := romanMap[num]
	if ok {
		return roman
	}
	res := ""
	r := len(romanSlice) - 1
	for num > 0 {

		roman, ok := romanMap[num]
		if ok {
			res += roman
			return res
		}
		// near := getNearNum(num)
		// res += romanMap[near]
		// num -= near

		nearPos := getNearNumV2(num, 0, r)
		// 这里改为 除法，直接跳过 num的整数倍
		compant(num/romanSlice[nearPos], romanMap[romanSlice[nearPos]], &res)
		// 然后获取 当前的num
		num %= romanSlice[nearPos]
		// 这里缩小 romanSlice的边界，由于 num已经缩小了 所以二分再从 len(romanSlice)开始就太low了 这里可以直接跳过
		r = nearPos

	}

	return res
}
func getNearNumV2(num, l, r int) int {

	// 这里的等于是必须的 若没有等于 会出现取994的时候 会取到 1000
	// 但是这里有了等于 在判断到下标为12的时候 就会继续循环，发现1000比994大 然后r=mid-1
	for l <= r {
		mid := (l + r) >> 1

		if romanSlice[mid] > num {
			r = mid - 1
		} else if romanSlice[mid] < num {
			l = mid + 1
		} else {
			return mid
		}

	}
	return r
}

func getNearNum(num int) int {

	for k, v := range romanSlice {
		if num < v {
			if k > 0 {
				return romanSlice[k-1]
			} else {
				return romanSlice[0]
			}
		}
	}

	return romanSlice[len(romanSlice)-1]
}
