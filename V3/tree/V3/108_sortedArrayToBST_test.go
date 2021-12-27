/*
 * @Author: lirui38
 * @Date: 2021-12-17 16:33:37
 * @LastEditTime: 2021-12-17 17:20:39
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: /leecode/V3/tree/V3/108_sortedArrayToBST_test.go
 */

package V3

import (
	"fmt"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {

	// nums := []int{-10, -3, 0, 5, 9}

	// root := sortedArrayToBST(nums)

	// res := inorderTraversal(root)

	// fmt.Println(res)

	nums := []int{1, 3}

	root := sortedArrayToBST(nums)

	res := inorderTraversal(root)

	fmt.Println(res)

}

func sortedArrayToBST(nums []int) *TreeNode {

	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}

// 这里 left 和 right 都是下标，从0开始
func sortedArrayToBSTHelper(nums []int, left, right int) *TreeNode {

	if left > right {
		return nil
	}

	// 若 left+right 为 偶数 则直接/2
	// 若为奇数 则  / 2 +1
	rootIdx := (left + right)

	if rootIdx&1 == 0 {
		rootIdx = rootIdx / 2
	} else {
		rootIdx = rootIdx/2 + 1
	}

	root := &TreeNode{Val: nums[rootIdx]}

	root.Left = sortedArrayToBSTHelper(nums, left, rootIdx-1)
	root.Right = sortedArrayToBSTHelper(nums, rootIdx+1, right)

	return root
}
