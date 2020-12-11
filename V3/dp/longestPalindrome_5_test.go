package dp

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	str := "abba"

	res := longestPalindrome(str)
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	maxLen := 0
	start := 0
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			// 这里 +1 是因为 在最后取 区间的时候 右边是开区间，
			// 若计算的出的 start=0 和 maxLen=3 那么3这个元素是取不到的。所以这里要加1
			// +1 不会影响最后的结果 因为 下面 maxLen 记录的时候也是+1之后的值
			// 这里判断 j-i+1 > maxLen 是为了减少计算次数。 若当前长度还没有大于 上次记录的长度 那么也没有必要取判断
			if j-i+1 > maxLen && validPalindromic(s, i, j) {
				start = i
				maxLen = j - i + 1
			}

		}
	}

	return s[start : start+maxLen]
}

func validPalindromic(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func longestPalindrome_dp(s string) string {
	if len(s) <= 1 {
		return s
	}

	start := 0
	maxLen := 0
	dp := make([][]bool, len(s))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	// 若 i , j 重合 都指向了同一个元素 那么该位置一定是true
	// 这里提前处理好
	for i := 0; i < len(dp); i++ {
		dp[i][i] = true
	}
	// 这里 i 是 边界指针 也就是right
	for i := 1; i < len(s); i++ {
		// j是 起始指针
		// 所有后面判断的时候 都是 i在右边
		for j := 0; j < i; j++ {
			// 这里就是不相等的情况
			if s[i] != s[j] {
				dp[j][i] = false
			} else {
				/*
					这里是判断 i，j 之间距离太小
					因为后面要使用 j+1 和 i-1 这两个位置。这里怕指到错误的位置去 所以做的检查
					这里使用距离来判断， 若 i-1 和 j+1 离得太近 那么就一定不对
						只要 距离小于1 那么就认为是检查完了。检查完了就直接给dp[j][i]赋值就好
					i-1-(j+1) < 1 => i - j < 3
				*/
				if i-j < 3 {
					dp[j][i] = true
				} else {
					// 这里就是没有处理完 那么当前的结果就要依赖 j+1 i-1 位置的结果了
					dp[j][i] = dp[j+1][i-1]
				}
			}
			// 这里用 i-j+1 也是因为 区间的问题 是左闭右开的 所以这里+1多给了一个
			if dp[j][i] && i-j+1 > maxLen {
				start = j
				maxLen = i - j + 1
			}
		}
	}

	return s[start : start+maxLen]
}
