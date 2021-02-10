package V2

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestMinDepth(t *testing.T) {
	l3 := &TreeNode{Val: 3}
	l9 := &TreeNode{Val: 9}
	l20 := &TreeNode{Val: 20}
	l15 := &TreeNode{Val: 15}
	l7 := &TreeNode{Val: 7}

	l3.Left = l9
	l3.Right = l20

	l20.Left = l15
	l20.Right = l7

	res := minDepth(l3)

	//l2 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//l4 := &TreeNode{Val: 4}
	//l5 := &TreeNode{Val: 5}
	//l6 := &TreeNode{Val: 6}
	//
	//l2.Right = l3
	//l3.Right = l4
	//l4.Right = l5
	//l5.Right = l6
	//
	//res := minDepth(l2)

	fmt.Println(res)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// dfs 解法
	//minVal := math.MaxInt64
	//
	//minDepthDFSHelper(root, 0, &minVal)

	// dfs 解法 V2
	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	// 这里表示 当前节点的 左子树 或者 右子树 为nil
	// 那么此时 我们应该返回 不为nil 的子树的 高度
	// 这里 left 或者 right 其中一个一定为0，
	// 那么这时返回不为 nil的子树的高度 就可以写为  left+right+1
	if root.Left == nil || root.Right == nil {
		return left + right + 1
	}
	return tools.Min(left, right) + 1

	// bfs 解法
	//que := make([]*TreeNode, 0)
	//que = append(que, root)
	//depths := make([]int, 0)
	//depths = append(depths, 1)
	//
	//for i := 0; i < len(que); i++ {
	//	depth := depths[i]
	//	node := que[i]
	//	if node.Left == nil && node.Right == nil {
	//		return depth
	//	}
	//
	//	if node.Left != nil {
	//		que = append(que, node.Left)
	//		depths = append(depths, depth+1)
	//	}
	//	if node.Right != nil {
	//		que = append(que, node.Right)
	//		depths = append(depths, depth+1)
	//	}
	//}
	//
	//return depths[len(depths)-1]
}

// dfs解法
func minDepthDFSHelper(root *TreeNode, curVal int, minVal *int) {
	if root == nil {
		return
	}

	// 这里表示是 叶子节点
	if root.Left == nil && root.Right == nil {
		tmp := curVal + 1
		if *minVal > tmp {
			*minVal = tmp
		}
		return
	}
	minDepthDFSHelper(root.Left, curVal+1, minVal)

	minDepthDFSHelper(root.Right, curVal+1, minVal)

}
