package V1

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPathSum(t *testing.T) {
	//
	//n1 := &TreeNode{Val: 5}
	//n2 := &TreeNode{Val: 4}
	//n3 := &TreeNode{Val: 8}
	//n4 := &TreeNode{Val: 11}
	//n5 := &TreeNode{Val: 13}
	//n6 := &TreeNode{Val: 4}
	//n7 := &TreeNode{Val: 7}
	//n8 := &TreeNode{Val: 2}
	//n9 := &TreeNode{Val: 5}
	//n10 := &TreeNode{Val: 1}
	//
	//n1.Left = n2
	//n1.Right = n3
	//
	//n2.Left = n4
	//
	//n4.Left = n7
	//n4.Right = n8
	//
	//n3.Left = n5
	//n3.Right = n6
	//
	//n5.Left = n9
	//n5.Right = n10
	//
	//res := pathSum(n1, 22)
	////
	//fmt.Println(res)
	//
	//res := levelOrder(n1)
	//
	//fmt.Println(res)
	//c := make([]int, 0, 2)
	//TempFunc(c)
	//fmt.Println(c)
	//size := unsafe.Sizeof(0)
	//head := unsafe.Pointer(&c)
	//p := (*[2]int)(unsafe.Pointer(*(*uintptr)(head)))
	//
	//fmt.Println(head)
	//fmt.Println(p)
	//fmt.Println(*(*int)(unsafe.Pointer(uintptr(head) + size)))
	//

	m := make(map[string]struct{})
	m["aaa"] = struct{}{}
	fmt.Printf("main %v\n", unsafe.Pointer(&m))
	TempMap(m)
}

func TempFunc(c []int) {
	fmt.Printf("tempFunc %v", unsafe.Pointer(&c))
	c = append(c, 3, 2, 1, 4)
}

func TempMap(m map[string]struct{}) {

	fmt.Printf("tempFunc %v\n", unsafe.Pointer(&m))
}

var res = 0

func pathSum_V1(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}

	pathSum_437_V1(root, sum)
	if root.Left != nil {
		pathSum_V1(root.Left, sum)
	}
	if root.Right != nil {
		pathSum_V1(root.Right, sum)
	}

	return res
}

func pathSum_437_V1(root *TreeNode, sum int) {

	if root == nil {
		return
	}

	if sum == root.Val {
		res += 1
	}

	pathSum_437_V1(root.Left, sum-root.Val)
	pathSum_437_V1(root.Right, sum-root.Val)

}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	tmp := make([]int, 0)
	return pathSum_437(root, tmp, sum)

}

func pathSum_437(root *TreeNode, sumList []int, sum int) int {

	if root == nil {
		return 0
	}
	tmp := make([]int, len(sumList))
	for i := 0; i < len(sumList); i++ {
		tmp[i] = sumList[i]
		tmp[i] += root.Val
	}
	count := 0
	tmp = append(tmp, root.Val)
	fmt.Println(tmp, root.Val)
	for _, v := range tmp {
		if v == sum {
			count += 1
		}
	}

	return count + pathSum_437(root.Left, tmp, sum) + pathSum_437(root.Right, tmp, sum)
}
