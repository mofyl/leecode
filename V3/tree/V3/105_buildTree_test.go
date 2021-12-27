/*
 * @Author: lirui38
 * @Date: 2021-12-24 10:11:21
 * @LastEditTime: 2021-12-24 10:52:16
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/105_buildTree_test.go
 */

package V3

import (
	"math"
	"testing"
)

func TestBuildTree_105(t *testing.T) {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	preIdx := 0
	inIdx := 0

	return buildTreeHelper(preorder, inorder, &preIdx, &inIdx, math.MaxInt)
}

func buildTreeHelper(preorder []int, inorder []int, preIdx *int, inIdx *int, stop int) *TreeNode {

	if *preIdx == len(preorder) {
		return nil
	}

	if inorder[*inIdx] == stop {
		*inIdx++
		return nil
	}

	val := preorder[*preIdx]
	root := &TreeNode{Val: val}

	*preIdx++
	root.Left = buildTreeHelper(preorder, inorder, preIdx, inIdx, val)
	root.Right = buildTreeHelper(preorder, inorder, preIdx, inIdx, stop)
	return root
}
