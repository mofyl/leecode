package V2

import (
	"fmt"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	res := generateTrees(3)

	for i := 0; i < len(res); i++ {
		v := inorderTraversal(res[i])
		fmt.Println(v)
	}
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return generateTreesHelper(1, n)
}

func generateTreesHelper(start, end int) []*TreeNode {

	if start > end {
		// 这里一定要给一个 空对象，保证 返回的[]*TreeNode 长度不为0
		// 防止 left || right 没有值的时候 无法构成树
		return []*TreeNode{nil}
	}
	allTree := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {

		left := generateTreesHelper(start, i-1)
		right := generateTreesHelper(i+1, end)

		for _, leftTree := range left {
			for _, rightTree := range right {
				allTree = append(allTree,
					&TreeNode{Val: i, Left: leftTree, Right: rightTree})
			}

		}

	}
	return allTree
}
