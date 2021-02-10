package V2

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l5 := &TreeNode{Val: 5}

	l1.Left = l2
	l1.Right = l3

	l2.Right = l5

	res := binaryTreePaths(l1)
	fmt.Println(res)
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := make([]string, 0)
	binaryTreePathsHelper(root, "", &res)

	return res
}

func binaryTreePathsHelper(root *TreeNode, str string, res *[]string) {

	if root == nil {
		return
	}

	str += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		*res = append(*res, str)
		return
	}

	binaryTreePathsHelper(root.Left, str+"->", res)
	binaryTreePathsHelper(root.Right, str+"->", res)
}
