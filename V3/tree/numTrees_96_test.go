package tree

import (
	"fmt"
	"testing"
)

func TestNumTrees(t *testing.T) {
	n := 3

	res := numTrees(n)

	fmt.Println(res)
	//
	//nums := []int{1, 2, 3}
	//
	//newNums := append(nums[:0], nums[1:]...)
	//
	//fmt.Println(nums)
	//fmt.Println(newNums)

}

func numTrees(n int) int {

	if n < 1 {
		return 0
	}
	var res int

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	for i := 0; i < len(nums); i++ {
		root := &TreeNode{Val: nums[i]}

		newNum := append([]int{}, nums[:i]...)
		newNum = append(newNum, nums[i+1:]...)
		numTreesLoop(newNum, root, root, &res)
	}
	return res
}

func numTreesLoop(nums []int, root *TreeNode, curNode *TreeNode, res *int) {

	if len(nums) == 0 {
		*res++
		node := inorderTraversal(root)
		fmt.Println(node)
		return
	}

	for i := 0; i < len(nums); i++ {
		tmp := &TreeNode{
			Val: nums[i],
		}
		newNum := append([]int{}, nums[:i]...)
		newNum = append(newNum, nums[i+1:]...)
		if tmp.Val > root.Val {
			curNode.Right = tmp
			numTreesLoop(newNum, root, curNode.Right, res)
			curNode.Right = nil
		} else {
			curNode.Left = tmp
			numTreesLoop(newNum, root, curNode.Left, res)
			curNode.Left = nil
		}

	}

}
