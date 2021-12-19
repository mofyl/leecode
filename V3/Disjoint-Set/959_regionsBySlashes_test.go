package Disjoint_Set

import (
	"fmt"
	"testing"
)

func TestRegionsBySlashes(t *testing.T) {
	//grid := []string{
	//	" /",
	//	"/ ",
	//}
	//grid := []string{
	//	" /",
	//	"  ",
	//}

	//grid := []string{
	//	"\\/",
	//	"/\\",
	//}

	grid := []string{
		"/\\",
		"\\/",
	}
	res := regionsBySlashes(grid)
	fmt.Println(res)
}

// 使用并查集 ,这用的是 路径压缩+秩的方式
type unionFind struct {
	parent []int
	rank   []int
	//unionNum int // 这里表示集合个数，刚开始所有的元素自成一个集合，所以就是n个，合并一次就少一个集合
}

func newUnionFind(n int) *unionFind {

	p := make([]int, n)
	r := make([]int, n)

	for i := 0; i < n; i++ {
		// 这里初始化每个集合的根节点都是自己
		p[i] = i
		// 每个集合的 秩 都是1
		r[i] = 1
	}

	return &unionFind{
		parent: p,
		rank:   r,
		//unionNum: n - 1,
	}

}

// 合并 i，j 元素的根节点
func (uf *unionFind) merge(i, j int) {
	// 先找到 i, j 的根节点
	rootI := uf.find(i)
	rootJ := uf.find(j)
	// 这里表示 这两个元素的 根节点相同，这里直接返回
	if rootI == rootJ {
		return
	}

	if uf.rank[rootI] <= uf.rank[rootJ] {
		uf.parent[rootI] = rootJ
	} else {
		uf.parent[rootJ] = rootI
	}

	if uf.rank[rootI] == uf.rank[rootJ] {
		// 修改 rootJ 的 rank 因为若rank相同 我们上面是将 rootI 的父节点置为了 rootJ
		uf.rank[rootJ]++
	}

	//uf.unionNum--
}

func (uf *unionFind) getUnionNum() int {
	res := 0
	for i := 0; i < len(uf.parent); i++ {
		if uf.parent[i] == i {
			res++
		}
	}
	return res
}

// 找到 x 的根节点
func (uf *unionFind) find(x int) int {
	//// 这里表示已经找到根节点了，根节点的标志就是 parent == idx
	if uf.parent[x] == x {
		return uf.parent[x]
	}
	// 这里使用 路径压缩，将 x 的父节点设置为根节点
	uf.parent[x] = uf.find(uf.parent[x])
	// 由于上面已经设置过了，所以这里x 的父节点就是根节点，直接返回即可
	return uf.parent[x]
}

/*
	题解见 笔记
*/
func regionsBySlashes(grid []string) int {
	n := len(grid)

	uf := newUnionFind(4 * n * n)
	/*
		这里确认 1*1格子的编号公式 idx = i*n+j
		根据 1*1格子确认 里面小三角形公式为 idx*4+k  k表示要定位那个三角形，
			一个格子中会有4个三角形 所以这里要乘以4
	*/
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			idx := i*n + j

			// 若 i 不是最后一行，那么i的第2个小三角形一定要和  (i+1 , j)方格 的第0个小三角形合并
			if i < n-1 {
				belowOnIdx := idx + n
				uf.merge(idx*4+2, belowOnIdx*4+0)
			}
			/*
				若 j 不是最后一列，那么就让idx 的第1号三角形和  (i,j+1)的方格的第3号小三角形合并
			*/
			if j < n-1 {
				right := idx + 1
				uf.merge(idx*4+1, right*4+3)
			}

			// 然后处理 grid中限定的合并
			if grid[i][j] == '\\' {
				uf.merge(idx*4+0, idx*4+1)
				uf.merge(idx*4+2, idx*4+3)
			} else if grid[i][j] == '/' {
				uf.merge(idx*4+0, idx*4+3)
				uf.merge(idx*4+1, idx*4+2)
			} else {
				// 若是空格 则将当前格子的所有三角形都合并
				uf.merge(idx*4+0, idx*4+1)
				uf.merge(idx*4+1, idx*4+2)
				uf.merge(idx*4+2, idx*4+3)
			}

		}
	}
	return uf.getUnionNum()
}
