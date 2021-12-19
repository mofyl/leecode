package Disjoint_Set

import (
	"fmt"
	"sort"
	"testing"
)

func TestAccountsMerge(t *testing.T) {
	//accounts := [][]string{
	//
	//	{"John", "johnsmith@mail.com", "john00@mail.com"},
	//	{"John", "johnnybravo@mail.com"},
	//	{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
	//	{"Mary", "mary@mail.com"},
	//}

	accounts := [][]string{
		{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
		{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
		{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
		{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
		{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
	}

	res := accountsMerge(accounts)
	//res := accountsMerge1(accounts)

	fmt.Println(res)
}

type unionFind721 struct {
	parent []int
	rank   []int
	num    int
}

func buildUnionFind721(n int) unionFind721 {

	p := make([]int, n)
	r := make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = i
	}

	for i := 0; i < n; i++ {
		r[i] = 1
	}

	return unionFind721{
		parent: p,
		rank:   r,
	}
}

// 返回 find的根节点
func (uf unionFind721) find(x int) int {

	if uf.parent[x] == x {
		return x
	}

	// 路径压缩
	uf.parent[x] = uf.find(uf.parent[x])
	return uf.parent[x]
}

func (uf unionFind721) merge(i, j int) {

	rootI := uf.find(i)
	rootJ := uf.find(j)

	// 若已经在同一个集合中了 那么就不用合并了
	if rootI == rootJ {
		return
	}

	// root 不同，将rank 小的合并到 rank 大的上

	if uf.rank[rootI] <= uf.rank[rootJ] {
		uf.parent[rootI] = rootJ
	} else {
		uf.parent[rootJ] = rootI
	}

	// 只有在相等的时候才会去调整rank。若rank本身就有差别，那么就不用调整了
	// 反正也可以分别出来
	if uf.rank[rootI] == uf.rank[rootJ] {
		uf.rank[rootI] += uf.rank[rootJ]
	}

}

func accountsMerge(accounts [][]string) [][]string {

	// 这里 map的值是int，就是为了将 email 和uf 联系起来
	emailUfIdxMap := make(map[string]int, len(accounts))
	nameMap := make(map[string]string, len(accounts))
	for i := 0; i < len(accounts); i++ {
		name := accounts[i][0]
		// 从1开始因为，数据格式 0位置是名字，1... 才是邮箱
		for j := 1; j < len(accounts[i]); j++ {
			// 记录 email 和对应的Num 第一次出现才会去统计.后面重复出现的就不用统计了
			if _, ok := emailUfIdxMap[accounts[i][j]]; !ok {
				emailUfIdxMap[accounts[i][j]] = len(emailUfIdxMap)
				// 保存 每个email 对应的name ，相同email的name一定相同，这里是题目保证的
				nameMap[accounts[i][j]] = name
			}
		}

	}

	uf := buildUnionFind721(len(emailUfIdxMap))

	// 合并 出现相同email 的两个集合
	for i := 0; i < len(accounts); i++ {

		emailFirst := accounts[i][1] // 默认都往 第一个上面 merge
		emailNum := emailUfIdxMap[emailFirst]

		for j := 2; j < len(accounts[i]); j++ {
			emailJ := accounts[i][j]
			emailNumJ := emailUfIdxMap[emailJ]
			uf.merge(emailNum, emailNumJ)
		}

	}

	// 统计uf中每个集合中出现的email
	emailCountMap := make(map[int][]string, len(emailUfIdxMap))
	for email, idx := range emailUfIdxMap {
		// 通过当前email 的idx 找到 root 的idx,然后统计下来.这里只需要统计那些不重复的就行了
		// 因为相同集合的 root Idx 都是相同的，所以这样就可以直接统计
		rootIdx := uf.find(idx)
		emailCountMap[rootIdx] = append(emailCountMap[rootIdx], email)
	}
	// 拼接结果
	res := make([][]string, 0, len(accounts))

	for _, emails := range emailCountMap {
		// 题目要求 emails中内容要按照字典序 所以这里加了排序
		sort.Strings(emails)
		account := append([]string{nameMap[emails[0]]}, emails...)
		res = append(res, account)
	}

	return res
}
