/*
 * @Author: lirui38
 * @Date: 2021-12-24 11:10:40
 * @LastEditTime: 2021-12-24 15:18:47
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/671_findSecondMinimumValue_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestFindSecondMinimumValue(t *testing.T) {

	l21 := &TreeNode{Val: 2}
	l22 := &TreeNode{Val: 2}
	l51 := &TreeNode{Val: 5}
	l52 := &TreeNode{Val: 5}
	l7 := &TreeNode{Val: 7}

	l51.Left = l52
	l51.Right = l7

	l21.Left = l22
	l21.Right = l51

	res := findSecondMinimumValue(l21)

	fmt.Println(res)

	// l21 := &TreeNode{Val: 2}
	// l22 := &TreeNode{Val: 2}
	// l23 := &TreeNode{Val: 2}

	// l21.Left = l22
	// l21.Right = l23

	// res := findSecondMinimumValue(l21)

	// fmt.Println(res)

}

func findSecondMinimumValue(root *TreeNode) int {
	ans := -1
	findSecondMinimumValueHelper(root, root.Val, &ans)

	return ans
}

func findSecondMinimumValueHelper(root *TreeNode, minValue int, ans *int) {
	if root == nil {
		return
	}

	// 增加剪枝
	if *ans != -1 && root.Val > *ans {
		return
	}

	if root.Val > minValue {
		*ans = root.Val
	}

	findSecondMinimumValueHelper(root.Left, minValue, ans)
	findSecondMinimumValueHelper(root.Right, minValue, ans)

}
