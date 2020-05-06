package main

import (
	"bytes"
	"fmt"
	"math"
)

func main5() {
	str := "ac"

	fmt.Println(longestPalindrome(str))
}

func manacherStr(str string) string {
	res := bytes.NewBuffer([]byte{})

	// res.WriteString("$")
	for i := 0; i < len(str); i++ {
		res.WriteString("#")
		res.WriteByte(str[i])
	}
	res.WriteString("#")
	return res.String()

}

func longestPalindrome(str string) string {

	if len(str) == 1 || len(str) == 0 {
		return str
	}
	// 这里是为了去掉判断len(str)奇数还是偶数
	// 这里统一为str 添加#  这样以来len(newStr) 一定是奇数    偶数+奇数 一定会是奇数
	newStr := manacherStr(str)
	newStrLen := len(newStr)
	// p中保存的是 以每个newStr中字符 为中心的，最大回文半径
	p := make([]int, len(newStr))

	maxRight := 0

	pos := 0
	// 最大回文中心
	resCenter := 0
	// 最大回文半径
	resLen := 0

	for i := 0; i < newStrLen; i++ {

		if maxRight > i {
			//
			p[i] = int(math.Min(float64(p[2*pos-i]), float64(maxRight-i)))
		} else {
			// 这里其实 用1 和 0 出来的结果一样，但是用1 可以在求回文半径的时候 少一些循环
			/*

				结果一样的原因：
					当走这个分支的时候 是因为当前的回文半径 比当前元素位置小
					0 的时候： i-0 = i  i+0 = i 所以当p[i] = 0 的时候 一定会 走到p[i]++ 就会变成1
					1 的时候： 这里不满足条件就是1  满足就是2 其实跟上面一样  但是会少一些循环，直接被判断条件阻挡了
			*/
			p[i] = 1
		}
		// 求回文半径
		for i-p[i] > 0 && i+p[i] < newStrLen && newStr[i+p[i]] == newStr[i-p[i]] {
			p[i]++
		}
		// 当  当前回文半径 在当前元素的左边界之内的时候 maxRight 和 pos就会 增加
		if maxRight < i+p[i] {
			maxRight = i + p[i]
			pos = i
		}

		if resLen < p[i] {
			resCenter = i
			resLen = p[i]
		}

	}

	left := (resCenter - resLen + 1) / 2
	right := (resCenter + resLen) / 2

	return str[left:right]

}
