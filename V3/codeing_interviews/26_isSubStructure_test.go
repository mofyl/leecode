package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestIsSubStructure(t *testing.T) {

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
	l41 := &TreeNode{Val: 4}

	l41.Left = l11

	res := isSubStructure(l3, l41)

	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//
	//l1.Left = l2
	//l1.Right = l3
	//
	//l31 := &TreeNode{Val: 3}
	//l11 := &TreeNode{Val: 1}
	//
	//l31.Left = l11
	//
	//res := isSubStructure(l1, l31)

	fmt.Println(res)
}

// 遍历 A 中的每个节点
func isSubStructure(A *TreeNode, B *TreeNode) bool {

	// 若 A == nil || B == nil 则直接返回 false 因为若B
	/*
		1.A 为空 时 B无论是否为空 都要返回false
			1) A 为空 B 为空 ，题目规定 空树不是 字结构 返回false
			2) A 为空 B 不为空，不是字结构 返回false
	*/
	if A == nil || B == nil {
		return false
	}

	// B 是否是 A的子结构
	// 以A 的左子树 开始遍历， 判断 B 是否是 子结构
	// 以A 的右子树 开始遍历，判断 B 是否是 子结构
	return isSubStructureHelper(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// 判断 当前B 是否是 当前A 的子结构
func isSubStructureHelper(A *TreeNode, B *TreeNode) bool {

	// 表示 B 已经匹配完成了
	if B == nil {
		return true
	}
	// B 不为空 A 为空 匹配失败
	if A == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	}

	return isSubStructureHelper(A.Left, B.Left) && isSubStructureHelper(A.Right, B.Right)

}
