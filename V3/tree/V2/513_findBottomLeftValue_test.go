package V2

import (
	"fmt"
	"math"
	"testing"
)

func TestFindBottomLeftValue(t *testing.T) {

	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//
	//l2.Left = l1
	//l2.Right = l3
	//
	//v := findBottomLeftValue(l1)

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}
	l3 := &TreeNode{Val: 3}
	l4 := &TreeNode{Val: 4}
	l5 := &TreeNode{Val: 5}
	l6 := &TreeNode{Val: 6}
	l7 := &TreeNode{Val: 7}

	l1.Left = l2
	l1.Right = l3

	l2.Left = l4

	l3.Left = l5
	l3.Right = l6

	l5.Left = l7

	v := findBottomLeftValue(l1)
	fmt.Println(v)
}

func findBottomLeftValue(root *TreeNode) int {

	if root == nil {
		return 0
	}

	//que := make([]int, 0)
	//depth := 1
	//
	//findBottomLeftValueHelper(root, depth, &que)
	//
	//return que[len(que)-1]

	res := 0
	maxLevel := math.MinInt64

	findBottomLeftValueHelper_V2(root, &res, &maxLevel, 0)
	return res
}

// 这里将每一层的最左边的节点都记录下了，没有必要
func findBottomLeftValueHelper(root *TreeNode, depth int, que *[]int) {

	if root == nil {
		return
	}

	if depth > len(*que) {
		*que = append(*que, root.Val)
	}

	findBottomLeftValueHelper(root.Left, depth+1, que)

	findBottomLeftValueHelper(root.Right, depth+1, que)

}

/*
	既然要找 左下角的节点
	那么就可以一直递归左子树 当递归到叶子节点的时候 记录当前深度，和当前的 root.Val
	若后面还有深度比 当前大的，那么就继续更换值,这里没有等于 是因为 我们要找最左边的值，防止同层的节点值将最左边的值覆盖
*/
func findBottomLeftValueHelper_V2(root *TreeNode, res, maxLevel *int, curLevel int) {

	if root == nil {
		return
	}

	findBottomLeftValueHelper_V2(root.Left, res, maxLevel, curLevel+1)

	if curLevel > *maxLevel {
		*maxLevel = curLevel
		*res = root.Val
	}

	findBottomLeftValueHelper_V2(root.Right, res, maxLevel, curLevel+1)
}
