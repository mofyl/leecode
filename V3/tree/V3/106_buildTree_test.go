/*
 * @Author: lirui38
 * @Date: 2021-12-24 10:40:42
 * @LastEditTime: 2021-12-24 11:02:05
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/106_buildTree_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestBuildTree_106(t *testing.T) {

	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	root := buildTree_106(inorder, postorder)

	res := inorderTraversal(root)

	fmt.Println(res)
}

func buildTree_106(inorder []int, postorder []int) *TreeNode {

	inIdx := len(inorder) - 1
	postIdx := len(postorder) - 1
	stop := 65535 // 随便给个一个大数字，最好可以用 math.MaxInt. 但是leecode 有时候会报 math库 不存在

	return buildTree106Helper(inorder, postorder, &inIdx, &postIdx, stop)

}

func buildTree106Helper(inorder []int, postorder []int, inIdx *int, postIdx *int, stop int) *TreeNode {

	if *postIdx < 0 {
		return nil
	}

	if inorder[*inIdx] == stop {
		*inIdx--
		return nil
	}

	val := postorder[*postIdx]
	root := &TreeNode{Val: val}

	*postIdx--

	root.Right = buildTree106Helper(inorder, postorder, inIdx, postIdx, val)
	root.Left = buildTree106Helper(inorder, postorder, inIdx, postIdx, stop)

	return root
}
