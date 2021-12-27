package V3

import "testing"

func TestGenerateTrees(t *testing.T) {

	n := 3
	res := generateTrees(n)

	t.Log(res)
}

func generateTrees(n int) []*TreeNode {

	return generateTreesHelper(1, n)
}

func generateTreesHelper(start, end int) []*TreeNode {

	if start > end {
		// 这里不直接返回nil 的原因是 无法组成树的话，空树也是一颗树
		return []*TreeNode{nil}
	}

	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {

		leftNodes := generateTreesHelper(start, i-1)
		rightNodes := generateTreesHelper(i+1, end)

		// 双层遍历是因为 总的结果 = left * right
		for _, leftNode := range leftNodes {
			for _, rightNode := range rightNodes {
				res = append(res, &TreeNode{Val: i, Left: leftNode, Right: rightNode})
			}
		}

	}

	return res

}
