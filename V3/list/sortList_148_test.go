package list

import (
	"encoding/asn1"
	"github.com/xwb1989/sqlparser"
	"testing"
)

func TestSortList(t *testing.T) {
	head := &ListNode{}
	AddNode(head, 4)
	AddNode(head, 2)
	AddNode(head, 1)
	AddNode(head, 3)

	res := sortList(head.Next)

	PrintNode(res)
}

// 该算法要求使用 O( nlog(n) )
// O( nlog(n) ) 排序有 归并 堆排序 快速排序
// 这里要求使用 常数空间 所以不能使用 递归 快排就排除了。
// 这里是 递归版的 归并 是 由顶到底 的方式 , 还可以由 底 到 顶
func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	// 寻找 当前的中点
	// 最终 slow 会指向右中点
	fast, slow := head.Next, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	tmp := slow.Next
	// 这里一定要将 链表从 slow 处断开
	slow.Next = nil
	left := sortList(head)
	right := sortList(tmp)

	// 合并 将 left 和 right 合并为有序链表
	newHead := &ListNode{}
	p := newHead
	for left != nil && right != nil {
		if left.Val <= right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
	}

	var afterNode *ListNode

	if left != nil {
		afterNode = left
	} else if right != nil {
		afterNode = right
	}

	p.Next = afterNode

	return newHead.Next
}

func sortList_V2(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	invt := 1

	p := head
	length := 0

	for p != nil {
		p = p.Next
		length++
	}

	for invt < length {

		count := 1
		start := head
		tail := head

		for count <= invt {
			tail
			count++
		}

	}

}
