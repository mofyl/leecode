package tools

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int32) int {
	// a 若是正数 则 左移 31位 结果为 0 因为正数的 符号位 为0
	// a 若是负数 则 左移 31位 结果为 -1 ,本来是1，但是最高位也是1 所以就是-1
	// 负数就是对应 正数的补码表示：也就是取反+1
	y := a >> 31
	// 若 a是正数 那么 a ^ 0 = a ,  a - 0 = a
	// 若 a是负数 那么 a ^ -1 = a按位取反  a - (-1) = a+1  所以结果就是 取反+1
	return int(a ^ y - y)
}

func Abs_(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Swap(a, b *int) {

	if *a != *b {
		*a = *a ^ *b
		*b = *a ^ *b
		*a = *a ^ *b
	}
}
