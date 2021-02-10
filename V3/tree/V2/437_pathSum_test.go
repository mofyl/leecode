package V2

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

func pathSum_437(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}

	prefix := make(map[int]int)
	prefix[0] = 1
	res := 0

	pathSum_437Helper(root, prefix, 0, &res, sum)

	return res
}

func pathSum_437Helper(root *TreeNode, prefix map[int]int, curPrefix int, res *int, target int) {
	if root == nil {
		return
	}

	curPrefix += root.Val

	*res += prefix[curPrefix-target]

	prefix[curPrefix] = prefix[curPrefix] + 1

	pathSum_437Helper(root.Left, prefix, curPrefix, res, target)
	pathSum_437Helper(root.Right, prefix, curPrefix, res, target)

	prefix[curPrefix] = prefix[curPrefix] - 1
}
