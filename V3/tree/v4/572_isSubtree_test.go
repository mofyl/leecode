package v4

import (
	"fmt"
	"testing"
)

func TestIsSubtree(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}

	l3.Left = l4
	l3.Right = l5

	l4.Left = l1
	l4.Right = l2

	l11 := &TreeNode{Val: 1}
	l21 := &TreeNode{Val: 2}
	l41 := &TreeNode{Val: 4}

	l41.Left = l11
	l41.Right = l21

	res := isSubtree(l3, l41)

	fmt.Println(res)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	// 这里 空子树 也是子树
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	//if root.Val != subRoot.Val {
	//	return false
	//}
	// 判断 subRoot 是否为 root 的子节点 || subRoot 是否是root.Left 的子节点 || subRoot 是否是 root.Right 的子节点
	return isSubtreeHelper(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}

func isSubtreeHelper(root *TreeNode, subRoot *TreeNode) bool {

	if subRoot == nil && root == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return isSubtreeHelper(root.Left, subRoot.Left) && isSubtreeHelper(root.Right, subRoot.Right)

}
