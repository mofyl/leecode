package list

import (
	"fmt"
	"leecode/tools"
	"testing"
)

func TestNextLargerNodes(t *testing.T) {
	head := &tools.ListNode{}
	tools.AddNode(head, 2)
	tools.AddNode(head, 1)
	tools.AddNode(head, 5)
	//AddNode(head, 3)
	//AddNode(head, 5)

	res := nextLargerNodes(head.Next)
	fmt.Println(res)
}

func nextLargerNodes(head *tools.ListNode) []int {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return []int{0}
	}

	/*
		找下一个比Val 大的数字，就是当前val的res。
	*/

	res := make([]int, 0)

	p := head

	for p != nil {
		huge := findNextHuge(p)
		res = append(res, huge)
		p = p.Next
	}
	return res
}

func findNextHuge(head *tools.ListNode) int {

	p := head.Next
	val := head.Val

	for p != nil {
		if p.Val > val {
			return p.Val
		}
		p = p.Next
	}

	return 0
}

func nextLargerNodes_v2(head *tools.ListNode) []int {
	if head == nil {
		return nil
	}
	data := []int{}
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	stack := []int{}
	res := make([]int, len(data))
	for i, v := range data {
		for len(stack) > 0 && data[stack[len(stack)-1]] < v {
			res[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
