package codeing_interviews

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中叙遍历
func print(node *TreeNode) {

	if node == nil {
		return
	}
	print(node.Left)
	fmt.Println(node.Val)
	print(node.Right)
}
