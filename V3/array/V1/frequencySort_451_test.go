package V1

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"unsafe"
)

func TestFrequencySort(t *testing.T) {
	str := "Aabb"

	res := frequencySort(str)

	fmt.Println(res)
}

type Segment []string

func (s Segment) Len() int {
	return len(s)
}

func (s Segment) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] > s[j]
	}
	return len(s[i]) > len(s[j])
}

func (s Segment) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func frequencySort(s string) string {
	// 这里表示声明了长度。这张表的长度就是 ASCII码 表的长度。因为ASCII码表最后一个字符是 小写z
	// 这样后面的字符 就 可以直接跟 自己的 ASCII 码做对应，自己的ASCII码就是自己的下标
	// ASCII 码表 不是从0 开始的 是从 '0' 开始的
	counter := ['z' + 1]int{}
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}
	ss := make([]string, 0, len(s))

	for i := 0; i < len(counter); i++ {
		if counter[i] == 0 {
			continue
		}
		n := counter[i]
		tmp := makeStr(byte(i), n)
		ss = append(ss, tmp)
	}

	sort.Sort(Segment(ss))
	builder := strings.Builder{}

	for i := 0; i < len(ss); i++ {
		builder.WriteString(ss[i])
	}
	return builder.String()
}

func makeStr(char byte, n int) string {

	buf := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		buf = append(buf, char)
	}
	// byte 和 string 底层都是一样的 可以直接转换
	return *(*string)(unsafe.Pointer(&buf))
}
