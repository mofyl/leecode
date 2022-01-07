/*
 * @Author: lirui
 * @Date: 2022-01-05 13:33:10
 * @LastEditTime: 2022-01-07 14:27:41
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\tree\V3\437_pathSum_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestPathSum_437(t *testing.T) {

	l5 := &TreeNode{Val: 5}
	l51 := &TreeNode{Val: 5}

	l4 := &TreeNode{Val: 4}
	l41 := &TreeNode{Val: 4}

	l8 := &TreeNode{Val: 8}
	l11 := &TreeNode{Val: 11}
	l13 := &TreeNode{Val: 13}
	l7 := &TreeNode{Val: 7}
	l2 := &TreeNode{Val: 2}
	l1 := &TreeNode{Val: 1}

	l5.Left = l4
	l5.Right = l8

	l4.Left = l11

	l8.Left = l13
	l8.Right = l41

	l11.Left = l7
	l11.Right = l2

	l41.Left = l51
	l41.Right = l1

	res := pathSum_437(l5, 22)

	fmt.Println(res)
}

func pathSum_437(root *TreeNode, targetSum int) int {

	res := 0

	data := make(map[int]int)
	// 我们每次都用 curSum - targetSum , 若一上来的 curSum == targetSum 那么这条路径我们就统计不到了
	// 所以提前加上 0 这条路径
	data[0] = 1

	pathSum_437Helper(root, data, 0, &res, targetSum)

	return res
}

func pathSum_437Helper(root *TreeNode, data map[int]int, curSum int, res *int, tS int) {

	if root == nil {
		return
	}

	curSum += root.Val

	*res += data[curSum-tS]

	data[curSum] = data[curSum] + 1

	pathSum_437Helper(root.Left, data, curSum, res, tS)
	pathSum_437Helper(root.Right, data, curSum, res, tS)

	data[curSum] = data[curSum] - 1

}
