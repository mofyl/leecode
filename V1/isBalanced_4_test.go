package V1

import (
	"fmt"
	"testing"
)

/*
	通过检查每个节点距离根节点的最大高度 是否符合来判断整个数是否符合
		因为若树中有某个节点不符合的话，那么整颗树一定不符合


*/

func TestIsBalance(t *testing.T) {
	n := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	fmt.Println(isBalanced(n))
}

var isBalance = true

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	bfs(root)
	return isBalance
}

func bfs(root *TreeNode) int {

	if root == nil || !isBalance {
		// 这里返回0: 若当前节点为nil，那么对于该节点的父节点来说
		// 当前儿子节点的高度就是0
		return 0
	}

	l := bfs(root.Left)
	r := bfs(root.Right)

	if abs_isBalanced(l-r) > 1 {
		isBalance = false
	}

	return max_isBalanced(l, r) + 1 // 这里由于是 返回上一层，所以层数应该+1

}

func max_isBalanced(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs_isBalanced(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
