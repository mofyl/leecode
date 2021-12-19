package array

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(nums, k)

	fmt.Println(nums)
}

func rotate(nums []int, k int) {

	/*

		环状替换：
		直接找到 nums[i] 的下一个位置next，然后把next和next上的值记录下来，
			把nums[i] 放过去就行了。
			next=(i+k) % len(nums) 将数据想象成一个环形的。这样循环的将元素替换过去
			直到next=0。循环结束。全部放完

		但是这里要分 len(nums) 和 k 是否有公约数。
			若没有则一次循环就交换完了
			若有 则一次循环不完。 比如 4和2。一轮分组之后 只交换了 0和2. 1,3就没有交换到
		所以最大的交换次数为  len(nums) 和 k的最大公约数(最大公因数)。
			若除1外没有公约数，那么 就是交换1次
			若有则找出公约数,循环这么多次就好了

	*/
	// 这里要分 len(nums),k 是否有公约数。
	//

	n := len(nums)
	// 这里 防止k比n大
	k %= n

	for i := 0; i < gcd(n, k); i++ {

		curNum := nums[i]
		cur := i

		for ok := true; ok; ok = cur != i {
			next := (cur + k) % n
			tmp := nums[next]
			nums[next] = curNum

			curNum = tmp
			cur = next
		}

	}

}

// 求最大公约数 辗转相除法
// 较大的数为a , 较小的数为 b
func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
