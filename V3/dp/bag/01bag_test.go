package bag

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func Test01Bag(t *testing.T) {

	vol := []int{1, 2, 3, 4, 5}
	value := []int{5, 4, 3, 2, 1}
	W := 10

	res := bag_01(vol, value, W)

	fmt.Println(res)
}

// 求出
func bag_01(vol []int, values []int, W int) int {

	if W == 0 {
		return 0
	}

	//return bag_01Helper(vol, values, W, 0, 0)
	return bag_01DPWay(vol, values, W)
}

// 递归的定义为: 从idx 位置到结束位置 在剩余重量为 curW 的情况下 返回可以获得的最大价值
func bag_01Helper(vol []int, values []int, W int, idx int, curW int) int {

	// 这里表示 若当前重量已经达到上限 当前是无法获取任何价值的
	// 因为或要获取价值就必须去减去重量，所以无法获取价值
	if curW >= W {
		return 0
	}
	// 这里价值数组 已经来到了末尾，还是没有填满背包，
	// 这里由于价值数组已经到头了，所以无法获取价值
	if idx == len(values) {
		return 0
	}

	// 对于每个重量来说 我们当前可以进行选择，也可以不选择

	// 不选择 idx 位置的物品
	noOp := bag_01Helper(vol, values, W, idx+1, curW)
	// 选择 idx 位置的物品
	choice := values[idx] + bag_01Helper(vol, values, W, idx+1, curW+vol[idx])

	// 返回最大的价值
	return tools.Max(noOp, choice)

}

func bag_01DPWay(vol []int, values []int, W int) int {

	// 这里可以取到  len(val) 的位置 所以给 len(val)+1
	dp := make([][]int, len(values)+1)

	for i := 0; i < len(dp); i++ {
		// 这里给W+1 应为可以取到W
		dp[i] = make([]int, W+1)
	}
	// 这里其实不需要的 应为初始化后 默认就是0
	//for i := 0; i < W+1; i++ {
	//	dp[len(dp)-1][i] = 0
	//}
	//

	// 最后一行已经填好了，都是0 , 所以从倒数第二行开始
	for i := len(dp) - 2; i >= 0; i-- {

		for j := 0; j <= W; j++ {

			// 这里表示超重了， 所以价值为0
			if j+vol[i] > W {
				dp[i][j] = 0
			} else {
				dp[i][j] = tools.Max(dp[i+1][j], values[i]+dp[i+1][j+vol[i]])
			}

		}

	}

	return dp[0][0]
}
