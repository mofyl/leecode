package tree

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestTree2str(t *testing.T) {

	//l1 := &TreeNode{Val: 1}
	//l2 := &TreeNode{Val: 2}
	//l3 := &TreeNode{Val: 3}
	//l4 := &TreeNode{Val: 4}
	//
	//l1.Left = l2
	//l1.Right = l3
	//
	//l2.Right = l4
	//
	//str := tree2str(l1)

	l1 := &TreeNode{Val: 1}
	l2 := &TreeNode{Val: 2}

	l1.Right = l2

	str := tree2str(l1)

	fmt.Println(str)

}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	builder := strings.Builder{}

	tree2str_loop(t, &builder)

	return builder.String()
}

func tree2str_loop(t *TreeNode, str *strings.Builder) {

	if t == nil {
		return
	}
	//
	i := strconv.Itoa(t.Val)

	str.WriteString(i)

	if t.Left == nil && t.Right == nil {
		return
	}
	str.WriteString("(")
	tree2str_loop(t.Left, str)
	str.WriteString(")")
	if t.Right != nil {
		str.WriteString("(")
		tree2str_loop(t.Right, str)
		str.WriteString(")")
	}

}
