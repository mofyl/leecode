/*
 * @Author: lirui38
 * @Date: 2021-12-20 10:57:17
 * @LastEditTime: 2021-12-25 16:09:15
 * @LastEditors: lirui38
 * @Description:
 * @FilePath: \leecode\V3\tree\V3\113_pathSum_test.go
 */

package V3

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestPathSum(t *testing.T) {

	// l51 := &TreeNode{Val: 5}
	// l41 := &TreeNode{Val: 4}
	// l8 := &TreeNode{Val: 8}
	// l11 := &TreeNode{Val: 11}
	// l13 := &TreeNode{Val: 13}
	// l42 := &TreeNode{Val: 4}
	// l7 := &TreeNode{Val: 7}
	// l2 := &TreeNode{Val: 2}
	// l52 := &TreeNode{Val: 5}
	// l1 := &TreeNode{Val: 1}

	// l51.Left = l41
	// l51.Right = l8

	// l41.Left = l11

	// l8.Left = l13
	// l8.Right = l42

	// l11.Left = l7
	// l11.Right = l2

	// l42.Left = l52
	// l42.Right = l1

	// res := pathSum(l51, 22)

	// fmt.Println(res)

	// l1 := &TreeNode{Val: 1}
	// l2 := &TreeNode{Val: 2}
	// l3 := &TreeNode{Val: 3}

	// l1.Left = l2
	// l1.Right = l3

	// res := pathSum(l1, 22)
	// fmt.Println(res)

	c := make([]int, 0, 5)
	f1(c)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	sh.Len = 10

	fmt.Println(c)

	// fmt.Println(c)
}

func f1(cur []int) {
	cur = append(cur, 1)
	fmt.Println(cap(cur))

	cur = append(cur, 2)
	fmt.Println(cap(cur))

	cur = append(cur, 3)
	fmt.Println(cap(cur))

	// cur = append(cur, 4)
	// fmt.Println(cap(cur))

	f2(cur)

	fmt.Println(cur)
}

func f2(cur []int) {
	cur = append(cur, 5)

	fmt.Println("f2  ", cur)
}

func pathSum(root *TreeNode, targetSum int) [][]int {

	res := make([][]int, 0)
	pathSumHelper(root, targetSum, []int{}, &res)
	return res
}

func pathSumHelper(root *TreeNode, targetSum int, cur []int, res *[][]int) {

	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		cur = append(cur, root.Val)
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	sum := targetSum - root.Val
	cur = append(cur, root.Val)

	pathSumHelper(root.Left, sum, cur, res)
	pathSumHelper(root.Right, sum, cur, res)

}
