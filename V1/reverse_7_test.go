package V1

import (
	"fmt"
	"math"
)

func main7() {
	fmt.Println(reverse(120))
}

func reverse(x int) int {

	var res int
	for x != 0 {

		r := x % 10
		x /= 10

		// 这里 除以10  是因为下面 res要乘以10 ，若这里不做判断  res*10 可能越界
		// 这里也要对r进行判断， 若res这里== maxInt32/10  要是r < math.MaxInt32 % 10 那么后面 res = res*10+r 就不会越界 反之就会越界
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && r > math.MaxInt32%10) {
			return 0
		}

		if res < math.MinInt32/10 || (res == math.MinInt32/10 && r < math.MinInt32%10) {
			return 0
		}

		res = res*10 + r
	}

	return res

}
