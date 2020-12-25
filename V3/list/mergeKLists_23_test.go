package list

import "testing"

func TestMergeKLists(t *testing.T) {

	lists := make([]*ListNode, 0, 10)

	l1 := &ListNode{}
	AddNode(l1, 1)
	AddNode(l1, 4)
	AddNode(l1, 5)
	//
	l2 := &ListNode{}
	AddNode(l2, 1)
	AddNode(l2, 3)
	AddNode(l2, 4)

	l3 := &ListNode{}
	AddNode(l3, 2)
	AddNode(l3, 6)

	lists = append(lists, l1.Next, l2.Next, l3.Next)

	//lists = append(lists, l1.Next, nil, l3.Next)

	//res := mergeKLists(lists)
	res := mergeKListsV2(lists)

	PrintNode(res)

}

func mergeKLists(lists []*ListNode) *ListNode {
	// 先合并两个 剩余的都和之前的做比较
	if len(lists) < 1 {
		return nil
	}

	if len(lists) < 2 {
		return lists[0]
	}

	res := &ListNode{}

	tmp := mergeTwoLists(lists[0], lists[1])
	if tmp == nil {
		return nil
	}
	res.Next = tmp

	for i := 2; i < len(lists); i++ {
		if lists[i] != nil {
			tmp = mergeTwoLists(res.Next, lists[i])
			res.Next = tmp
		}
	}
	return res.Next
}

func mergeKListsV2(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(list []*ListNode, l, r int) *ListNode {
	if l == r {
		return list[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	//l1 := merge(list, l, mid)
	//l2 := merge(list, mid+1, r)
	//return mergeTwoLists(l1 , l2)
	return mergeTwoLists(merge(list, l, mid), merge(list, mid+1, r))
}
