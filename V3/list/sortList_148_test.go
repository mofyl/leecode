package list

import (
	"leecode/tools"
	"testing"
)

func TestSortList(t *testing.T) {
	head := &tools.ListNode{}
	tools.AddNode(head, 4)
	tools.AddNode(head, 2)
	tools.AddNode(head, 1)
	tools.AddNode(head, 3)

	//res := sortList(head.Next)
	res := sortList_V2(head.Next)

	tools.PrintNode(res)
}

// 该算法要求使用 O( nlog(n) )
// O( nlog(n) ) 排序有 归并 堆排序 快速排序
// 这里要求使用 常数空间 所以不能使用 递归 快排就排除了。
// 这里是 递归版的 归并 是 由顶到底 的方式 , 还可以由 底 到 顶
func sortList(head *tools.ListNode) *tools.ListNode {

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
	newHead := &tools.ListNode{}
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

	var afterNode *tools.ListNode

	if left != nil {
		afterNode = left
	} else if right != nil {
		afterNode = right
	}

	p.Next = afterNode

	return newHead.Next
}

func sortList_V2(head *tools.ListNode) *tools.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	p := head
	length := 0

	for p != nil {
		p = p.Next
		length++
	}

	// 由于要改第一个，所以增加一个newHead
	newHead := &tools.ListNode{}
	newHead.Next = head
	// 这里将head分割为 1,2,4,6,8
	for invt := 1; invt < length; invt *= 2 {

		prev := newHead
		cur := newHead.Next

		for cur != nil {
			h1 := cur
			h2 := getHead(cur, invt)
			cur = getHead(h2, invt)

			// 对 h1,h2 进行排序
			mergeHead := merge_148(h1, h2)
			prev.Next = mergeHead

			// 移动prev 到 排序好 部分的末尾 由于 h1,h2之间都已经断开了，所以只需要判断prev.next != nil 就好了
			for prev.Next != nil {
				prev = prev.Next
			}
		}

	}
	return newHead.Next
}

func merge_148(h1, h2 *tools.ListNode) *tools.ListNode {

	newHead := &tools.ListNode{}
	p := newHead
	for h1 != nil && h2 != nil {

		if h1.Val < h2.Val {
			p.Next = h1
			h1 = h1.Next
		} else {
			p.Next = h2
			h2 = h2.Next
		}
		p = p.Next
	}
	if h1 != nil {
		p.Next = h1
	}
	if h2 != nil {
		p.Next = h2
	}
	return newHead.Next
}

// 返回 start 后 invt步的节点，并将 这两部分 分割
func getHead(start *tools.ListNode, invt int) *tools.ListNode {
	if start == nil {
		return start
	}
	tail := start
	// 由于这里要将 invt步后的节点和 start的末尾分割 所以要从1开始 让tail 指向 返回节点的前一个
	for i := 1; i < invt && tail.Next != nil; i++ {
		tail = tail.Next
	}

	next := tail.Next
	tail.Next = nil

	return next
}
