package V2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestFindTilt(t *testing.T) {

	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//
	//l1.Left = l2
	//l1.Right = l3
	//
	//res := findTilt(l1)
	//fmt.Println(res)

	l4 := &TreeNode{Val: 4}
	l2 := &TreeNode{Val: 2}
	l9 := &TreeNode{Val: 9}
	l3 := &TreeNode{Val: 3}
	l5 := &TreeNode{Val: 5}
	l7 := &TreeNode{Val: 7}

	l4.Left = l2
	l4.Right = l9

	l2.Left = l3
	l2.Right = l5

	l9.Right = l7

	res := findTilt(l4)
	fmt.Println(res)
}

func findTilt(root *TreeNode) int {

	if root == nil {
		return 0
	}

	res := 0
	findTiltHelper(root, &res)

	return res
}

// 返回的是子树的和
func findTiltHelper(root *TreeNode, res *int) int {

	if root == nil {
		return 0
	}

	leftSum := findTiltHelper(root.Left, res)
	rightSum := findTiltHelper(root.Right, res)

	*res += tools.Abs(leftSum - rightSum)

	return leftSum + rightSum + root.Val

}
