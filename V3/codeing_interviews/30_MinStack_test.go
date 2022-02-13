package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {

	ms := ConstructorMinStack()

	//ms.Push(-2)
	//ms.Push(0)
	//ms.Push(-3)

	ms.Push(0)
	ms.Push(1)
	ms.Push(0)

	fmt.Println(ms.Min())
	ms.Pop()
	fmt.Println(ms.Top())
	fmt.Println(ms.Min())
}

// 这里引入单调栈
type MinStack struct {
	// 栈是先入后出
	// 这里规定 数组尾部 为栈顶，每次入栈 从尾入，出栈从尾出
	stack    []int
	minStack []int // 单调最小栈，该栈严格单调 栈顶是最
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {

	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}

}

func (this *MinStack) Push(x int) {

	this.stack = append(this.stack, x)

	// 这里 等于 单调栈顶 元素 也可以入单调栈
	if len(this.minStack) == 0 || x <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, x)
	}

}

func (this *MinStack) Pop() {

	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]

	// 若 当前元素和 单调栈 minStack 顶元素一致才进行删除
	if val == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[0 : len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}
