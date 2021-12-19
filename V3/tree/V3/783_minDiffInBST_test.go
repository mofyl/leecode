package V3

// func TestMinDiffInBST(t *testing.T) {

// 	l1 := &TreeNode{Val: 1}
// 	l2 := &TreeNode{Val: 2}
// 	l3 := &TreeNode{Val: 3}
// 	l4 := &TreeNode{Val: 4}
// 	l6 := &TreeNode{Val: 6}

// 	l4.Left = l2
// 	l4.Right = l6

// 	l2.Left = l1
// 	l2.Right = l3

// 	res := minDiffInBST(l4)

// 	fmt.Println(res)

// }

// func minDiffInBST(root *TreeNode) int {

// 	if root == nil {
// 		return 0
// 	}

// 	if root.Left == nil && root.Right == nil {
// 		return root.Val
// 	}

// 	min := math.MaxInt64
// 	prev := -1

// 	minDiffInBSTHelper(root, &min, &prev)

// 	return min

// }

// // 返回 第一小和 第二小的值
// func minDiffInBSTHelper(root *TreeNode, min *int, prev *int) {

// 	if root == nil {
// 		return
// 	}

// 	minDiffInBSTHelper(root.Left, min, prev)

// 	if *prev != -1 {

// 		subV := tools.Abs(root.Val - *prev)

// 		if subV < *min {
// 			*min = subV
// 		}

// 	}

// 	*prev = root.Val

// 	minDiffInBSTHelper(root.Right, min, prev)

// }
