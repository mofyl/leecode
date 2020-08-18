package V1

import (
	"fmt"
	"testing"
)

func TestLCS(t *testing.T) {

	// str1 := "abcde"
	// str2 := "ace"

	str1 := "abc"
	str2 := "def"

	fmt.Println(longestCommonSubsequence(str1, str2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	/*

		使用LCS动态规划，首先找到边界，然后找到状态改变的条件,然后去套 状态改变公式就好
		这里的LCS会使用一个二维数组来存储状态

	*/
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	twoArr := make([][]int, len(text1)+1)

	for i := 0; i < len(text1)+1; i++ {
		twoArr[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {

		for j := 1; j <= len(text2); j++ {

			if text1[i-1] == text2[j-1] {
				twoArr[i][j] = twoArr[i-1][j-1] + 1

			} else {

				twoArr[i][j] = max(twoArr[i-1][j], twoArr[i][j-1])
			}

		}

	}

	fmt.Println(twoArr[len(text1)][len(text2)])
	fmt.Println(getCLS(text1, text2, twoArr))

	return twoArr[len(text1)][len(text2)]

}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

func getCLS(text1, text2 string, bitmap [][]int) string {

	var i, j = 1, 1
	var num1, num2, num = 0, 0, 0
	len1 := len(text1)
	len2 := len(text2)
	res := ""

	for i <= len1 && j <= len2 {

		if text1[i-1] == text2[j-1] {
			res += string(text1[i-1])
			i++
			j++
		} else {
			num1 = 0
			num2 = 0

			if i < len1 {
				num1 = bitmap[i+1][j]
			}
			if j < len2 {
				num2 = bitmap[i][j+1]
			}
			num = max(num1, num2)
			if num == num1 {
				i++
			} else {
				j++
			}
		}

	}

	return res

}
