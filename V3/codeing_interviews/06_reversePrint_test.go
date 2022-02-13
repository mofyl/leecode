package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestReversePrint(t *testing.T) {

	l1 := &ListNode{Val: 1}
	l3 := &ListNode{Val: 3}
	l2 := &ListNode{Val: 2}

	l1.Next = l3
	l3.Next = l2

	//res := reversePrint(l1)

	res := reversePrint_v2(l1)

	fmt.Println(res)
}

func reversePrint(head *ListNode) []int {

	p := head

	res := make([]int, 0)

	for p != nil {

		res = append([]int{p.Val}, res...)
		p = p.Next
	}

	return res

}

func reversePrint_v2(head *ListNode) []int {

	if head == nil {
		return nil
	}

	var arr []int

	if head.Next != nil {
		// 先将 后面的加入数组， 然后 在将本次的加入 ，就成了倒序了
		arr = reversePrint_v2(head.Next)
	}

	arr = append(arr, head.Val)

	return arr
}
