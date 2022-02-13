package codeing_interviews

import (
	"fmt"
	"testing"
)

func TestCQueue(t *testing.T) {

	obj := Constructor()
	obj.AppendTail(3)
	obj.AppendTail(4)
	fmt.Println(obj.DeleteHead())
}

type CQueue struct {
	que []int
}

func Constructor() CQueue {
	return CQueue{
		que: make([]int, 0),
	}
}

// 向末尾插入元素
func (this *CQueue) AppendTail(value int) {

	this.que = append(this.que, value)
}

// 删除头部元素
func (this *CQueue) DeleteHead() int {

	if len(this.que) < 1 {
		return -1
	}

	val := this.que[0]

	this.que = this.que[1:]

	return val
}
